package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleOrderedStream_Min() {
	ordered := stream.MapOrdered(stream.Range2Channel(10), func(num int) string { return fmt.Sprintf("[%d]", num) })

	fmt.Println(ordered.Min())

	// Output:
	// [0]
}
