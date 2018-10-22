# go-stream
Stream Collections for Go. Inspired in Java 8 Streams.

## Installation
To install the library and command line program, use the following:
```shell
go get -u github.com/xuender/go-stream
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	i, err := stream.New(arr).
		Peek(func(i int) { fmt.Println("peek1:", i) }).
		Filter(func(i int) bool { return i > 1 }).
		Peek(func(i int) { fmt.Println("peek2:", i) }).
		Map(func(i int) string { return fmt.Sprintf("id:%d", i) }).
		Peek(func(s string) { fmt.Println("peek3:", s) }).
		FindFirst()
	fmt.Println(i, err)
}
```
output
```shell
peek1: 1
peek1: 2
peek2: 2
peek3: id:2
id:2 <nil>
```

## Parallel
```go
package main

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	i, err := stream.New(arr).
		Parallel().
		Peek(func(i int) {
			fmt.Println("peek1:", i)
			time.Sleep(time.Second * time.Duration(i))
		}).
		Filter(func(i int) bool { return i > 1 }).
		Peek(func(i int) {
			fmt.Println("peek2:", i)
			time.Sleep(time.Second * time.Duration(i))
		}).
		FindFirst()
	fmt.Println(i, err)
}
```
output
```shell
peek1: 4
peek1: 6
peek1: 1
peek1: 5
peek1: 7
peek1: 8
peek1: 2
peek1: 3
peek2: 2
peek2: 3
peek2: 4
2 <nil>
```

## Functions
| Function | Type | State |
| - | - | - |
| Filter | Intermediate operations, Stateless | √ |
| Map | Intermediate operations, Stateless | √ |
| Peek | Intermediate operations, Stateless | √ |
| FindFirst | Terminal operations, short-circuiting | √ |
| ForEach | Terminal operations | √ |
| Max | Terminal operations | |
| Min | Terminal operations | |
| Count | Terminal operations | |
| AnyMatch | Terminal operations, short-circuiting | |
| AllMatch | Terminal operations, short-circuiting | |
| NoneMatch | Terminal operations, short-circuiting | |
| FlatMap | Intermediate operations, Stateless | |
| Distinct | Intermediate operations, Stateful | |
| Skip | Intermediate operations, Stateful | |
| Limit | Intermediate operations, Stateful | |
| Sorted | Intermediate operations, Stateful | |
