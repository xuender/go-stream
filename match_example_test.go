package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleBaseStream_AnyMatch() {
	base1 := stream.NewBase(stream.Range2Channel(1, 10))
	base2 := stream.NewBase(stream.Copy(base1))

	fmt.Println(base1.AnyMatch(func(num int) bool { return num > 8 }))
	fmt.Println(base2.AnyMatch(func(num int) bool { return num > 10 }))

	// Output:
	// true
	// false
}

func ExampleBaseStream_AllMatch() {
	base1 := stream.NewBase(stream.Range2Channel(1, 10))
	base2 := stream.NewBase(stream.Copy(base1))

	fmt.Println(base1.AllMatch(func(num int) bool { return num > 8 }))
	fmt.Println(base2.AllMatch(func(num int) bool { return num >= 0 }))

	// Output:
	// false
	// true
}

func ExampleBaseStream_NoneMatch() {
	base1 := stream.NewBase(stream.Range2Channel(1, 10))
	base2 := stream.NewBase(stream.Copy(base1))

	fmt.Println(base1.NoneMatch(func(num int) bool { return num > 10 }))
	fmt.Println(base2.NoneMatch(func(num int) bool { return num > 8 }))

	// Output:
	// true
	// false
}
