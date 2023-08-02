package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	base := stream.NewBase(stream.Range2Channel(10)).
		Filter(func(num int) bool { return num > 5 })

	fmt.Println(base.Count())
}
