package main

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	f, err := stream.New(arr).
		Parallel().
		Peek(func(i int) {
			fmt.Println("peek1:", i)
			time.Sleep(time.Second * time.Duration(i))
		}).
		Filter(func(i int) bool { return i > 1 }).
		Peek(func(i int) {
			fmt.Println("peek2:", i)
			time.Sleep(time.Second * time.Duration(i))
		}).
		FindFirst()

	fmt.Println(f, err)
}
