package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

// ExampleMap is an example function.
func ExampleMap() {
	input := make(chan int)
	defer close(input)

	base := stream.NewBase(input)

	base.Filter(func(num int) bool { return num > 5 })
	newBase := stream.Map(base.C, func(num int) string { return fmt.Sprintf("[%d]", num) })

	go func() {
		for i := range newBase.C {
			fmt.Println(i)
		}
	}()

	for i := 0; i < 10; i++ {
		input <- i
	}

	time.Sleep(time.Millisecond)

	// Output:
	// [6]
	// [7]
	// [8]
	// [9]
}

// ExampleMapOrdered is an example function.
func ExampleMapOrdered() {
	input := make(chan int)
	ordered := stream.MapOrdered(input, func(num int) string { return fmt.Sprintf("[%d]", num) })

	go func() {
		for i := 0; i < 10; i++ {
			input <- i
		}

		close(input)
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(ordered.Max())

	// Output:
	// [9]
}
