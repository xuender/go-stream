package stream

import (
	"fmt"
)

func ExampleStream_Min() {
	max, err := New([]int{3, 1, 4, 2}).
		Min(func(i, j int) bool { return i < j })

	fmt.Println(max, err)

	// Output:
	// 1 <nil>
}
