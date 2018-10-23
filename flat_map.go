package stream

import (
	"errors"
	"reflect"
)

// FlatMap returns a stream consisting of the results of replacing each element of
// this stream with the contents of a mapped stream produced by applying
// the provided mapping function to each element.
func (s *Stream) FlatMap(mapper interface{}) *Stream {
	if s.err != nil {
		return s
	}
	s.funcs = append(s.funcs, func(i *reflect.Value) []*reflect.Value {
		fn := reflect.ValueOf(mapper)
		if fn.Kind() != reflect.Func {
			s.err = errors.New("FlatMap mapper type is not Func")
			return emptyValues
		}
		if fn.Type().NumIn() != 1 {
			s.err = errors.New("FlatMap mapper's input parameter length not one")
			return emptyValues
		}
		if fn.Type().NumOut() != 1 {
			s.err = errors.New("FlatMap mapper's output parameter length not one")
			return emptyValues
		}
		if fn.Type().Out(0).Kind() != reflect.Slice && fn.Type().Out(0).Kind() != reflect.Array {
			s.err = errors.New("FlatMap mapper's output parameter type is not Slice and Array")
			return emptyValues
		}
		var param [1]reflect.Value
		param[0] = *i
		out := fn.Call(param[:])[0]
		ret := make([]*reflect.Value, out.Len(), out.Len())
		for i := 0; i < out.Len(); i++ {
			v := out.Index(i)
			ret[i] = &v
		}
		return ret
	})
	return s
}
