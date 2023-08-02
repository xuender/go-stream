# go-stream

[![GoCI](https://github.com/xuender/go-stream/workflows/Go/badge.svg)](https://github.com/xuender/go-stream/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuender/go-stream)](https://goreportcard.com/report/github.com/xuender/go-stream)
[![codecov](https://codecov.io/gh/xuender/go-stream/branch/master/graph/badge.svg?token=KCNTIM7DLH)](https://codecov.io/gh/xuender/go-stream)
[![GoDoc](https://godoc.org/github.com/xuender/go-stream?status.svg)](https://pkg.go.dev/github.com/xuender/go-stream)
[![GitHub license](https://img.shields.io/github/license/xuender/go-stream)](https://github.com/xuender/go-stream/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/xuender/go-stream)](https://github.com/xuender/go-stream/issues)
[![GitHub stars](https://img.shields.io/github/stars/xuender/go-stream)](https://github.com/xuender/gostream/stargazers)

Stream Collections for Go. Inspired in Java 8 Streams.

Based on channel and go1.18 generic support.

## Install

To install the library and command line program:

```shell
go get -u github.com/xuender/go-stream
```

## Base

BaseStream

```go
package main

import (
  "fmt"

  "github.com/xuender/go-stream"
)

func main() {
  base := stream.NewBase(stream.Range2Channel(1, 5)).
    Peek(func(num int) { fmt.Println("peek1:", num) }).
    Filter(func(num int) bool { return num > 2 }).
    Peek(func(num int) { fmt.Println("peek2:", num) })

  fmt.Println(base.Count())
}
```

Output:

```shell
peek1: 0
peek1: 1
peek1: 2
peek1: 3
peek1: 4
peek2: 3
peek2: 4
2
```

## Parallel

ParallelStream

```go
package main

import (
  "fmt"
  "math/rand"
  "time"

  "github.com/xuender/go-stream"
)

func main() {
  stream.NewBase(stream.Range2Channel(1, 1000)).
    Parallel(100).
    Filter(func(num int) bool { return num%7 == 0 }).
    ForEach(func(num int) {
      dur := time.Duration(rand.Intn(1000))

      time.Sleep(time.Millisecond * dur)
      fmt.Printf("%d\t%dms\n", num, dur)
    })
}
```

Output:

```shell
322     0ms  
651     2ms  
483     2ms  
182     15ms 
266     26ms 
567     33ms 
742     10ms 
175     47ms 
476     59ms 
7       59ms 
...
```

## Map

int to string.

```go
package main

import (
  "fmt"

  "github.com/xuender/go-stream"
)

func main() {
  base := stream.Map(
    stream.Range2Channel(1, 100),
    func(num int) string { return fmt.Sprintf("[%d]", num) },
  ).Limit(3)

  for i := range base.C {
    fmt.Println(i)
  }
}
```

Output:

```shell
[0]
[1]
[2]
```

## Sorted

```golang
package main

import (
  "fmt"

  "github.com/xuender/go-stream"
)

func main() {
  stream.NewOrdered(stream.Slice2Channel(1, 3, 2, 7, 1)).
    Sorted().
    ForEach(func(num int) {
      fmt.Println(num)
    })
}
```

Output:

```shell
1
2
3
7
```

## Functions

| Function | Type | State |
| - | - | - |
| ForEach | Terminal operations | √ |
| Count | Terminal operations | √ |
| Max | Terminal operations | √ |
| Min | Terminal operations | √ |
| Parallel | Intermediate operations, Stateful | √ |
| Limit | Intermediate operations, Stateful | √ |
| Skip | Intermediate operations, Stateful | √ |
| Filter | Intermediate operations, Stateless | √ |
| Peek | Intermediate operations, Stateless | √ |
| FindFirst | Terminal operations, short-circuiting | √ |
| AnyMatch | Terminal operations, short-circuiting | √ |
| AllMatch | Terminal operations, short-circuiting | √ |
| NoneMatch | Terminal operations, short-circuiting | √ |
| Distinct | Intermediate operations, Stateful | √ |
| Sorted | Intermediate operations, Stateful | √ |
| Sort | Intermediate operations, Stateful | √ |
| Reduce | Terminal operations | √ |
| Sequential | Intermediate operations, Stateful | √ |
| Map | Intermediate operations, Function | √ |
| FlatMap | Intermediate operations, Function | √ |

## License

© ender, 2023~time.Now

[MIT LICENSE](https://github.com/xuender/go-stream/blob/master/LICENSE)
