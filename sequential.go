package stream

import (
	"reflect"
)

// Sequential returns an equivalent stream that is sequential.
func (s *Stream) Sequential() *Stream {
	return s.setParallel(false)
}

func (s *Stream) run(value *reflect.Value, i int) []*reflect.Value {
	if s.stop {
		return emptyValues
	}
	os := s.funcs[i](value)
	i++
	if i >= len(s.funcs) {
		return os
	}
	num := 0
	for num < len(os) {
		ns := s.run(os[num], i)
		end := os[num+1:]
		os = append(os[:num], ns...)
		num++
		num += len(ns)
		os = append(os, end...)
	}
	return os
}

func (s *Stream) evaluate(terminalOp Operation) (*reflect.Value, error) {
	switch s.value.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < s.value.Len(); i++ {
			a := s.value.Index(i)
			if len(s.funcs) == 0 {
				if stop, e := terminalOp(&a); stop {
					return e, nil
				}
			} else {
				for _, v := range s.run(&a, 0) {
					if stop, e := terminalOp(v); stop {
						return e, nil
					}
				}
			}
		}
		return nil, errNotFound
	default:
		return nil, errArrayTypeError
	}
}
