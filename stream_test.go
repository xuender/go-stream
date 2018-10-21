package stream

import "fmt"

func ExampleNewStream() {
	arr := []int{1, 2, 3}
	i, err := NewStream(arr).
		Peek(func(i int) { fmt.Println("peek1:", i) }).
		Filter(func(i int) bool { return i > 1 }).
		Peek(func(i int) { fmt.Println("peek2:", i) }).
		Map(func(i int) int { return i + 5 }).
		Peek(func(i int) { fmt.Println("peek3:", i) }).
		FindFirst()
	fmt.Println(i, err)

	// Output:
	// peek1: 1
	// peek1: 2
	// peek2: 2
	// peek3: 7
	// 7 <nil>
}
