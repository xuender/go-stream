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

	s.funcs = append(s.funcs, func(i *reflect.Value) (bool, *reflect.Value) {
		fn := reflect.ValueOf(predicate)
		if fn.Kind() != reflect.Func {
			s.err = errors.New("Filter predicate type is not function.")
			return false, i
		}
		if fn.Type().NumIn() != 1 {
			s.err = errors.New("Filter predicate's in params length not one.")
			return false, i
		}
		if fn.Type().NumOut() != 1 {
			s.err = errors.New("Filter predicate's out params length not one.")
			return false, i
		}
		if fn.Type().Out(0).Kind() != reflect.Bool {
			s.err = errors.New("Filter predicate's out param type is not Bool.")
			return false, i
		}
		var param [1]reflect.Value
		param[0] = *i
		return fn.Call(param[:])[0].Bool(), i
	})

	return s
}
