package main

import (
	"github.com/xuender/go-stream"
)

func main() {
	input := make(chan int)
	base := stream.NewBase(input)

	base.Filter(func(t int) bool { return t%5 == 0 })

	go func(cha chan<- int) {
		for i := 0; i < 100; i++ {
			cha <- i
		}

		close(cha)
	}(input)

	for i := range base.C {
		println(i)
	}
}
