package stream

import (
	"fmt"
)

func ExampleStream_Limit() {
	New([]int{1, 2, 3}).
		Limit(2).
		ForEach(func(i int) { fmt.Println(i) })

	// Output:
	// 1
	// 2
}
