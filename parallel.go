package stream

import (
	"errors"
	"reflect"
)

type valueCallBack struct {
	v    *reflect.Value
	stop bool
}

// Parallel returns an equivalent stream that is parallel.
func (s *Stream) Parallel() *Stream {
	s.parallel = true
	return s
}
func (s *Stream) parallelEvaluate(terminalOp Operation) (*reflect.Value, error) {
	switch s.value.Kind() {
	case reflect.Slice, reflect.Array:
		s.stop = false
		cbC := make(chan *valueCallBack, 1)
		for i := 0; i < s.value.Len(); i++ {
			a := s.value.Index(i)
			go func(a *reflect.Value) {
				ok := true
				for _, f := range s.funcs {
					if s.stop {
						cbC <- &valueCallBack{
							v:    a,
							stop: true,
						}
						return
					}
					if stop, t := f(a); stop {
						ok = false
						break
					} else {
						a = t
					}
				}
				if s.stop {
					cbC <- &valueCallBack{
						v:    a,
						stop: true,
					}
					return
				}
				if ok {
					if stop, e := terminalOp(a); stop {
						s.stop = true
						cbC <- &valueCallBack{
							v:    e,
							stop: true,
						}
						return
					}
				}
				cbC <- &valueCallBack{
					v:    a,
					stop: false,
				}
			}(&a)
		}
		var ret *reflect.Value
		err := errors.New("Not find")
		for i := 0; i < s.value.Len(); i++ {
			cb := <-cbC
			if err != nil && cb.stop {
				ret = cb.v
				err = nil
			}
		}
		return ret, err
	default:
		return nil, errors.New("array type is not Slice or Array")
	}
}
