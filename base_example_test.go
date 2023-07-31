package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

// ExampleNewBase is an example function.
func ExampleNewBase() {
	input := make(chan int)

	chs := stream.NewBase(input)

	chs.Filter(func(num int) bool { return num > 5 })

	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}

		close(input)
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(chs.Count())

	// Output:
	// 4
}
