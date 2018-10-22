package stream

import (
	"fmt"
	"time"
)

func Example_newintStream() {
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
func Example_intStreamSorted() {
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

func Example_intStreamForEach() {
	arr := []int{1, 2, 3}
	newintStream(arr).
		Filter(func(i int) bool { return i > 1 }).
		ForEach(func(i int) { fmt.Println(i) })

	// Output:
	// 2
	// 3
}

func Example_intStreamParallel() {
	arr := []int{1, 2, 3}
	newintStream(arr).
		Filter(func(i int) bool { return i > 1 }).
		ForEach(func(i int) { fmt.Println(i) })

	// Output:
	// 2
	// 3
}

func Example_intStreamFindFirst() {
	arr := []int{1, 2, 3, 4}
	n, err := newintStream(arr).
		Parallel().
		Peek(func(i int) {
			time.Sleep(time.Second * time.Duration(i))
			fmt.Println("seep1:", i)
		}).
		Filter(func(i int) bool { return i > 0 }).
		Peek(func(i int) {
			time.Sleep(time.Second * time.Duration(i))
			fmt.Println("seep2:", i)
		}).
		FindFirst()

	fmt.Println(n, err)
}
