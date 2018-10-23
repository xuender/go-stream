package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	f, err := stream.New(arr).
		Peek(func(i int) { fmt.Println("peek1:", i) }).
		Filter(func(i int) bool { return i > 1 }).
		Peek(func(i int) { fmt.Println("peek2:", i) }).
		Map(func(i int) string { return fmt.Sprintf("id:%d", i) }).
		Peek(func(s string) { fmt.Println("peek3:", s) }).
		FindFirst()

	fmt.Println(f, err)
}
