// Package jp offers a highly performant json selector in the style of the jq command line
//
// Usage of this package involves the concept of an Op.  An Op is a transformation that converts a []byte into a []byte.
// To get started, use the Parse function to obtain an Op.
//
//     op, err := jq.Parse(".key")
//
// This will create an Op that will accept a JSON object in []byte format and return the value associated with "key."
// For example:
//
//     data, _ := op.Apply([]byte(`{"key":"value"}`))
//     fmt.Println(string(data))
//
// Will print the string ```"value"```
//
package jq
