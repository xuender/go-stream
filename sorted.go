package stream

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func (p *OrderedStream[O]) Sorted() *OrderedStream[O] {
	p.C = Sorted(p.C)

	return p
}

func Sorted[O constraints.Ordered](input <-chan O) chan O {
	output := make(chan O)

	go sorted(input, output)

	return output
}

func sorted[O constraints.Ordered](input <-chan O, output chan<- O) {
	elems := []O{}

	for i := range input {
		elems = append(elems, i)
	}

	sort.Slice(elems, func(i, j int) bool { return elems[i] < elems[j] })

	for _, elem := range elems {
		output <- elem
	}

	close(output)
}
