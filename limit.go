package stream

import (
	"reflect"
)

// Limit returns a stream consisting of the elements of this stream.
func (s *Stream) Limit(maxSize int) *Stream {
	if s.err != nil {
		return s
	}
	if maxSize > s.value.Len() {
		return s
	}

	size := 0
	t := s.value.Type()
	ret := reflect.New(t).Elem()
	ret.Set(reflect.MakeSlice(t, 0, maxSize))
	ret.SetLen(maxSize)
	o := func(i *reflect.Value) (bool, *reflect.Value) {
		if size >= maxSize {
			return true, i
		}
		ret.Index(size).Set(*i)
		size++
		return false, i
	}

	var err error
	if s.parallel {
		_, err = s.parallelEvaluate(o)
	} else {
		_, err = s.evaluate(o)
	}
	if err != nil && err != errNotFound {
		s.err = err
		return s
	}
	s.funcs = []Operation{}
	s.value = &ret
	return s
}
