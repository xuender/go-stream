package stream

import "reflect"

// Sequential returns an equivalent stream that is sequential.
func (s *Stream) Sequential() *Stream {
	return s.setParallel(false)
}

func (s *Stream) evaluate(terminalOp Operation) (*reflect.Value, error) {
	switch s.value.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < s.value.Len(); i++ {
			a := s.value.Index(i)
			ok := true
			for _, f := range s.funcs {
				if stop, t := f(&a); stop {
					ok = false
					break
				} else {
					a = *t
				}
			}
			if ok {
				if stop, e := terminalOp(&a); stop {
					return e, nil
				}
			}
		}
		return nil, errNotFound
	default:
		return nil, errArrayTypeError
	}
}
