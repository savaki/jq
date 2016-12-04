package jq_test

import (
	"testing"

	"github.com/savaki/jq"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {
	Convey("Verify dot selectors", t, func() {
		testCases := map[string]struct {
			In       string
			Op       string
			Expected string
			HasError bool
		}{
			"simple": {
				In:       `{"hello":"world"}`,
				Op:       ".hello",
				Expected: `"world"`,
			},
			"nested": {
				In:       `{"a":{"b":"world"}}`,
				Op:       ".a.b",
				Expected: `"world"`,
			},
			"index": {
				In:       `["a","b","c"]`,
				Op:       ".[1]",
				Expected: `"b"`,
			},
			"range": {
				In:       `["a","b","c"]`,
				Op:       ".[1:2]",
				Expected: `["b","c"]`,
			},
			"nested index": {
				In:       `{"abc":"-","def":["a","b","c"]}`,
				Op:       ".def.[1]",
				Expected: `"b"`,
			},
			"nested range": {
				In:       `{"abc":"-","def":["a","b","c"]}`,
				Op:       ".def.[1:2]",
				Expected: `["b","c"]`,
			},
		}

		for label, tc := range testCases {
			Convey(label, func() {
				op, err := jq.Parse(tc.Op)
				So(err, ShouldBeNil)

				data, err := op.Apply([]byte(tc.In))
				if tc.HasError {
					So(err, ShouldNotBeNil)
				} else {
					So(string(data), ShouldEqual, tc.Expected)
					So(err, ShouldBeNil)
				}
			})
		}
	})
}
