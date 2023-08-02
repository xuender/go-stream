package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleBaseStream_Filter() {
	stream.NewBase(stream.Range2Channel(1, 10)).
		Filter(func(num int) bool { return num > 5 }).
		ForEach(func(num int) {
			fmt.Println(num)
		})

	// Output:
	// 6
	// 7
	// 8
	// 9
}

func ExampleFilter() {
	chi := stream.Filter(
		stream.Range2Channel(1, 10),
		func(num int) bool { return num > 5 },
	)

	for i := range chi {
		fmt.Println(i)
	}

	// Output:
	// 6
	// 7
	// 8
	// 9
}

func ExampleParallelStream_Filter() {
	stream.NewBase(stream.Range2Channel(1, 10)).
		Parallel(3).
		Filter(func(num int) bool { return num > 5 }).
		ForEach(func(_ int) {
			fmt.Println("x")
		})

	// Output:
	// x
	// x
	// x
	// x
}
