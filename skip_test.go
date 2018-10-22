package stream

import (
	"fmt"
)

func ExampleStream_Skip() {
	err := New([]int{1, 2, 3, 4}).
		Skip(2).
		ForEach(func(i int) { fmt.Println(i) })

	fmt.Println(err)

	// Output:
	// 3
	// 4
	// <nil>
}
