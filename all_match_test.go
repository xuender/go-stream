package stream

import "fmt"

func ExampleStream_AllMatch() {
	all, err := New([]int{1, 2, 3, 4}).
		AllMatch(func(i int) bool { return i > 2 })
	fmt.Println(all, err)

	all, err = New([]int{3, 4}).
		AllMatch(func(i int) bool { return i > 2 })
	fmt.Println(all, err)

	// Output:
	// false <nil>
	// true <nil>
}
