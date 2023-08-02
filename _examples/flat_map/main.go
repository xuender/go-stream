package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	stream.FlatMap(
		stream.Slice2Channel(1, []int{0, 0}, []int{1, 2}, []int{2, 4}),
		func(num int) string { return fmt.Sprintf("[%d]", num) },
	).ForEach(func(str string) {
		fmt.Println(str)
	})
}
