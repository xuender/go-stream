package stream

import "fmt"

func ExampleStream_ForEach() {
	arr := []int{1, 2, 3}

	err := New(arr).
		Filter(func(i int) bool { return i > 1 }).
		ForEach(func(i int) { fmt.Println(i) })

	fmt.Println(err)

	// Output:
	// 2
	// 3
	// <nil>
}
