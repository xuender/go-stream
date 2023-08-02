package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

// ExampleSlice2Channel is an example function.
func ExampleSlice2Channel() {
	for i := range stream.Slice2Channel(9, 1, 2, 3) {
		fmt.Println(i)
	}

	// Output:
	// 1
	// 2
	// 3
}

// ExampleRange2Channel is an example function.
func ExampleRange2Channel() {
	for i := range stream.Range2Channel(1, 3) {
		fmt.Println(i)
	}

	for i := range stream.Range2Channel(1, -3) {
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

// ExampleRangeFrom2Channel is an example function.
func ExampleRangeFrom2Channel() {
	for i := range stream.RangeFrom2Channel(1, 3, 2) {
		fmt.Println(i)
	}

	for i := range stream.RangeFrom2Channel(1, 8, -3) {
		fmt.Println(i)
	}

	// Output:
	// 3
	// 4
	// 8
	// 7
	// 6
}

// ExampleRangeWithSteps2Channel is an example function.
func ExampleRangeWithSteps2Channel() {
	for i := range stream.RangeWithSteps2Channel(1, 0, 4, 3) {
		fmt.Println(i)
	}

	for i := range stream.RangeWithSteps2Channel(1, 9, 5, -3) {
		fmt.Println(i)
	}

	// Output:
	// 0
	// 3
	// 9
	// 6
}
