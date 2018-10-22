package stream

import (
	"fmt"
)

func ExampleStream_Sequential() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c, err := New(arr).
		Parallel().
		Filter(func(i int) bool { return i > 1 }).
		Sequential().
		Filter(func(i int) bool { return i > 3 }).
		Count()

	fmt.Println(c, err)

	// Output:
	// 6 <nil>
}
