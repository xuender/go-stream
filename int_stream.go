package stream

import (
	"errors"
	"sort"
)

// intStream test
type intStream struct {
	array    []int
	funcs    []func(int) (bool, int)
	err      error
	parallel bool
	stop     bool
}

func (s *intStream) Parallel() *intStream {
	s.parallel = true
	return s
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

type callBack struct {
	i    int
	stop bool
}

func (s *intStream) parallelEvaluate(terminalOp func(i int) (bool, int)) (int, error) {
	s.stop = false
	cbC := make(chan *callBack, 1)
	for _, aa := range s.array {
		go func(a int) {
			ok := true
			for _, f := range s.funcs {
				if s.stop {
					cbC <- &callBack{
						i:    0,
						stop: true,
					}
					return
				}
				if o, t := f(a); o {
					a = t
				} else {
					ok = false
					break
				}
			}
			if s.stop {
				cbC <- &callBack{
					i:    0,
					stop: true,
				}
				return
			}
			if ok {
				if ret, e := terminalOp(a); ret {
					s.stop = true
					cbC <- &callBack{
						i:    e,
						stop: true,
					}
					return
				}
			}
			cbC <- &callBack{
				i:    0,
				stop: false,
			}
		}(aa)
	}
	var ret int
	err := errors.New("Not find")
	for range s.array {
		cb := <-cbC
		if err != nil && cb.stop {
			ret = cb.i
			err = nil
		}
	}
	return ret, err
}
func (s *intStream) evaluate(terminalOp func(i int) (bool, int)) (int, error) {
	for _, a := range s.array {
		ok := true
		for _, f := range s.funcs {
			if stop, t := f(a); stop {
				ok = false
				break
			} else {
				a = t
			}
		}
		if ok {
			if ret, e := terminalOp(a); ret {
				return e, nil
			}
		}
	}
	return 0, errors.New("Not find")
}
func (s *intStream) ForEach(iteratee func(int)) error {
	if s.err != nil {
		return s.err
	}
	if s.parallel {
		s.parallelEvaluate(func(i int) (bool, int) {
			iteratee(i)
			return false, 0
		})
	} else {
		s.evaluate(func(i int) (bool, int) {
			iteratee(i)
			return false, 0
		})
	}
	return nil
}
func (s *intStream) Peek(iteratee func(int)) *intStream {
	s.funcs = append(s.funcs, func(i int) (bool, int) {
		iteratee(i)
		return true, i
	})
	return s
}
func (s *intStream) Sorted() *intStream {
	if s.err != nil {
		return s
	}
	na := []int{}
	if s.parallel {
		s.parallelEvaluate(func(i int) (bool, int) {
			na = append(na, i)
			return false, i
		})
	} else {
		s.evaluate(func(i int) (bool, int) {
			na = append(na, i)
			return false, i
		})
	}

	sort.Ints(na)
	s.array = na
	s.funcs = []func(int) (bool, int){}
	return s
}

func (s *intStream) FindFirst() (int, error) {
	if s.err != nil {
		return 0, s.err
	}

	if s.parallel {
		return s.parallelEvaluate(func(i int) (bool, int) {
			return true, i
		})
	}

	return s.evaluate(func(i int) (bool, int) {
		return true, i
	})
}

func newintStream(array []int) *intStream {
	return &intStream{
		array: array,
		funcs: []func(int) (bool, int){},
	}
}
