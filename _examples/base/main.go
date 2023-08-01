package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	input := make(chan int)
	base := stream.NewBase(input).
		Peek(func(num int) { fmt.Println("peek1:", num) }).
		Filter(func(num int) bool { return num > 1 }).
		Peek(func(num int) { fmt.Println("peek2:", num) })

	go func() {
		for i := 1; i < 5; i++ {
			input <- i
		}

		close(input)
	}()

	fmt.Println(base.Count())
}
