package stream

import (
	"reflect"
)

// Stream inspired in Java 8 Streams.
type Stream struct {
	value *reflect.Value
	funcs []func(*reflect.Value) (bool, *reflect.Value)
	err   error
}

// NewStream returns a sequential Stream.
func NewStream(array interface{}) *Stream {
	v := reflect.ValueOf(array)
	return &Stream{
		value: &v,
		funcs: []func(*reflect.Value) (bool, *reflect.Value){},
	}
}
