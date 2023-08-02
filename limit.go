package stream

// Limit returns a stream consisting of the elements of this stream,
// truncated to be no longer than maxSize in length.
func (p *BaseStream[T]) Limit(maxSize int) *BaseStream[T] {
	p.C = Limit(p.C, maxSize)

	return p
}

// Limit returns a channel consisting of the elements of input channel,
// truncated to be no longer than maxSize in length.
func Limit[T any](input <-chan T, maxSize int) chan T {
	output := make(chan T)

	if maxSize <= 0 {
		go Count(input)
		close(output)

		return output
	}

	go func() {
		count := 0

		for i := range input {
			output <- i

			count++
			if count >= maxSize {
				break
			}
		}

		go Count(input)
		close(output)
	}()

	return output
}
