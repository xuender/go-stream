package stream

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
	for i := range input {
		action(i)
		output <- i
	}

	close(output)
}
