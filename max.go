package stream

import (
	"errors"
	"reflect"
	"sort"
)

// Max returns the maximum element of this stream according to the provided `less`.
func (s *Stream) Max(less interface{}) (interface{}, error) {
	if s.err != nil {
		return nil, s.err
	}
	fn := reflect.ValueOf(less)
	if fn.Kind() != reflect.Func {
		return nil, errors.New("Max less type is not Func")
	}
	if fn.Type().NumIn() != 2 {
		return nil, errors.New("Max less's input parameter length not two")
	}
	if fn.Type().NumOut() != 1 {
		return nil, errors.New("Max less's output parameter length not one")
	}
	if fn.Type().Out(0).Kind() != reflect.Bool {
		return nil, errors.New("Max less's output parameter type is not Bool")
	}

	t := s.value.Type()
	ret := reflect.New(t).Elem()
	ret.Set(reflect.MakeSlice(t, 0, s.value.Len()))
	ret.SetLen(s.value.Len())
	n := 0
	operation := func(i *reflect.Value) (bool, *reflect.Value) {
		ret.Index(n).Set(*i)
		n++
		return false, i
	}

	var err error
	if s.parallel {
		_, err = s.parallelEvaluate(operation)
	} else {
		_, err = s.evaluate(operation)
	}
	if err != nil && err != errNotFound {
		return nil, err
	}
	ret.SetLen(n)
	s.value = &ret
	s.sortFunc = &fn
	sort.Sort(s)
	return s.value.Index(s.value.Len() - 1).Interface(), nil
}
