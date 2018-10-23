package stream

import (
	"errors"
	"reflect"
	"sort"
)

// Min returns the minimum element of this stream according to the provided `less`.
func (s *Stream) Min(less interface{}) (interface{}, error) {
	if s.err != nil {
		return nil, s.err
	}
	fn := reflect.ValueOf(less)
	if fn.Kind() != reflect.Func {
		return nil, errors.New("Min less type is not Func")
	}
	if fn.Type().NumIn() != 2 {
		return nil, errors.New("Min less's input parameter length not two")
	}
	if fn.Type().NumOut() != 1 {
		return nil, errors.New("Min less's output parameter length not one")
	}
	if fn.Type().Out(0).Kind() != reflect.Bool {
		return nil, errors.New("Min less's output parameter type is not Bool")
	}

	t := s.value.Type()
	ret := reflect.New(t).Elem()
	ret.Set(reflect.MakeSlice(t, 0, s.value.Len()))
	ret.SetLen(s.value.Len())
	n := 0
	o := func(i *reflect.Value) (bool, *reflect.Value) {
		ret.Index(n).Set(*i)
		n++
		return false, i
	}

	var err error
	if s.parallel {
		_, err = s.parallelEvaluate(o)
	} else {
		_, err = s.evaluate(o)
	}
	if err != nil && err != errNotFound {
		return nil, err
	}
	ret.SetLen(n)
	s.value = &ret
	s.sortFunc = &fn
	sort.Sort(s)
	return s.value.Index(0).Interface(), nil
}
