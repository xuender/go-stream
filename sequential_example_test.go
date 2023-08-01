package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

// ExampleParallelStream_Sequential is an example function.
func ExampleParallelStream_Sequential() {
	input := make(chan int)
	strea := stream.NewParallel(input, 10).
		Sequential().
		Filter(func(num int) bool { return num > 5 })

	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}

		close(input)
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(strea.FindFirst())

	// Output:
	// 6
}
