package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	base := stream.NewBase(stream.Range2Channel(100)).
		Peek(func(i int) { println("P1:", i) }).
		Filter(func(t int) bool { return t%5 == 0 }).
		Peek(func(i int) { println("P2:", i) })

	for num := range base.C {
		fmt.Println(num)
	}
}
