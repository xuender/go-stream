package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

// ExampleBaseStream_ForEach is an example function.
func ExampleBaseStream_ForEach() {
	input := make(chan int)
	chs := stream.NewBase(input)

	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}

		close(input)
	}()

	chs.ForEach(func(num int) { fmt.Println(num) })

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
}

// ExampleParallelStream_ForEach is an example function.
func ExampleParallelStream_ForEach() {
	input := make(chan int)
	chs := stream.NewParallel(input, 3)

	go func() {
		for i := 0; i < 3; i++ {
			input <- i
		}

		close(input)
	}()

	chs.ForEach(func(num int) {
		time.Sleep(time.Duration(3-num) * time.Millisecond)
		fmt.Println(num)
	})

	// Output:
	// 2
	// 1
	// 0
}
