package stream

import "sort"

type LessAction[T any] func(T, T) bool

func (p *BaseStream[T]) Sort(action LessAction[T]) *BaseStream[T] {
	p.C = Sort(p.C, action)

	return p
}

func Sort[T any](input <-chan T, action LessAction[T]) chan T {
	output := make(chan T)

	go sortFun(input, output, action)

	return output
}

func sortFun[T any](input <-chan T, output chan<- T, action LessAction[T]) {
	elems := []T{}

	for i := range input {
		elems = append(elems, i)
	}

	sort.Slice(elems, func(i, j int) bool { return action(elems[i], elems[j]) })

	for _, elem := range elems {
		output <- elem
	}

	close(output)
}
