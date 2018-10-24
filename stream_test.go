package stream

import (
	"fmt"
	"sort"
	"testing"
)

func ExampleNew() {
	arr := []int{1, 2, 3, 4, 5}
	i, err := New(arr).
		Peek(func(i int) { fmt.Println("peek1:", i) }).
		Filter(func(i int) bool { return i > 1 }).
		Peek(func(i int) { fmt.Println("peek2:", i) }).
		Map(func(i int) string { return fmt.Sprintf("id:%d", i) }).
		Peek(func(s string) { fmt.Println("peek3:", s) }).
		FindFirst()
	fmt.Println(i, err)

	// Output:
	// peek1: 1
	// peek1: 2
	// peek2: 2
	// peek3: id:2
	// id:2 <nil>
}
func ExampleStream() {
	arr := []int{1, 2, 3, 4, 5}
	s := New(arr).
		Filter(func(i int) bool { return i > 1 })
	count, err := s.Count()
	fmt.Println(count, err)
	first, err := s.FindFirst()
	fmt.Println(first, err)

	// Output:
	// 4 <nil>
	// 2 <nil>
}
func BenchmarkStream(b *testing.B) {
	b.StopTimer()
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = i
	}
	s := New(arr).
		Filter(func(i int) bool { return i > 33 }).
		Map(func(i int) int { return i * 2 }).
		Filter(func(i int) bool { return i > 100 })
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		l, _ := s.Sorted(func(x, y int) bool { return x > y }).FindFirst()
		if l != 198 {
			fmt.Println("error")
		}
	}
}

func BenchmarkStream2(b *testing.B) {
	b.StopTimer()
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = i
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		na := []int{}
		for _, a := range arr {
			if a > 33 {
				na = append(na, a*2)
			}
		}
		ret := []int{}
		for _, a := range na {
			if a > 100 {
				ret = append(ret, a)
			}
		}
		sort.Ints(ret)
		l := ret[len(ret)-1]
		if l != 198 {
			fmt.Println("error")
		}
	}
}
