package stream

func Count[T any](input <-chan T) int {
	count := 0

	for range input {
		count++
	}

	return count
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

func Range2Channel(length int) chan int {
	output := make(chan int)
	step := 1

	if length < 0 {
		step = -1
	}

	go range2Channel(output, 0, length, step)

	return output
}

func RangeFrom2Channel(start, length int) chan int {
	output := make(chan int)
	step := 1

	if length < 0 {
		step = -1
	}

	length += start

	go range2Channel(output, start, length, step)

	return output
}

func RangeWithSteps2Channel(start, end, step int) chan int {
	output := make(chan int)

	go range2Channel(output, start, end, step)

	return output
}

func Slice2Channel[T any](elems ...T) chan T {
	output := make(chan T)

	go slice2Channel(output, elems)

	return output
}

func range2Channel(output chan<- int, start, end, step int) {
	for i := start; (step > 0 && i < end) || (step < 0 && i > end); i += step {
		output <- i
	}

	close(output)
}

func slice2Channel[T any](output chan<- T, elems []T) {
	for _, elem := range elems {
		output <- elem
	}

	close(output)
}
