package stream

import "sync"

type FilterAction[T any] func(T) bool

// Filter returns a stream consisting of the elements of this stream that match
// the given action.
func (p *BaseStream[T]) Filter(action FilterAction[T]) *BaseStream[T] {
	p.C = Filter(p.C, action)

	return p
}

// Filter returns a channel consisting of the elements of input channel that match
// the given action.
func Filter[T any](input <-chan T, action FilterAction[T]) chan T {
	output := make(chan T)

	go filter(input, output, action)

	return output
}

func filter[T any](input <-chan T, output chan<- T, action FilterAction[T]) {
	for elem := range input {
		if action(elem) {
			output <- elem
		}
	}

	close(output)
}

// Filter returns a stream consisting of the elements of this stream that match
// the given action, parallel.
func (p *ParallelStream[T]) Filter(action FilterAction[T]) *ParallelStream[T] {
	p.C = FilterParallel(p.C, p.Size, action)

	return p
}

// FilterParallel returns a channel consisting of the elements of input channel that match
// the given action, parallel.
func FilterParallel[T any](input <-chan T, size int, action FilterAction[T]) chan T {
	output := make(chan T)
	group := &sync.WaitGroup{}

	group.Add(size)

	for i := 0; i < size; i++ {
		go filterParallel(input, output, group, action)
	}

	go waitAndClose(output, group)

	return output
}

func waitAndClose[T any](output chan<- T, group *sync.WaitGroup) {
	group.Wait()
	close(output)
}

func filterParallel[T any](input <-chan T, output chan<- T, group *sync.WaitGroup, action FilterAction[T]) {
	for elem := range input {
		if action(elem) {
			output <- elem
		}
	}

	group.Done()
}
