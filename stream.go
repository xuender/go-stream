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
	sortFunc *reflect.Value
}

func (s *Stream) Len() int {
	return s.value.Len()
}

func (s *Stream) Swap(i, j int) {
	v := s.value.Index(i).Interface()
	s.value.Index(i).Set(s.value.Index(j))
	s.value.Index(j).Set(reflect.ValueOf(v))
}

func (s Stream) Less(i, j int) bool {
	var param [2]reflect.Value
	param[0] = s.value.Index(i)
	param[1] = s.value.Index(j)
	return s.sortFunc.Call(param[:])[0].Bool()
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
