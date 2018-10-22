package stream

import (
	"errors"
	"reflect"
)

// AllMatch returns whether all elements of this stream match the provided predicate.
func (s *Stream) AllMatch(predicate interface{}) (bool, error) {
	if s.err != nil {
		return false, s.err
	}
	fn := reflect.ValueOf(predicate)
	if fn.Kind() != reflect.Func {
		return false, errors.New("AllMatch predicate type is not Func")
	}
	if fn.Type().NumIn() != 1 {
		return false, errors.New("AllMatch predicate's input parameter length not one")
	}
	if fn.Type().NumOut() != 1 {
		return false, errors.New("AllMatch predicate's output parameter length not one")
	}
	if fn.Type().Out(0).Kind() != reflect.Bool {
		return false, errors.New("AllMatch predicate's out param type is not Bool")
	}

	o := func(i *reflect.Value) (bool, *reflect.Value) {
		var param [1]reflect.Value
		param[0] = *i
		if !fn.Call(param[:])[0].Bool() {
			return true, i
		}
		return false, i
	}

	var err error
	if s.parallel {
		_, err = s.parallelEvaluate(o)
	} else {
		_, err = s.evaluate(o)
	}
	if err == errNotFound {
		return true, nil
	}
	return false, err
}
