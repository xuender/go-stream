package stream

import "fmt"

func ExampleStream_FindFirst() {
	arr := []int{1, 2, 3}
	i, err := New(arr).
		Filter(func(i int) bool { return i > 1 }).
		FindFirst()
	fmt.Println(i, err)

	// Output:
	// 2 <nil>
}
