# debuglog

Minimal Go logging library with zero runtime cost.

## Zero runtime cost? How?

Go build tags and compiler optimizations.

All functions in this package are wrapped in an `if enabled` condition, where
`enabled` is a const defined in build-tag-dependent files. If the build tag is
not set, the compiler will optimize out the entire block of code, leaving no
trace of it in the final binary:

```sh
$ go test -tags debuglog -bench .
goos: linux
goarch: amd64
pkg: github.com/paskozdilar/debuglog
cpu: 13th Gen Intel(R) Core(TM) i7-13800H
BenchmarkDebugLog-20            49902048                24.34 ns/op           32 B/op          1 allocs/op
BenchmarkNoDebugLog-20          1000000000               0.1356 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/paskozdilar/debuglog 1.394s

$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/paskozdilar/debuglog
cpu: 13th Gen Intel(R) Core(TM) i7-13800H
BenchmarkDebugLog-20            1000000000               0.1326 ns/op          0 B/op          0 allocs/op
BenchmarkNoDebugLog-20          1000000000               0.1315 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/paskozdilar/debuglog 0.302s
```

## Usage

Write logging code as usual:

```go
package main

import (
    "github.com/paskozdilar/debuglog"
)

func main() {
    debuglog.Print("Hello, world!")
    debuglog.Printf("Hello, %s!", "world")
    debuglog.Println("Hello,", "world!")
}
```

To enable it, add Go build tag `debuglog` to your commands:

```sh
go run -tags debuglog main.go
go build -tags debuglog main.go
go test -tags debuglog
```
