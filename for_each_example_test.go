package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

func ExampleBaseStream_ForEach() {
	stream.NewBase(stream.Range2Channel(1, 10)).
		ForEach(func(num int) { fmt.Println(num) })

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
}

func ExampleParallelStream_ForEach() {
	stream.NewParallel(stream.Range2Channel(1, 3), 3).
		ForEach(func(num int) {
			time.Sleep(time.Duration(3-num) * time.Millisecond)
			fmt.Println(num)
		})

	// Output:
	// 2
	// 1
	// 0
}
