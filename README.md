# go-stream

[![GoCI](https://github.com/xuender/go-stream/workflows/Go/badge.svg)](https://github.com/xuender/go-stream/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/xuender/go-stream)](https://goreportcard.com/report/github.com/xuender/go-stream)
[![codecov](https://codecov.io/gh/xuender/go-stream/branch/master/graph/badge.svg?token=KCNTIM7DLH)](https://codecov.io/gh/xuender/go-stream)
[![GoDoc](https://godoc.org/github.com/xuender/go-stream?status.svg)](https://pkg.go.dev/github.com/xuender/go-stream)
[![GitHub license](https://img.shields.io/github/license/xuender/go-stream)](https://github.com/xuender/go-stream/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/xuender/go-stream)](https://github.com/xuender/go-stream/issues)
[![GitHub stars](https://img.shields.io/github/stars/xuender/go-stream)](https://github.com/xuender/gostream/stargazers)

Stream Collections for Go. Inspired in Java 8 Streams.

Use channel and Go1.18+ generic support.

✨ **`xuender/go-stream` is a Java 8 Streams style Go library based on Go 1.18+ Generics.**

## 🚀 Install

To install the library and command line program:

```shell
go get -u github.com/xuender/go-stream@latest
```

## 💡 Usage

You can import `stream` using:

```go
import "github.com/xuender/go-stream"
```

### NewBase

New BaseStream.

```go
package main

import (
  "fmt"

  "github.com/xuender/go-stream"
)

func main() {
  base := stream.NewBase(stream.Range2Channel(5)).
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

[[play](https://go.dev/play/p/LxULGyhWCi7)]

### Parallel

BaseStream to ParallelStream.

```go
package main

import (
  "fmt"
  "math/rand"
  "time"

  "github.com/xuender/go-stream"
)

func main() {
  stream.NewBase(stream.Range2Channel(1000)).
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

[[play](https://go.dev/play/p/wd5qrtmiwnS)]

### Map

Integer to string.

```go
package main

import (
  "fmt"

  "github.com/xuender/go-stream"
)

func main() {
  base := stream.Map(
    stream.Range2Channel(100),
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

[[play](https://go.dev/play/p/5QlT-D-Cv3V)]

### FlatMap

```go
package main

import (
  "fmt"

  "github.com/xuender/go-stream"
)

func main() {
  stream.FlatMap(
    stream.Slice2Channel([]int{0, 0}, []int{1, 2}, []int{2, 4}),
    func(num int) string { return fmt.Sprintf("[%d]", num) },
  ).ForEach(func(str string) {
    fmt.Println(str)
  })
}
```

Output:

```shell
[0]
[0]
[1]
[2]
[2]
[4]
```

[[play](https://go.dev/play/p/HXeeZOOkD2y)]

### Sorted

OrderedStream sorted.

```go
package main

import (
  "fmt"

  "github.com/xuender/go-stream"
)

func main() {
  stream.NewOrdered(stream.Slice2Channel(3, 2, 7, 1)).
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

[[play](https://go.dev/play/p/mSrDBfV_-d1)]

## 🛩 Functions

| Function | Type | State |
| - | - | - |
| AnyMatch | Terminal operations, short-circuiting | √ |
| AllMatch | Terminal operations, short-circuiting | √ |
| Count | Terminal operations | √ |
| Filter | Intermediate operations, Parallel | √ |
| FindFirst | Terminal operations, short-circuiting | √ |
| ForEach | Terminal operations, Parallel | √ |
| Limit | Intermediate operations | √ |
| NoneMatch | Terminal operations, short-circuiting | √ |
| Parallel | Intermediate operations | √ |
| Peek | Intermediate operations, Parallel | √ |
| Skip | Intermediate operations | √ |
| Sort | Intermediate operations | √ |
| Distinct | Intermediate operations, Comparable | √ |
| Max | Terminal operations, Ordered | √ |
| Min | Terminal operations, Ordered | √ |
| Reduce | Terminal operations, Ordered | √ |
| Sorted | Intermediate operations, Ordered | √ |
| Sequential | Intermediate operations, Parallel | √ |
| Map | Function | √ |
| FlatMap | Function | √ |

## 📝 License

© ender, 2023~time.Now

[MIT LICENSE](https://github.com/xuender/go-stream/blob/master/LICENSE)
