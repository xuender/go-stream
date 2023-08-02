package stream

func Slice2Channel[T any](size int, elems ...T) chan T {
	output := make(chan T, size)

	go slice2Channel(output, elems)

	return output
}

func slice2Channel[T any](output chan<- T, elems []T) {
	for _, elem := range elems {
		output <- elem
	}

	close(output)
}

func Range2Channel(size, length int) chan int {
	output := make(chan int, size)
	step := 1

	if length < 0 {
		step = -1
	}

	go range2Channel(output, 0, length, step)

	return output
}

func RangeFrom2Channel(size, start, length int) chan int {
	output := make(chan int, size)
	step := 1

	if length < 0 {
		step = -1
	}

	length += start

	go range2Channel(output, start, length, step)

	return output
}

func RangeWithSteps2Channel(size, start, end, step int) chan int {
	output := make(chan int, size)

	go range2Channel(output, start, end, step)

	return output
}

func range2Channel(output chan<- int, start, end, step int) {
	for i := start; (step > 0 && i < end) || (step < 0 && i > end); i += step {
		output <- i
	}

	close(output)
}

func Distribute[T any](input <-chan T, output1, output2 chan<- T) {
	for elem := range input {
		select {
		case output1 <- elem:
		case output2 <- elem:
		}
	}

	close(output1)
	close(output2)
}

func Count[T any](input <-chan T) int {
	count := 0

	for range input {
		count++
	}

	return count
}
