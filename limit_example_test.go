package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

// ExampleBaseStream_Limit is an example function.
func ExampleBaseStream_Limit() {
	input := make(chan int)
	defer close(input)

	chs := stream.NewBase(input)

	chs.Limit(3)

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
	// 0
	// 1
	// 2
}

// ExampleLimit is an example function.
func ExampleLimit() {
	input := make(chan int)
	defer close(input)

	chi := stream.Limit(input, 0)

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
}

func ExampleLimit_out() {
	input := make(chan int)
	defer close(input)

	chi := stream.Limit(input, 3)

	go func() {
		for i := range chi {
			fmt.Println(i)
		}
	}()

	for i := 0; i < 2; i++ {
		input <- i
	}

	time.Sleep(time.Millisecond)

	// Output:
	// 0
	// 1
}
