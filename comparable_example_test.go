package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

// ExampleComparableStream_Distinct is an example function.
func ExampleComparableStream_Distinct() {
	input := make(chan int)
	defer close(input)

	chs := stream.NewComparable(input).
		Distinct()

	go func() {
		for i := range chs.C {
			fmt.Println(i)
		}
	}()

	input <- 1
	input <- 1
	input <- 2
	input <- 3
	input <- 3
	input <- 4

	time.Sleep(time.Millisecond)

	// Output:
	// 1
	// 2
	// 3
	// 4
}
