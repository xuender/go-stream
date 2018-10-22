package stream

import "fmt"

func ExampleStream_NoneMatch() {
	all, err := New([]int{1, 2, 3, 4}).
		NoneMatch(func(i int) bool { return i > 2 })
	fmt.Println(all, err)

	all, err = New([]int{3, 4}).
		NoneMatch(func(i int) bool { return i > 4 })
	fmt.Println(all, err)

	// Output:
	// false <nil>
	// true <nil>
}
