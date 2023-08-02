package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleNewBase() {
	base := stream.NewBase(stream.Range2Channel(10)).
		Filter(func(num int) bool { return num > 5 })

	fmt.Println(base.Count())

	// Output:
	// 4
}

func ExampleBaseStream_FindFirst() {
	base := stream.NewBase(stream.Range2Channel(10)).
		Filter(func(num int) bool { return num > 5 })

	fmt.Println(base.FindFirst())

	// Output:
	// 6
}
