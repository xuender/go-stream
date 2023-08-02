package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleBaseStream_Peek() {
	count1 := 0
	count2 := 0
	base := stream.NewBase(stream.Range2Channel(1, 3)).
		Peek(func(num int) { count1++ }).
		Filter(func(num int) bool { return num%2 == 0 }).
		Peek(func(num int) { count2++ })

	fmt.Println(base.Count())
	fmt.Println(count1)
	fmt.Println(count2)

	// Output:
	// 2
	// 3
	// 2
}
