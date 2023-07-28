package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

// ExampleNew is an example function.
func ExampleNew() {
	array := stream.New(1, 2, 3)
	fmt.Println(array.Count())

	// Output:
	// 3
}
