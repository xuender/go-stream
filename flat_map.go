package stream

import "golang.org/x/exp/constraints"

func FlatMap[I, O any](input chan []I, action MapAction[I, O]) *BaseStream[O] {
	output := make(chan O)

	go flatMap(input, output, action)

	return NewBase[O](output)
}

func FlatMapComparable[I any, O comparable](input chan []I, action MapAction[I, O]) *ComparableStream[O] {
	output := make(chan O)

	go flatMap(input, output, action)

	return NewComparable(output)
}

func FlatMapOrdered[I any, O constraints.Ordered](input chan []I, action MapAction[I, O]) *OrderedStream[O] {
	output := make(chan O)

	go flatMap(input, output, action)

	return NewOrdered(output)
}

func flatMap[I, O any](input <-chan []I, output chan<- O, action MapAction[I, O]) {
	for elems := range input {
		for _, elem := range elems {
			output <- action(elem)
		}
	}

	close(output)
}
