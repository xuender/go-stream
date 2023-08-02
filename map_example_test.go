package stream_test

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func ExampleMap() {
	base1 := stream.NewBase(stream.Range2Channel(10)).
		Filter(func(num int) bool { return num > 5 })
	base2 := stream.Map(
		base1.C,
		func(num int) string { return fmt.Sprintf("[%d]", num) },
	)

	for num := range base2.C {
		fmt.Println(num)
	}

	// Output:
	// [6]
	// [7]
	// [8]
	// [9]
}

func ExampleMapOrdered() {
	ordered := stream.MapOrdered(
		stream.Range2Channel(10),
		func(num int) string { return fmt.Sprintf("[%d]", num) },
	)

	fmt.Println(ordered.Max())

	// Output:
	// [9]
}

func ExampleMapComparable() {
	com := stream.MapComparable(
		stream.Slice2Channel(1, 1, 1, 2, 3, 3, 4),
		func(num int) string { return fmt.Sprintf("[%d]", num) },
	).Distinct()

	for i := range com.C {
		fmt.Println(i)
	}

	// Output:
	// [1]
	// [2]
	// [3]
	// [4]
}
