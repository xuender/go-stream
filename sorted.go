package stream

import (
	"errors"
	"reflect"
	"sort"
)

// Sorted returns a stream consisting of the elements of this stream,
// sorted according to the provided `less`.
func (s *Stream) Sorted(less interface{}) *Stream {
	if s.err != nil {
		return s
	}
	fn := reflect.ValueOf(less)
	if fn.Kind() != reflect.Func {
		s.err = errors.New("Sorted less type is not Func")
		return s
	}
	if fn.Type().NumIn() != 2 {
		s.err = errors.New("Sorted less's input parameter length not two")
		return s
	}
	if fn.Type().NumOut() != 1 {
		s.err = errors.New("Sorted less's output parameter length not one")
		return s
	}
	if fn.Type().Out(0).Kind() != reflect.Bool {
		s.err = errors.New("Sorted less's output parameter type is not Bool")
		return s
	}

	operation := func(i *reflect.Value) (bool, *reflect.Value) { return false, i }

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
	s.sortFunc = &fn
	sort.Sort(s)
	return s
}
