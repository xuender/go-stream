package stream

import (
	"reflect"
)

// FindFirst returns the first element of this stream.
func (s *Stream) FindFirst() (interface{}, error) {
	if s.err != nil {
		return nil, s.err
	}

	operation := func(i *reflect.Value) (bool, *reflect.Value) { return true, i }

	var v *reflect.Value
	var err error
	if s.parallel {
		v, err = s.parallelEvaluate(operation)
	} else {
		v, err = s.evaluate(operation)
	}
	if err == nil {
		return v.Interface(), nil
	}
	return nil, err
}
