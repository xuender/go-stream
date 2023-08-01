package stream

func (p *BaseStream[T]) Skip(size int) *BaseStream[T] {
	p.C = Skip(p.C, size)

	return p
}

func Skip[T any](input <-chan T, size int) chan T {
	output := make(chan T)

	go sikp(input, output, size)

	return output
}

func sikp[T any](input <-chan T, output chan<- T, size int) {
	count := 0

	for elem := range input {
		if count < size {
			count++

			continue
		}

		output <- elem
	}

	close(output)
}
