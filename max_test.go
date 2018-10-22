package stream

import (
	"fmt"
)

func ExampleStream_Max() {
	max, err := New([]int{3, 1, 4, 2}).
		Max(func(i, j int) bool { return i < j })

	fmt.Println(max, err)

	// Output:
	// 4 <nil>
}
