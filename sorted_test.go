package stream

import (
	"fmt"
)

func ExampleStream_Sorted() {
	err := New([]int{3, 1, 4, 2, 5, 6}).
		Filter(func(i int) bool { return i < 6 }).
		Filter(func(i int) bool { return i > 1 }).
		Sorted(func(i, j int) bool { return i < j }).
		ForEach(func(i int) { fmt.Println(i) })

	fmt.Println(err)

	// Output:
	// 2
	// 3
	// 4
	// 5
	// <nil>
}
