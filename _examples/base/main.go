package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	base := stream.NewBase(stream.Range2Channel(5)).
		Peek(func(num int) { fmt.Println("peek1:", num) }).
		Filter(func(num int) bool { return num > 2 }).
		Peek(func(num int) { fmt.Println("peek2:", num) })

	fmt.Println(base.Count())
}
