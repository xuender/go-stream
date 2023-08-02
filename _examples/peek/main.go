package main

import "github.com/xuender/go-stream"

func main() {
	base := stream.NewBase(stream.Range2Channel(1, 100)).
		Peek(func(i int) { println("P1:", i) }).
		Filter(func(t int) bool { return t%5 == 0 }).
		Peek(func(i int) { println("P2:", i) })

	for i := range base.C {
		println(i)
	}
}
