package stream

import (
	"fmt"
)

func ExampleStream_Reduce() {
	sum, err := New([]int{1, 2, 3, 4}).
		Filter(func(i int) bool { return i > 1 }).
		Reduce(func(i, j int) int { return i + j })

	fmt.Println(sum, err)

	// Output:
	// 9 <nil>
}
