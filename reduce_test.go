package stream

import (
	"fmt"
)

func ExampleStream_Reduce() {
	sum, err := New([]int{1, 2, 3, 4, 1}).
		Filter(func(i int) bool { return i > 1 }).
		Map(func(i int) int { return i + 100 }).
		Reduce(func(i, j int) int { return i + j })

	fmt.Println(sum, err)

	// Output:
	// 309 <nil>
}
