package stream_test

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

func ExampleBaseStream_Peek() {
	count1 := 0
	count2 := 0
	base := stream.NewBase(stream.Range2Channel(3)).
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

func ExampleParallelStream_Peek() {
	count := stream.NewParallel(stream.Range2Channel(3), 3).
		Peek(func(num int) {
			time.Sleep(time.Duration((3-num)*100) * time.Millisecond)
			fmt.Println(num)
		}).Count()

	fmt.Println(count)

	// Output:
	// 2
	// 1
	// 0
	// 3
}
