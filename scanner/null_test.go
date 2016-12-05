package scanner_test

import (
	"testing"

	"github.com/savaki/jq/scanner"
)

func BenchmarkNull(t *testing.B) {
	data := []byte("null")
	for i := 0; i < t.N; i++ {
		pos, err := scanner.Null(data, 0)
		if err != nil {
			t.FailNow()
		}
		if pos != 4 {
			t.FailNow()
		}
	}
}
