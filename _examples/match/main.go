package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	base1 := stream.NewBase(stream.Range2Channel(10))
	base2 := stream.NewBase(stream.Copy(base1))

	fmt.Println(base1.AnyMatch(func(num int) bool { return num > 100 }))
	fmt.Println(base2.AnyMatch(func(num int) bool { return num > 1 }))
}
