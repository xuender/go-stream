package stream

import (
	"reflect"
)

// FindFirst returns the first element of this stream.
func (s *Stream) FindFirst() (interface{}, error) {
	if s.Error != nil {
		return nil, s.Error
	}
	if s.empty {
		return nil, errEmpty
	}

	o := func(i *reflect.Value) (bool, *reflect.Value) { return true, i }

	var v *reflect.Value
	var err error
	if s.parallel {
		v, err = s.parallelEvaluate(o)
	} else {
		v, err = s.evaluate(o)
	}
	if err == nil {
		return v.Interface(), nil
	}
	return nil, err
}
