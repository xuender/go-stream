package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleBaseStream_Limit() {
	stream.NewBase(stream.Range2Channel(10)).
		Limit(3).
		ForEach(func(num int) {
			fmt.Println(num)
		})

	// Output:
	// 0
	// 1
	// 2
}
