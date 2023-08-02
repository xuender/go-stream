package stream

import "sync"

// Peek returns a stream consisting of the elements of this stream, additionally
// performing the provided action on each element as elements are consumed
// from the resulting stream.
func (p *BaseStream[T]) Peek(action Action[T]) *BaseStream[T] {
	p.C = Peek(p.C, action)

	return p
}

// Peek returns a channel consisting of the elements of input channel, additionally
// performing the provided action on each element as elements are consumed
// from the resulting stream.
func Peek[T any](input <-chan T, action Action[T]) chan T {
	output := make(chan T)

	go peek(input, output, action)

	return output
}

func peek[T any](input <-chan T, output chan<- T, action Action[T]) {
	for elem := range input {
		action(elem)
		output <- elem
	}

	close(output)
}

func (p *ParallelStream[T]) Peek(action Action[T]) *ParallelStream[T] {
	p.C = PeekParallel(p.C, p.Size, action)

	return p
}

func PeekParallel[T any](input <-chan T, size int, action Action[T]) chan T {
	output := make(chan T)
	group := &sync.WaitGroup{}

	group.Add(size)

	for i := 0; i < size; i++ {
		go peekParallel(input, output, group, action)
	}

	go waitAndClose(output, group)

	return output
}

func peekParallel[T any](input <-chan T, output chan<- T, group *sync.WaitGroup, action Action[T]) {
	for elem := range input {
		action(elem)
		output <- elem
	}

	group.Done()
}
