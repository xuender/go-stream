package stream

import "golang.org/x/exp/constraints"

type MapAction[I, O any] func(I) O

func Map[I, O any](input chan I, action MapAction[I, O]) *BaseStream[O] {
	output := make(chan O)

	go mapRun(input, output, action)

	return NewBase[O](output)
}

func MapComparable[I any, O comparable](input chan I, action MapAction[I, O]) *ComparableStream[O] {
	output := make(chan O)

	go mapRun(input, output, action)

	return NewComparable(output)
}

func MapOrdered[I any, O constraints.Ordered](input chan I, action MapAction[I, O]) *OrderedStream[O] {
	output := make(chan O)

	go mapRun(input, output, action)

	return NewOrdered(output)
}

func mapRun[I, O any](input <-chan I, output chan<- O, action MapAction[I, O]) {
	for elem := range input {
		output <- action(elem)
	}

	close(output)
}
