package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

// ExampleBaseStream_Sort is an example function.
func ExampleBaseStream_Sort() {
	input := make(chan int)
	order := stream.NewBase(input).
		Sort(func(num1, num2 int) bool { return num2 < num1 })

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
