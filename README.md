# go-stream

[![Go Report Card](https://goreportcard.com/badge/github.com/xuender/go-stream)](https://goreportcard.com/report/github.com/xuender/go-stream)

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
  input := make(chan int)
  base := stream.NewBase(input).
    Peek(func(num int) { fmt.Println("peek1:", num) }).
    Filter(func(num int) bool { return num > 1 }).
    Peek(func(num int) { fmt.Println("peek2:", num) })

  go func() {
    for i := 1; i < 5; i++ {
      input <- i
    }

    close(input)
  }()

  fmt.Println(base.Count())
}
```

Output:

```shell
peek1: 1
peek1: 2
peek1: 3
peek2: 2
peek2: 3
peek1: 4
peek2: 4
3
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
  input := make(chan int)
  parallel := stream.NewBase(input).
    Parallel(100).
    Filter(func(t int) bool { return t%7 == 0 })

  go func() {
    for i := 0; i < 1000; i++ {
      input <- i
    }

    close(input)
  }()

  parallel.ForEach(func(num int) {
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
  input := make(chan int)
  base := stream.Map(input, func(num int) string {
    return fmt.Sprintf("[%d]", num)
  }).Limit(3)

  go func(cha chan<- int) {
    for i := 0; i < 100; i++ {
      cha <- i
    }

    close(cha)
  }(input)

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
