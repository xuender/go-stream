package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleSlice2Channel() {
	for i := range stream.Slice2Channel(1, 2, 3) {
		fmt.Println(i)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleRange2Channel() {
	for i := range stream.Range2Channel(3) {
		fmt.Println(i)
	}

	for i := range stream.Range2Channel(-3) {
		fmt.Println(i)
	}

	// Output:
	// 0
	// 1
	// 2
	// 0
	// -1
	// -2
}

func ExampleRangeFrom2Channel() {
	for i := range stream.RangeFrom2Channel(3, 2) {
		fmt.Println(i)
	}

	for i := range stream.RangeFrom2Channel(8, -3) {
		fmt.Println(i)
	}

	// Output:
	// 3
	// 4
	// 8
	// 7
	// 6
}

func ExampleRangeWithSteps2Channel() {
	for i := range stream.RangeWithSteps2Channel(0, 4, 3) {
		fmt.Println(i)
	}

	for i := range stream.RangeWithSteps2Channel(9, 5, -3) {
		fmt.Println(i)
	}

	// Output:
	// 0
	// 3
	// 9
	// 6
}
