package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	base := stream.Map(
		stream.Range2Channel(1, 100),
		func(num int) string {
			return fmt.Sprintf("[%d]", num)
		},
	).Limit(3)

	for i := range base.C {
		fmt.Println(i)
	}
}
