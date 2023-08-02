package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleParallelStream_Sequential() {
	strea := stream.NewParallel(stream.Range2Channel(1, 10), 10).
		Sequential().
		Filter(func(num int) bool { return num > 5 })

	fmt.Println(strea.FindFirst())

	// Output:
	// 6
}
