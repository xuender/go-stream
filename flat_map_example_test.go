package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleFlatMap() {
	stream.FlatMap(
		stream.Slice2Channel([]int{0, 0}, []int{1, 2}, []int{2, 4}),
		func(num int) string { return fmt.Sprintf("[%d]", num) },
	).ForEach(func(str string) {
		fmt.Println(str)
	})

	// Output:
	// [0]
	// [0]
	// [1]
	// [2]
	// [2]
	// [4]
}

func ExampleFlatMapComparable() {
	stream.FlatMapOrdered(
		stream.Slice2Channel([]int{1, 2}, []int{2, 4}, []int{3, 6}),
		func(num int) string { return fmt.Sprintf("[%d]", num) },
	).
		Sorted().
		ForEach(func(str string) {
			fmt.Println(str)
		})

	// Output:
	// [1]
	// [2]
	// [2]
	// [3]
	// [4]
	// [6]
}

func ExampleFlatMapOrdered() {
	stream.FlatMapComparable(stream.Slice2Channel([]int{1, 2}, []int{1, 2}, []int{2, 4},
		[]int{3, 6}, []int{3, 6}, []int{4, 8}),
		func(num int) string { return fmt.Sprintf("[%d]", num) },
	).
		Distinct().
		ForEach(func(str string) {
			fmt.Println(str)
		})

	// Output:
	// [1]
	// [2]
	// [4]
	// [3]
	// [6]
	// [8]
}
