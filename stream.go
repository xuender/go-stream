package stream

import (
	"errors"
	"reflect"
)

// Stream inspired in Java 8 Streams.
type Stream struct {
	value    *reflect.Value
	funcs    []Operation
	err      error
	parallel bool
	stop     bool
}

// Operation terminal operation.
type Operation func(*reflect.Value) (bool, *reflect.Value)

func (s *Stream) evaluate(terminalOp Operation) (*reflect.Value, error) {
	switch s.value.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < s.value.Len(); i++ {
			a := s.value.Index(i)
			ok := true
			for _, f := range s.funcs {
				if stop, t := f(&a); stop {
					ok = false
					break
				} else {
					a = *t
				}
			}
			if ok {
				if stop, e := terminalOp(&a); stop {
					return e, nil
				}
			}
		}
		return nil, errors.New("No find")
	default:
		return nil, errors.New("array type is not Slice or Array")
	}
}

// NewStream returns a sequential Stream.
func NewStream(array interface{}) *Stream {
	v := reflect.ValueOf(array)
	return &Stream{
		value: &v,
		funcs: []Operation{},
	}
}
