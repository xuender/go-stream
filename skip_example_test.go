package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleBaseStream_Skip() {
	stream.NewBase(stream.Range2Channel(1, 6)).
		Skip(4).
		ForEach(func(num int) {
			fmt.Println(num)
		})

	// Output:
	// 4
	// 5
}
