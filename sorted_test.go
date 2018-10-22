package stream

import (
	"fmt"
)

func ExampleStream_Sorted() {
	err := New([]int{3, 1, 4, 2}).
		Sorted(func(i, j int) bool { return i < j }).
		ForEach(func(i int) { fmt.Println(i) })

	fmt.Println(err)

	// Output:
	// 1
	// 2
	// 3
	// 4
	// <nil>
}
