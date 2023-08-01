package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

// ExampleFlatMap is an example function.
func ExampleFlatMap() {
	input := make(chan []int)
	defer close(input)

	base := stream.NewBase(input)
	mapBase := stream.FlatMap(base.C, func(num int) string { return fmt.Sprintf("[%d]", num) })

	go func() {
		for i := range mapBase.C {
			fmt.Println(i)
		}
	}()

	for i := 0; i < 3; i++ {
		input <- []int{i, i * 2}
	}

	time.Sleep(time.Millisecond)

	// Output:
	// [0]
	// [0]
	// [1]
	// [2]
	// [2]
	// [4]
}

// ExampleFlatMapComparable is an example function.
func ExampleFlatMapComparable() {
	input := make(chan []int)
	ordered := stream.FlatMapOrdered(input, func(num int) string { return fmt.Sprintf("[%d]", num) })

	go func() {
		for i := 1; i < 4; i++ {
			input <- []int{i, i * 2}
		}

		close(input)
	}()

	time.Sleep(time.Millisecond)
	// fmt.Println(ordered.Max())
	for i := range ordered.C {
		fmt.Println(i)
	}

	// Output:
	// [1]
	// [2]
	// [2]
	// [4]
	// [3]
	// [6]
}

// ExampleFlatMapOrdered is an example function.
func ExampleFlatMapOrdered() {
	input := make(chan []int)
	com := stream.FlatMapComparable(input, func(num int) string { return fmt.Sprintf("[%d]", num) }).
		Distinct()

	go func() {
		input <- []int{1, 2}
		input <- []int{1, 2}
		input <- []int{2, 4}
		input <- []int{3, 6}
		input <- []int{3, 6}
		input <- []int{4, 8}

		close(input)
	}()

	time.Sleep(time.Millisecond)

	for i := range com.C {
		fmt.Println(i)
	}

	// Output:
	// [1]
	// [2]
	// [4]
	// [3]
	// [6]
	// [8]
}
