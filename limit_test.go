package stream

import (
	"fmt"
)

func ExampleStream_Limit() {
	err := New([]int{1, 2, 3}).
		Limit(2).
		ForEach(func(i int) { fmt.Println(i) })

	fmt.Println(err)

	// Output:
	// 1
	// 2
	// <nil>
}
