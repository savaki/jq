package scanner_test

import (
	"testing"

	"github.com/savaki/jq/scanner"
	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkAny(t *testing.B) {
	data := []byte(`"Hello, 世界 - 生日快乐"`)

	for i := 0; i < t.N; i++ {
		end, err := scanner.Any(data, 0)
		if err != nil {
			t.FailNow()
			return
		}

		if end == 0 {
			t.FailNow()
			return
		}
	}
}

func TestAny(t *testing.T) {
	Convey("Verify trims space", t, func() {
		testCases := map[string]struct {
			In  string
			Out string
		}{
			"string": {
				In:  `"hello"`,
				Out: `"hello"`,
			},
			"array": {
				In:  `["a","b","c"]`,
				Out: `["a","b","c"]`,
			},
			"object": {
				In:  `{"a":"b"}`,
				Out: `{"a":"b"}`,
			},
			"number": {
				In:  `1.234e+10`,
				Out: `1.234e+10`,
			},
		}

		for label, tc := range testCases {
			Convey(label, func() {
				end, err := scanner.Any([]byte(tc.In), 0)
				So(err, ShouldBeNil)
				data := tc.In[0:end]
				So(string(data), ShouldEqual, tc.Out)
			})
		}
	})
}
