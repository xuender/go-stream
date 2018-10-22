package stream

import (
	"reflect"
)

// Skip returns a stream consisting of the remaining elements of this stream
// after discarding the first `n` elements of the stream.
func (s *Stream) Skip(n int) *Stream {
	if s.Error != nil {
		return s
	}
	if s.empty {
		s.Error = errEmpty
		return s
	}

	size := s.value.Len() - n
	if size < 0 {
		size = 0
	}
	if n > s.value.Len() {
		n = s.value.Len()
	}

	t := s.value.Type()
	ret := reflect.New(t).Elem()
	ret.Set(reflect.MakeSlice(t, 0, size))
	ret.SetLen(size)
	num := 0
	o := func(i *reflect.Value) (bool, *reflect.Value) {
		if num >= n {
			ret.Index(num - n).Set(*i)
		}
		num++
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
	s.value = &ret
	return s
}
