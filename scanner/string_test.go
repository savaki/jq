package scanner_test

import (
	"testing"

	"unicode/utf8"

	"github.com/savaki/jq/scanner"
	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	Convey("Verify String", t, func() {
		testCases := map[string]struct {
			In     string
			Out    string
			HasErr bool
		}{
			"simple": {
				In:  `"hello"`,
				Out: `"hello"`,
			},
			"array": {
				In:  `"hello", "world"`,
				Out: `"hello"`,
			},
			"escaped": {
				In:  `"hello\"\"world"`,
				Out: `"hello\"\"world"`,
			},
			"unclosed": {
				In:     `"hello`,
				HasErr: true,
			},
			"unclosed escape": {
				In:     `"hello\"`,
				HasErr: true,
			},
			"utf8": {
				In:  `"生日快乐"`,
				Out: `"生日快乐"`,
			},
		}

		for label, tc := range testCases {
			Convey(label, func() {
				end, err := scanner.String([]byte(tc.In), 0)
				if tc.HasErr {
					So(err, ShouldNotBeNil)
				} else {
					data := tc.In[0:end]
					So(string(data), ShouldEqual, tc.Out)
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

func TestDecode(t *testing.T) {
	Convey("Verify can decode", t, func() {
		v := ""
		_, size := utf8.DecodeRune([]byte(v))
		So(size, ShouldEqual, 0)
	})
}
