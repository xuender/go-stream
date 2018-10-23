package stream

import (
	"errors"
	"reflect"
)

// Reduce performs a reduction on the elements of this stream.
func (s *Stream) Reduce(accumulator interface{}) (interface{}, error) {
	if s.err != nil {
		return nil, s.err
	}
	fn := reflect.ValueOf(accumulator)
	if fn.Kind() != reflect.Func {
		return nil, errors.New("Reduce accumlator type is not Func")
	}
	if fn.Type().NumIn() != 2 {
		return nil, errors.New("Reduce accumlator's input parameter length not two")
	}
	if fn.Type().NumOut() != 1 {
		return nil, errors.New("Reduce accumlator's output parameter length not one")
	}
	if fn.Type().Out(0).Kind() != fn.Type().In(0).Kind() && fn.Type().Out(0).Kind() != fn.Type().In(1).Kind() {
		return nil, errors.New("Reduce accumlator's output parameter type is not input parameter")
	}

	var ret *reflect.Value
	isNew := true
	operation := func(i *reflect.Value) (bool, *reflect.Value) {
		if isNew {
			isNew = false
			ret = i
		} else {
			var param [2]reflect.Value
			param[0] = *ret
			param[1] = *i
			v := fn.Call(param[:])[0]
			ret = &v
		}
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
	if isNew {
		return nil, nil
	}
	return ret.Interface(), nil
}
