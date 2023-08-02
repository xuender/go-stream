package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleComparableStream_Distinct() {
	stream.NewComparable(stream.Slice2Channel(1, 1, 2, 3, 3, 4)).
		Distinct().
		ForEach(func(num int) {
			fmt.Println(num)
		})

	// Output:
	// 1
	// 2
	// 3
	// 4
}
