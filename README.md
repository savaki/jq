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


