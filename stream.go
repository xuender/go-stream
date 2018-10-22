package stream

import (
	"errors"
	"reflect"
)

// Stream inspired in Java 8 Streams.
type Stream struct {
	value    *reflect.Value
	funcs    []Operation
	Error    error
	parallel bool
	stop     bool
	empty    bool
}

// Operation terminal operation.
type Operation func(*reflect.Value) (bool, *reflect.Value)

var errNotFound = errors.New("Not found")
var errArrayTypeError = errors.New("array type is not Slice and Array")
var errEmpty = errors.New("Stream is empty")

// New returns a Stream.
func New(array interface{}) *Stream {
	v := reflect.ValueOf(array)
	return &Stream{
		value:    &v,
		funcs:    []Operation{},
		parallel: false,
		stop:     false,
		empty:    false,
	}
}

// NewEmpty returns a empty Stream.
func NewEmpty() *Stream {
	return &Stream{
		funcs:    []Operation{},
		parallel: false,
		stop:     false,
		empty:    true,
	}
}

// Set data
func (s *Stream) Set(i interface{}) *Stream {
	v := reflect.ValueOf(i)
	s.value = &v
	s.empty = false
	return s
}
