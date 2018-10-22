package stream

import (
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
				if stop, v := f(&a); stop {
					ok = false
					break
				} else {
					a = *v
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
		s.funcs = []Operation{}
		return s
	default:
		s.err = errArrayTypeError
		return s
	}
}
