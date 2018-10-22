package stream

import (
	"errors"
	"reflect"
)

// Filter returns a stream consisting of the elements of this stream that match the given predicate.
func (s *Stream) Filter(predicate interface{}) *Stream {
	if s.Error != nil {
		return s
	}

	s.funcs = append(s.funcs, func(i *reflect.Value) (bool, *reflect.Value) {
		fn := reflect.ValueOf(predicate)
		if fn.Kind() != reflect.Func {
			s.Error = errors.New("Filter predicate type is not Func")
			return true, i
		}
		if fn.Type().NumIn() != 1 {
			s.Error = errors.New("Filter predicate's input parameter length not one")
			return true, i
		}
		if fn.Type().NumOut() != 1 {
			s.Error = errors.New("Filter predicate's output parameter length not one")
			return true, i
		}
		if fn.Type().Out(0).Kind() != reflect.Bool {
			s.Error = errors.New("Filter predicate's out param type is not Bool")
			return true, i
		}
		var param [1]reflect.Value
		param[0] = *i
		return !fn.Call(param[:])[0].Bool(), i
	})

	return s
}
