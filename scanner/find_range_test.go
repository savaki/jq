package scanner_test

import (
	"testing"

	"github.com/savaki/jq/scanner"
	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkFindRange(t *testing.B) {
	data := []byte(`["a","b","c","d","e"]`)

	for i := 0; i < t.N; i++ {
		out, err := scanner.FindRange(data, 0, 1, 2)
		if err != nil {
			t.FailNow()
			return
		}

		if string(out) != `["b","c"]` {
			t.FailNow()
			return
		}
	}
}

func TestFindRange(t *testing.T) {
	Convey("Verify FindKey", t, func() {
		testCases := map[string]struct {
			In       string
			From     int
			To       int
			Expected string
			HasErr   bool
		}{
			"simple": {
				In:       `["a","b","c","d","e"]`,
				From:     1,
				To:       2,
				Expected: `["b","c"]`,
			},
			"single": {
				In:       `["a","b","c","d","e"]`,
				From:     1,
				To:       1,
				Expected: `["b"]`,
			},
			"mixed": {
				In:       `["a",{"hello":"world"},"c","d","e"]`,
				From:     1,
				To:       1,
				Expected: `[{"hello":"world"}]`,
			},
			"ordering": {
				In:     `["a",{"hello":"world"},"c","d","e"]`,
				From:   1,
				To:     0,
				HasErr: true,
			},
			"out of bounds": {
				In:     `["a",{"hello":"world"},"c","d","e"]`,
				From:   1,
				To:     20,
				HasErr: true,
			},
		}

		for label, tc := range testCases {
			Convey(label, func() {
				data, err := scanner.FindRange([]byte(tc.In), 0, tc.From, tc.To)
				if tc.HasErr {
					So(err, ShouldNotBeNil)
				} else {
					So(string(data), ShouldEqual, tc.Expected)
					So(err, ShouldBeNil)
				}
			})
		}
	})
}
