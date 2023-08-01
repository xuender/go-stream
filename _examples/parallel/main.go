package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xuender/go-stream"
)

func main() {
	input := make(chan int)
	parallel := stream.NewBase(input).
		Parallel(100).
		Filter(func(t int) bool { return t%7 == 0 })

	go func() {
		for i := 0; i < 1000; i++ {
			input <- i
		}

		close(input)
	}()

	parallel.ForEach(func(num int) {
		dur := time.Duration(rand.Intn(1000))

		time.Sleep(time.Millisecond * dur)
		fmt.Printf("%d\t%dms\n", num, dur)
	})
}
