package stream

func (p *BaseStream[T]) Peek(action Action[T]) *BaseStream[T] {
	p.C = Peek(p.C, action)

	return p
}

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
