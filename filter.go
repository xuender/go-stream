package stream

import (
	"errors"
	"reflect"
)

// Filter returns a stream consisting of the elements of this stream that match the given predicate.
func (s *Stream) Filter(predicate interface{}) *Stream {
	if s.err != nil {
		return s
	}

	s.funcs = append(s.funcs, func(i *reflect.Value) []*reflect.Value {
		fn := reflect.ValueOf(predicate)
		if fn.Kind() != reflect.Func {
			s.err = errors.New("Filter predicate type is not Func")
			return emptyValues
		}
		if fn.Type().NumIn() != 1 {
			s.err = errors.New("Filter predicate's input parameter length not one")
			return emptyValues
		}
		if fn.Type().NumOut() != 1 {
			s.err = errors.New("Filter predicate's output parameter length not one")
			return emptyValues
		}
		if fn.Type().Out(0).Kind() != reflect.Bool {
			s.err = errors.New("Filter predicate's output parameter type is not Bool")
			return emptyValues
		}
		var param [1]reflect.Value
		param[0] = *i
		if fn.Call(param[:])[0].Bool() {
			return []*reflect.Value{i}
		}
		return emptyValues
	})

	return s
}
