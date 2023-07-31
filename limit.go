package stream

func (p *BaseStream[T]) Limit(size int) *BaseStream[T] {
	p.C = Limit(p.C, size)

	return p
}

func Consume[T any](input <-chan T) {
	// nolint: revive
	for range input {
	}
}

func Limit[T any](input <-chan T, size int) chan T {
	output := make(chan T)

	if size <= 0 {
		go Consume(input)
		close(output)

		return output
	}

	go func() {
		count := 0

		for i := range input {
			output <- i

			count++
			if count >= size {
				break
			}
		}

		go Consume(input)
		close(output)
	}()

	return output
}
