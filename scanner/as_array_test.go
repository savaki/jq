package scanner_test

import (
	"bytes"
	"testing"

	"github.com/savaki/jq/scanner"
)

func BenchmarkAsArray(t *testing.B) {
	data := []byte(`["hello","world"]`)

	for i := 0; i < t.N; i++ {
		out, err := scanner.AsArray(data, 0)
		if err != nil {
			t.Errorf("expected nil err; got %v", err)
			return
		}
		if v := len(out); v != 2 {
			t.Errorf("want %v, got %v", 2, v)
			return
		}
	}
}

func TestAsArray(t *testing.T) {
	testCases := map[string]struct {
		In     string
		Out    []string
		HasErr bool
	}{
		"simple": {
			In:  `["hello","world"]`,
			Out: []string{`"hello"`, `"world"`},
		},
		"spaced": {
			In:  ` [ "hello" , "world" ] `,
			Out: []string{`"hello"`, `"world"`},
		},
		"all types": {
			In:  ` [ "hello" , 123, {"hello":"world"} ] `,
			Out: []string{`"hello"`, `123`, `{"hello":"world"}`},
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			out, err := scanner.AsArray([]byte(tc.In), 0)
			if tc.HasErr {
				if err == nil {
					t.FailNow()
				}

			} else {
				if err != nil {
					t.Errorf("expected nil err; got %v", err)
					return
				}
				if len(out) != len(tc.Out) {
					t.Errorf("expected output lengths to match; want %v, got %v", len(tc.Out), len(out))
					return
				}
				for index, item := range tc.Out {
					if v := out[index]; bytes.Compare(v, []byte(item)) != 0 {
						t.Errorf("expected content at index %v to match; want %v, got %v", index, item, string(v))
						return
					}
				}
			}
		})
	}
}
