package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleOrderedStream_Sorted() {
	stream.NewOrdered(stream.Slice2Channel(3, 2, 7, 1)).
		Sorted().ForEach(func(num int) {
		fmt.Println(num)
	})

	// Output:
	// 1
	// 2
	// 3
	// 7
}
