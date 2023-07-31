package main

import (
	"time"

	"github.com/xuender/go-stream"
)

func main() {
	input := make(chan int)
	base := stream.NewBase(input)

	base.Filter(func(num int) bool { return num > 5 })

	go func(cha chan<- int) {
		for i := 0; i < 10; i++ {
			cha <- i
		}

		close(cha)
	}(input)

	time.Sleep(time.Millisecond)
	println(base.Count())
}
