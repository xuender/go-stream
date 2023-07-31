package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

// ExampleOrderedStream_Min is an example function.
func ExampleOrderedStream_Min() {
	input := make(chan int)
	ordered := stream.MapOrdered(input, func(num int) string { return fmt.Sprintf("[%d]", num) })

	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}

		close(input)
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(ordered.Min())

	// Output:
	// [0]
}
