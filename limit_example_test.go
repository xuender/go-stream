package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleBaseStream_Limit() {
	stream.NewBase(stream.Range2Channel(1, 10)).
		Limit(3).
		ForEach(func(num int) {
			fmt.Println(num)
		})

	// Output:
	// 0
	// 1
	// 2
}

func ExampleLimit() {
	for num := range stream.Limit(stream.Range2Channel(1, 10), 0) {
		fmt.Println(num)
	}

	// Output:
}

func ExampleLimit_out() {
	for num := range stream.Limit(stream.Range2Channel(1, 2), 30) {
		fmt.Println(num)
	}

	// Output:
	// 0
	// 1
}
