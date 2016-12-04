package jq

import (
	"strings"

	"github.com/savaki/jq/scanner"
)

// Dot extract the specific key from the map provided; to extract a nested value, use the Dot Op in conjunction with the
// Chain Op
func Dot(key string) OpFunc {
	key = strings.TrimSpace(key)
	if key == "" {
		return func(in []byte) ([]byte, error) { return in, nil }
	}

	k := []byte(key)

	return func(in []byte) ([]byte, error) {
		return scanner.FindKey(in, 0, k)
	}
}
