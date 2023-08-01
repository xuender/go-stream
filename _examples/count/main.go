package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	input := make(chan int)
	base := stream.NewBase(input).
		Filter(func(num int) bool { return num > 5 })

	go func(cha chan<- int) {
		for i := 0; i < 10; i++ {
			cha <- i
		}

		close(cha)
	}(input)

	fmt.Println(base.Count())
}
