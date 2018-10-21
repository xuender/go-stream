package stream

import "fmt"

func ExampleNewintStream() {
	arr := []int{7, 2, 3}
	h, _ := newintStream(arr).
		Filter(func(i int) bool { return i > 1 }).
		Map(func(i int) int { return i + 5 }).
		Sorted().
		FindFirst()
	fmt.Println(h)

	// Output:
	// 7
}
