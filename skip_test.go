package stream

import (
	"fmt"
)

func ExampleStream_Skip() {
	New([]int{1, 2, 3, 4}).
		Skip(2).
		ForEach(func(i int) { fmt.Println(i) })

	// Output:
	// 3
	// 4
}
