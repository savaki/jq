# jq

[![GoDoc](https://godoc.org/github.com/savaki/jq?status.svg)](https://godoc.org/github.com/savaki/jq)
[![Build Status](https://snap-ci.com/savaki/jq/branch/master/build_image)](https://snap-ci.com/savaki/jq/branch/master)

Golang implementation of incredibly useful jq command line tool.

## Installation

```
go get github.com/savaki/jq
```

## Example

```go
package main

import (
	"fmt"

	"github.com/savaki/jq"
)

func main() {
	data := []byte(`{"hello":"world"}`)
	op, _ := jq.Parse(".hello")
	value, _ := op.Apply(data) // value == '"world"'
	fmt.Println(string(value))
}
```

## Syntax

| syntax | meaning|
| :--- | :--- |
| . |  unchanged input |
| .foo |  value at key |
| .foo.bar |  value at nested key |
| .[0] | value at specified element of array | 
| .[0:1] | array of specified elements of array, inclusive |
| .foo.[0] | nest value |

## Performance

```
BenchmarkChain-8   	 5000000	       285 ns/op	       0 B/op	       0 allocs/op
BenchmarkDot-8     	20000000	       115 ns/op	       0 B/op	       0 allocs/op

BenchmarkAny-8         	20000000	        80.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkArray-8       	20000000	       108 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindIndex-8   	10000000	       125 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindKey-8     	10000000	       125 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindRange-8   	10000000	       186 ns/op	      16 B/op	       1 allocs/op
BenchmarkNumber-8      	50000000	        28.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkObject-8      	20000000	        98.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkString-8      	30000000	        40.4 ns/op	       0 B/op	       0 allocs/op
```

