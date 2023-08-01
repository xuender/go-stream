package stream

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func (p *OrderedStream[O]) Reduce() *OrderedStream[O] {
	p.C = Reduce(p.C)

	return p
}

func Reduce[O constraints.Ordered](input <-chan O) chan O {
	output := make(chan O)

	go reduce(input, output)

	return output
}

func reduce[O constraints.Ordered](input <-chan O, output chan<- O) {
	elems := []O{}

	for i := range input {
		elems = append(elems, i)
	}

	sort.Slice(elems, func(i, j int) bool { return elems[j] < elems[i] })

	for _, elem := range elems {
		output <- elem
	}

	close(output)
}
