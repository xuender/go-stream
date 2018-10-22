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

var errNotFound = errors.New("Not found")
var errArrayTypeError = errors.New("array type is not Slice and Array")

// New returns a Stream.
func New(array interface{}) *Stream {
	v := reflect.ValueOf(array)
	return &Stream{
		value:    &v,
		funcs:    []Operation{},
		parallel: false,
		stop:     false,
	}
}
