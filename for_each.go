package stream

import (
	"errors"
	"reflect"
)

// ForEach performs an action for each element of this stream.
func (s *Stream) ForEach(action interface{}) error {
	if s.err != nil {
		return s.err
	}
	fn := reflect.ValueOf(action)
	if fn.Kind() != reflect.Func {
		return errors.New("ForEach action type is not Fun")
	}
	if fn.Type().NumIn() != 1 {
		return errors.New("ForEach action's input parameter length not one")
	}
	if fn.Type().NumOut() != 0 {
		return errors.New("ForEach action's output parameter length not zero")
	}

	o := func(i *reflect.Value) (bool, *reflect.Value) {
		var param [1]reflect.Value
		param[0] = *i
		fn.Call(param[:])
		return false, i
	}

	var err error
	if s.parallel {
		_, err = s.parallelEvaluate(o)
	} else {
		_, err = s.evaluate(o)
	}
	if err != nil && err != errNotFound {
		return err
	}
	return nil
}
