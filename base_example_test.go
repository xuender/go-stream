package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

// ExampleNewBase is an example function.
func ExampleNewBase() {
	input := make(chan int)
	base := stream.NewBase(input).
		Filter(func(num int) bool { return num > 5 })

	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}

		close(input)
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(base.Count())

	// Output:
	// 4
}

// ExampleBaseStream_FindFirst is an example function.
func ExampleBaseStream_FindFirst() {
	input := make(chan int)
	base := stream.NewBase(input).
		Filter(func(num int) bool { return num > 5 })

	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}

		close(input)
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(base.FindFirst())

	// Output:
	// 6
}
