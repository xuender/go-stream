package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	input := make(chan int)
	array := stream.Map(input, func(num int) string {
		return fmt.Sprintf("[%d]", num)
	}).Limit(3)

	go func(cha chan<- int) {
		for i := 0; i < 100; i++ {
			cha <- i
		}

		close(cha)
	}(input)

	for i := range array.C {
		fmt.Println(i)
	}
}
