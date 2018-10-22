package stream

import (
	"errors"
	"reflect"
)

// AnyMatch Returns whether any elements of this stream match the provided predicate.
func (s *Stream) AnyMatch(predicate interface{}) (bool, error) {
	if s.err != nil {
		return false, s.err
	}
	fn := reflect.ValueOf(predicate)
	if fn.Kind() != reflect.Func {
		return false, errors.New("AnyMatch predicate type is not Func")
	}
	if fn.Type().NumIn() != 1 {
		return false, errors.New("AnyMatch predicate's input parameter length not one")
	}
	if fn.Type().NumOut() != 1 {
		return false, errors.New("AnyMatch predicate's output parameter length not one")
	}
	if fn.Type().Out(0).Kind() != reflect.Bool {
		return false, errors.New("AnyMatch predicate's output parameter type is not Bool")
	}

	o := func(i *reflect.Value) (bool, *reflect.Value) {
		var param [1]reflect.Value
		param[0] = *i
		return fn.Call(param[:])[0].Bool(), i
	}

	var err error
	if s.parallel {
		_, err = s.parallelEvaluate(o)
	} else {
		_, err = s.evaluate(o)
	}
	if err == nil {
		return true, nil
	}
	if err == errNotFound {
		return false, nil
	}
	return false, err
}
