package jq

import "strings"

// Parse takes a string representation of a selector and returns the corresponding Op definition
func Parse(selector string) (Op, error) {
	segments := strings.Split(selector, ".")

	ops := make([]Op, 0, len(segments))
	for _, segment := range segments {
		key := strings.TrimSpace(segment)
		if key == "" {
			continue
		}

		ops = append(ops, Dot(key))
	}

	return Chain(ops...), nil
}
