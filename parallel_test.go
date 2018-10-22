package stream

import "fmt"

func ExampleStream_Parallel() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	NewStream(arr).
		Parallel().
		Filter(func(i int) bool { return i > 1 }).
		ForEach(func(i int) { fmt.Println(i) })
}
