package stream

import "fmt"

func ExampleStream_AnyMatch() {
	b, err := New([]int{1, 2, 3, 4}).
		AnyMatch(func(i int) bool { return i > 2 })
	fmt.Println(b, err)

	b, err = New([]int{1, 0}).
		AnyMatch(func(i int) bool { return i > 2 })
	fmt.Println(b, err)

	// Output:
	// true <nil>
	// false <nil>
}
