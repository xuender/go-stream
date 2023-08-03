package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xuender/go-stream"
)

func main() {
	stream.NewBase(stream.Range2Channel(1000)).
		Parallel(100).
		Filter(func(num int) bool { return num%7 == 0 }).
		ForEach(func(num int) {
			dur := time.Duration(rand.Intn(1000)) * time.Millisecond

			time.Sleep(dur)
			fmt.Printf("%d\t%s\n", num, dur)
		})
}
