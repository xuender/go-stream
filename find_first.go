package stream

import (
	"errors"
	"reflect"
)

// FindFirst returns the first element of this stream.
func (s *Stream) FindFirst() (interface{}, error) {
	if s.err != nil {
		return nil, s.err
	}

	switch s.value.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < s.value.Len(); i++ {
			a := s.value.Index(i)
			ok := true
			for _, f := range s.funcs {
				if o, i := f(&a); o {
					a = *i
				} else {
					ok = false
					break
				}
			}
			if ok {
				return a.Interface(), nil
			}
		}
		return nil, errors.New("No find")
	default:
		return nil, errors.New("array type is not Slice or Array")
	}
}
