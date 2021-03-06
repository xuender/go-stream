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

	s.funcs = append(s.funcs, func(i *reflect.Value) []*reflect.Value {
		fn := reflect.ValueOf(action)
		if fn.Kind() != reflect.Func {
			s.err = errors.New("Peek action type is not Fun")
			return emptyValues
		}
		if fn.Type().NumIn() != 1 {
			s.err = errors.New("Peek action's input parameter length is not one")
			return emptyValues
		}
		if fn.Type().NumOut() != 0 {
			s.err = errors.New("Peek action's output parameter length is not zero")
			return emptyValues
		}
		var param [1]reflect.Value
		param[0] = *i
		fn.Call(param[:])
		return []*reflect.Value{i}
	})

	return s
}
