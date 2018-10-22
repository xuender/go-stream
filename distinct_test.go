package stream

import (
	"fmt"
)

func ExampleStream_Distinct() {
	c, err := New([]int{1, 2, 3, 3, 2}).
		Filter(func(i int) bool { return i > 1 }).
		Distinct().
		Count()

	fmt.Println(c, err)

	// Output:
	// 2 <nil>
}
