package stream

import (
	"errors"
	"reflect"
)

// Sorted TODO
func (s *Stream) Sorted() *Stream {
	if s.err != nil {
		return s
	}
	switch s.value.Kind() {
	case reflect.Slice, reflect.Array:
		na := []interface{}{}
		for i := 0; i < s.value.Len(); i++ {
			a := s.value.Index(i)
			ok := true
			for _, f := range s.funcs {
				if o, v := f(&a); o {
					a = *v
				} else {
					ok = false
					break
				}
			}
			if ok {
				na = append(na, a.Interface())
			}
		}
		// TODO Sort
		// sort.Ints(na)
		v := reflect.ValueOf(na)
		s.value = &v
		s.funcs = []func(*reflect.Value) (bool, *reflect.Value){}
		return s
	default:
		s.err = errors.New("array type is not Slice or Array")
		return s
	}
}
