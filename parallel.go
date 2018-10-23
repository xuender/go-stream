package stream

import (
	"reflect"
)

type valueCallBack struct {
	v    *reflect.Value
	stop bool
}

// Parallel returns an equivalent stream that is parallel.
func (s *Stream) Parallel() *Stream {
	return s.setParallel(true)
}

func (s *Stream) setParallel(p bool) *Stream {
	if s.err != nil {
		return s
	}

	o := func(i *reflect.Value) (bool, *reflect.Value) { return false, i }

	var err error
	if s.parallel {
		_, err = s.parallelEvaluate(o)
	} else {
		_, err = s.evaluate(o)
	}
	if err != nil && err != errNotFound {
		s.err = err
		return s
	}
	s.funcs = []func(*reflect.Value) []*reflect.Value{}
	s.parallel = p
	return s
}

func (s *Stream) parallelEvaluate(terminalOp Operation) (*reflect.Value, error) {
	switch s.value.Kind() {
	case reflect.Slice, reflect.Array:
		s.stop = false
		cbC := make(chan *valueCallBack, 1)
		for i := 0; i < s.value.Len(); i++ {
			a := s.value.Index(i)
			if len(s.funcs) == 0 {
				return &a, nil
			}
			go func(a *reflect.Value) {
				var ret *reflect.Value
				for _, v := range s.run(a, 0) {
					if stop, e := terminalOp(v); stop {
						ret = e
						s.stop = true
						break
					}
				}
				cbC <- &valueCallBack{
					v:    ret,
					stop: false,
				}
			}(&a)
		}
		var ret *reflect.Value
		err := errNotFound
		for i := 0; i < s.value.Len(); i++ {
			cb := <-cbC
			if err != nil && cb.stop {
				ret = cb.v
				err = nil
			}
		}
		return ret, err
	default:
		return nil, errArrayTypeError
	}
}
