package stream

import (
	"fmt"
)

func ExampleStream_Peek() {
	arr := []int{1, 2, 3}
	c, err := New(arr).
		Peek(func(i int) { fmt.Println(i) }).
		Count()
	fmt.Println(c, err)

	// Output:
	// 1
	// 2
	// 3
	// 3 <nil>
}
