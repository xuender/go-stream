package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	arr := []string{"go", "stream", "is", "good"}

	sum, err := stream.New(arr).
		FlatMap(func(s string) []byte { return []byte(s) }).
		Map(func(s byte) int { return int(s) }).
		Reduce(func(x, y int) int { return x + y })

	fmt.Println(sum, err)
}
