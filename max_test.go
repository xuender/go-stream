package stream

import (
	"fmt"
)

func ExampleStream_Max() {
	max, err := New([]int{3, 1, 4, 2}).
		Filter(func(i int) bool { return i < 4 }).
		Max(func(i, j int) bool { return i < j })

	fmt.Println(max, err)

	// Output:
	// 3 <nil>
}
