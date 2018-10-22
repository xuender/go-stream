package stream

import "reflect"

// Count returns the count of elements in this stream.
func (s *Stream) Count() (int, error) {
	if s.err != nil {
		return 0, s.err
	}

	count := 0
	o := func(i *reflect.Value) (bool, *reflect.Value) {
		count++
		return false, i
	}

	var err error
	if s.parallel {
		_, err = s.parallelEvaluate(o)
	} else {
		_, err = s.evaluate(o)
	}
	if err != nil && err != errNotFound {
		return 0, err
	}
	return count, nil
}
