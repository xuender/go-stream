package stream

import (
	"reflect"
)

// Skip returns a stream consisting of the remaining elements of this stream
// after discarding the first `n` elements of the stream.
func (s *Stream) Skip(n int) *Stream {
	if s.err != nil {
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
	operation := func(i *reflect.Value) (bool, *reflect.Value) {
		if num >= n {
			ret.Index(num - n).Set(*i)
		}
		num++
		return false, i
	}

	var err error
	if s.parallel {
		_, err = s.parallelEvaluate(operation)
	} else {
		_, err = s.evaluate(operation)
	}
	if err != nil && err != errNotFound {
		s.err = err
		return s
	}
	s.funcs = []func(*reflect.Value) []*reflect.Value{}
	s.value = &ret
	return s
}
