package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleBaseStream_Sort() {
	stream.NewBase(stream.Slice2Channel(1, 3, 2, 7, 1)).
		Sort(func(num1, num2 int) bool { return num2 < num1 }).
		ForEach(func(num int) {
			fmt.Println(num)
		})

	// Output:
	// 7
	// 3
	// 2
	// 1
}
