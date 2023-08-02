package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleOrderedStream_Reduce() {
	stream.NewOrdered(stream.Slice2Channel(3, 2, 7, 1)).
		Reduce().
		ForEach(func(num int) {
			fmt.Println(num)
		})

	// Output:
	// 7
	// 3
	// 2
	// 1
}
