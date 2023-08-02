package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	base := stream.NewBase(stream.Range2Channel(1, 100)).
		Filter(func(num int) bool { return num%7 == 0 })

	for num := range base.C {
		fmt.Println(num)
	}
}
