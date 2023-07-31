package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

// ExampleBaseStream_AnyMatch is an example function.
func ExampleBaseStream_AnyMatch() {
	input := make(chan int)
	base1 := stream.NewBase(input)
	base2 := stream.NewBase(stream.Copy(base1))

	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}

		close(input)
	}()

	fmt.Println(base1.AnyMatch(func(num int) bool { return num > 8 }))
	fmt.Println(base2.AnyMatch(func(num int) bool { return num > 10 }))

	// Output:
	// true
	// false
}
