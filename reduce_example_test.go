package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

// ExampleOrderedStream_Reduce is an example function.
func ExampleOrderedStream_Reduce() {
	input := make(chan int)
	order := stream.NewOrdered(input).
		Reduce()

	go func() {
		input <- 3
		input <- 2
		input <- 7
		input <- 1

		close(input)
	}()

	for elem := range order.C {
		fmt.Println(elem)
	}

	// Output:
	// 7
	// 3
	// 2
	// 1
}
