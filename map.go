package stream

import (
	"errors"
	"reflect"
)

// Map returns a stream consisting of the results of applying the given function to the elements of this stream.
// - mapper: a non-interfering, stateless function to apply to each element
func (s *Stream) Map(mapper interface{}) *Stream {
	if s.Error != nil {
		return s
	}

	s.funcs = append(s.funcs, func(i *reflect.Value) (bool, *reflect.Value) {
		fn := reflect.ValueOf(mapper)
		if fn.Kind() != reflect.Func {
			s.Error = errors.New("Map mapper type is not Fun")
			return true, i
		}
		if fn.Type().NumIn() != 1 {
			s.Error = errors.New("Map mapper's input parameter length is not one")
			return true, i
		}
		if fn.Type().NumOut() != 1 {
			s.Error = errors.New("Map mapper's output parameter length is not one")
			return true, i
		}
		var param [1]reflect.Value
		param[0] = *i
		return false, &fn.Call(param[:])[0]
	})

	return s
}
