package stream

import (
	"errors"
	"reflect"
)

// Peek returns a stream consisting of the elements of this stream.
func (s *Stream) Peek(action interface{}) *Stream {
	if s.err != nil {
		return s
	}

	s.funcs = append(s.funcs, func(i *reflect.Value) (bool, *reflect.Value) {
		fn := reflect.ValueOf(action)
		if fn.Kind() != reflect.Func {
			s.err = errors.New("Peek action type is not function.")
			return false, i
		}
		if fn.Type().NumIn() != 1 {
			s.err = errors.New("Peek action's in params length is not one.")
			return false, i
		}
		if fn.Type().NumOut() != 0 {
			s.err = errors.New("Peek action's out params length is not zero.")
			return false, i
		}
		var param [1]reflect.Value
		param[0] = *i
		fn.Call(param[:])
		return true, i
	})

	return s
}
