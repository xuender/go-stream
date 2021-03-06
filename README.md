# go-stream
Stream Collections for Go. Inspired in Java 8 Streams.

Because using reflect is slow, it is **not recommended for production environment**.

Expect the support of GO2 version generic.

## Installation
To install the library and command line program, use the following:
```shell
go get -u github.com/xuender/go-stream
```

## Usage example
Sequential stream:
```go
package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	f, err := stream.New(arr).
		Peek(func(i int) { fmt.Println("peek1:", i) }).
		Filter(func(i int) bool { return i > 1 }).
		Peek(func(i int) { fmt.Println("peek2:", i) }).
		Map(func(i int) string { return fmt.Sprintf("id:%d", i) }).
		Peek(func(s string) { fmt.Println("peek3:", s) }).
		FindFirst()

	fmt.Println(f, err)
}
```
Output:
```
peek1: 1
peek1: 2
peek2: 2
peek3: id:2
id:2 <nil>
```

## Parallel example
Parallel stream:
```go
package main

import (
	"fmt"
	"time"

	"github.com/xuender/go-stream"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	f, err := stream.New(arr).
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

	fmt.Println(f, err)
}
```
Output:
```
peek1: 4
peek1: 6
peek1: 1
peek1: 5
peek1: 2
peek1: 3
peek2: 2
peek2: 3
peek2: 4
2 <nil>
```

## Complex example
```go
package main

import (
	"fmt"

	"github.com/xuender/go-stream"
)

func main() {
  arr := []string{"go", "stream", "is", "good"}

	sum, err := stream.New(arr).
		FlatMap(func(s string) []byte { return []byte(s) }).
		Map(func(s byte) int { return int(s) }).
		Reduce(func(x, y int) int { return x + y })

	fmt.Println(sum, err)
}
```
Output:
```
1511 <nil>
```

## Functions
| Function | Type | State |
| - | - | - |
| Sequential | Intermediate operations, Stateful | √ |
| Parallel | Intermediate operations, Stateful | √ |
| Skip | Intermediate operations, Stateful | √ |
| Limit | Intermediate operations, Stateful | √ |
| Distinct | Intermediate operations, Stateful | √ |
| Sorted | Intermediate operations, Stateful | √ |
| Filter | Intermediate operations, Stateless | √ |
| Map | Intermediate operations, Stateless | √ |
| FlatMap | Intermediate operations, Stateless | √ |
| Peek | Intermediate operations, Stateless | √ |
| FindFirst | Terminal operations, short-circuiting | √ |
| AnyMatch | Terminal operations, short-circuiting | √ |
| AllMatch | Terminal operations, short-circuiting | √ |
| NoneMatch | Terminal operations, short-circuiting | √ |
| ForEach | Terminal operations | √ |
| Count | Terminal operations | √ |
| Max | Terminal operations | √ |
| Min | Terminal operations | √ |
| Reduce | Terminal operations | √ |
