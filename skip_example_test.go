package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

// ExampleBaseStream_Skip is an example function.
func ExampleBaseStream_Skip() {
	input := make(chan int)
	defer close(input)

	base := stream.NewBase(input).
		Skip(4)

	go func() {
		for i := range base.C {
			fmt.Println(i)
		}
	}()

	for i := 0; i < 6; i++ {
		input <- i
	}

	time.Sleep(time.Millisecond)

	// Output:
	// 4
	// 5
}
