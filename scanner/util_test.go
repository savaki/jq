package scanner

import (
	"fmt"
	"testing"
)

func TestSkipSpace(t *testing.T) {
	content := []byte(" \t\n\r!")
	end, err := skipSpace(content, 0)
	if err != nil {
		t.FailNow()
	}
	if end+1 != len(content) {
		t.FailNow()
	}
}

func TestExpect(t *testing.T) {
	testCases := map[string]struct {
		In       string
		Expected string
		HasError bool
	}{
		"simple": {
			In:       "abc",
			Expected: "abc",
		},
		"extra": {
			In:       "abcdef",
			Expected: "abc",
		},
		"no match": {
			In:       "abc",
			Expected: "def",
			HasError: true,
		},
		"unexpected EOF": {
			In:       "ab",
			Expected: "abc",
			HasError: true,
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			pos, err := expect([]byte(tc.In), 0, []byte(tc.Expected)...)
			if tc.HasError {
				if err == nil {
					t.FailNow()
				}

			} else {
				if err != nil {
					fmt.Println(err)
					t.FailNow()
				}
				if pos != len([]byte(tc.Expected)) {
					t.FailNow()
				}

			}
		})
	}
}
