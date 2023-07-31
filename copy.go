package stream

const _hundred = 100

func Copy[T any](source *BaseStream[T]) chan T {
	output1 := make(chan T, _hundred)
	output2 := make(chan T, _hundred)

	go func(input <-chan T) {
		for i := range input {
			output1 <- i
			output2 <- i
		}

		close(output1)
		close(output2)
	}(source.C)

	source.C = output1

	return output2
}
