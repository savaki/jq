# jq

Golang implementation of incredibly useful jq command line tool.

## Example


```go
data := `{"hello":"world"}`
op := jq.Parse(".hello")
value, _ := op.Apply(data) // value == '"world"'
```


