# jq

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
| .foo.bar |  nested value at key |


