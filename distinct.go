package stream

import (
	"reflect"
)

// Distinct returns a stream consisting of the distinct elements of this stream.
func (s *Stream) Distinct() *Stream {
	if s.Error != nil {
		return s
	}

	t := s.value.Type()
	m := map[interface{}]*reflect.Value{}
	o := func(i *reflect.Value) (bool, *reflect.Value) {
		m[i.Interface()] = i
		return false, i
	}

	var err error
	if s.parallel {
		_, err = s.parallelEvaluate(o)
	} else {
		_, err = s.evaluate(o)
	}
	if err != nil && err != errNotFound {
		s.Error = err
		return s
	}
	s.funcs = []Operation{}
	ret := reflect.New(t).Elem()
	ret.Set(reflect.MakeSlice(t, 0, len(m)))
	ret.SetLen(len(m))
	i := 0
	for _, v := range m {
		ret.Index(i).Set(*v)
		i++
	}
	s.value = &ret
	return s
}
