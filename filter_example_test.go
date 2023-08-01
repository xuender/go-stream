package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

// ExampleBaseStream_Filter is an example function.
func ExampleBaseStream_Filter() {
	input := make(chan int)
	defer close(input)

	chs := stream.NewBase(input).
		Filter(func(num int) bool { return num > 5 })

	go func() {
		for i := range chs.C {
			fmt.Println(i)
		}
	}()

	for i := 0; i < 10; i++ {
		input <- i
	}

	time.Sleep(time.Millisecond)

	// Output:
	// 6
	// 7
	// 8
	// 9
}

// ExampleFilter is an example function.
func ExampleFilter() {
	input := make(chan int)
	defer close(input)

	chi := stream.Filter(input, func(num int) bool { return num > 5 })

	go func() {
		for i := range chi {
			fmt.Println(i)
		}
	}()

	for i := 0; i < 10; i++ {
		input <- i
	}

	time.Sleep(time.Millisecond)

	// Output:
	// 6
	// 7
	// 8
	// 9
}

// ExampleParallelStream_Filter is an example function.
func ExampleParallelStream_Filter() {
	input := make(chan int)
	defer close(input)

	chs := stream.NewBase(input).
		Parallel(3).
		Filter(func(num int) bool { return num > 5 })

	go func() {
		for range chs.C {
			fmt.Println("x")
		}
	}()

	for i := 0; i < 10; i++ {
		input <- i
	}

	time.Sleep(time.Millisecond)

	// Output:
	// x
	// x
	// x
	// x
}
