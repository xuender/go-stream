package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

// ExampleBaseStream_ForEach is an example function.
func ExampleBaseStream_Peek() {
	input := make(chan int)
	count1 := 0
	count2 := 0
	base := stream.NewBase(input).
		Peek(func(num int) { count1++ }).
		Filter(func(num int) bool { return num%2 == 0 }).
		Peek(func(num int) { count2++ })

	go func() {
		for i := 0; i < 3; i++ {
			input <- i
		}

		close(input)
	}()

	fmt.Println(base.Count())
	fmt.Println(count1)
	fmt.Println(count2)

	// Output:
	// 2
	// 3
	// 2
}
