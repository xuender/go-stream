package stream

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleStream_FlatMap() {
	arr := []string{"abc", "ab", "a"}

	err := New(arr).
		FlatMap(func(s string) []string { return strings.Split(s, "") }).
		ForEach(func(s string) { fmt.Println(s) })

	fmt.Println(err)

	// Output:
	// a
	// b
	// c
	// a
	// b
	// a
	// <nil>
}

func TestStream_FlatMap(t *testing.T) {
	arr := []string{"abc", "ab", "a"}

	c, err := New(arr).
		FlatMap(func(s string) []string { return strings.Split(s, "") }).
		Count()
	if err != nil {
		t.Errorf("Error is not nil %s", err)
	}
	if c != 6 {
		t.Errorf("count is %d not 6", c)
	}
}
