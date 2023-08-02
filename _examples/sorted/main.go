package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	stream.NewOrdered(stream.Slice2Channel(3, 2, 7, 1)).
		Sorted().
		ForEach(func(num int) {
			fmt.Println(num)
		})
}
