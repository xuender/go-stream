package stream

import "golang.org/x/exp/constraints"

func Map[I, O any](input chan I, predicate func(I) O) *BaseStream[O] {
	output := make(chan O)
	ret := NewBase[O](output)

	go func() {
		for i := range input {
			output <- predicate(i)
		}

		close(output)
	}()

	return ret
}

func MapOrdered[I any, O constraints.Ordered](input chan I, predicate func(I) O) *OrderedStream[O] {
	output := make(chan O)
	ret := NewOrdered(output)

	go func() {
		for i := range input {
			output <- predicate(i)
		}

		close(output)
	}()

	return ret
}
