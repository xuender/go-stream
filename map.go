package stream

import (
	"errors"
	"reflect"
)

// Map returns a stream consisting of the results of applying the given function to the elements of this stream.
// - mapper: a non-interfering, stateless function to apply to each element
func (s *Stream) Map(mapper interface{}) *Stream {
	if s.err != nil {
		return s
	}

	s.funcs = append(s.funcs, func(i *reflect.Value) (bool, *reflect.Value) {
		fn := reflect.ValueOf(mapper)
		if fn.Kind() != reflect.Func {
			s.err = errors.New("Map mapper type is not function")
			return false, i
		}
		if fn.Type().NumIn() != 1 {
			s.err = errors.New("Map mapper's in params length is not one.")
			return false, i
		}
		if fn.Type().NumOut() != 1 {
			s.err = errors.New("Map mapper's out params length is not one.")
			return false, i
		}
		var param [1]reflect.Value
		param[0] = *i
		return true, &fn.Call(param[:])[0]
	})

	return s
}
