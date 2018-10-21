package stream

import (
	"errors"
	"sort"
)

// intStream test
type intStream struct {
	array []int
	funcs []func(int) (bool, int)
}

func (s *intStream) Filter(predicate func(int) bool) *intStream {
	s.funcs = append(s.funcs, func(i int) (bool, int) {
		return predicate(i), i
	})
	return s
}
func (s *intStream) Map(iteratee func(int) int) *intStream {
	s.funcs = append(s.funcs, func(i int) (bool, int) {
		return true, iteratee(i)
	})
	return s
}
func (s *intStream) Peek(iteratee func(int)) *intStream {
	s.funcs = append(s.funcs, func(i int) (bool, int) {
		iteratee(i)
		return true, i
	})
	return s
}
func (s *intStream) Sorted() *intStream {
	na := []int{}
	for _, a := range s.array {
		ok := true
		for _, f := range s.funcs {
			if o, i := f(a); o {
				a = i
			} else {
				ok = false
				break
			}
		}
		if ok {
			na = append(na, a)
		}
	}
	sort.Ints(na)
	s.array = na
	s.funcs = []func(int) (bool, int){}
	return s
}

func (s *intStream) FindFirst() (int, error) {
	for _, a := range s.array {
		ok := true
		for _, f := range s.funcs {
			if o, i := f(a); o {
				a = i
			} else {
				ok = false
				break
			}
		}
		if ok {
			return a, nil
		}
	}
	return 0, errors.New("No find")
}

func newintStream(array []int) *intStream {
	return &intStream{
		array: array,
		funcs: []func(int) (bool, int){},
	}
}
