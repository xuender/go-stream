package stream

import "fmt"

func ExampleNewStream() {
	arr := []int{1, 2, 3, 4, 5}
	i, err := NewStream(arr).
		Peek(func(i int) { fmt.Println("peek1:", i) }).
		Filter(func(i int) bool { return i > 1 }).
		Peek(func(i int) { fmt.Println("peek2:", i) }).
		Map(func(i int) string { return fmt.Sprintf("id:%d", i) }).
		Peek(func(s string) { fmt.Println("peek3:", s) }).
		FindFirst()
	fmt.Println(i, err)

	// Output:
	// peek1: 1
	// peek1: 2
	// peek2: 2
	// peek3: id:2
	// id:2 <nil>
}
