package stream

import "reflect"

// Count returns the count of elements in this stream.
func (s *Stream) Count() (int, error) {
	if s.Error != nil {
		return 0, s.Error
	}
	if s.empty {
		return 0, errEmpty
	}

	count := 0
	o := func(i *reflect.Value) (bool, *reflect.Value) {
		count++
		return false, i
	}

	if s.parallel {
		s.parallelEvaluate(o)
		return count, nil
	}
	s.evaluate(o)
	return count, nil
}
