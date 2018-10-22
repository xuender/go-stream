package stream

import "fmt"

func ExampleStream_Count() {
	arr := []int{1, 2, 3, 4, 5}

	i, err := New(arr).
		Parallel().
		Filter(func(i int) bool { return i > 1 }).
		Count()

	fmt.Println(i, err)

	// Output:
	// 4 <nil>
}
