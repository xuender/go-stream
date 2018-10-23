package stream

import "fmt"

func ExampleStream_Map() {
	arr := []int{1, 2, 3}

	err := New(arr).
		Map(func(i int) bool { return i > 1 }).
		ForEach(func(b bool) { fmt.Println(b) })

	fmt.Println(err)

	// Output:
	// false
	// true
	// true
	// <nil>
}
