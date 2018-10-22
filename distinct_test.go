package stream

import (
	"fmt"
)

func ExampleStream_Distinct() {
	c, err := New([]int{1, 2, 3, 3, 2}).
		Distinct().
		Count()

	fmt.Println(c, err)

	// Output:
	// 3 <nil>
}
