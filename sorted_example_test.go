package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

// ExampleOrderedStream_Sorted is an example function.
func ExampleOrderedStream_Sorted() {
	input := make(chan int)
	order := stream.NewOrdered(input).
		Sorted()

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
	// 1
	// 2
	// 3
	// 7
}
