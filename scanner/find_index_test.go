package scanner_test

import (
	"testing"

	"github.com/savaki/jq/scanner"
	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkFindIndex(t *testing.B) {
	data := []byte(`["hello","world"]`)

	for i := 0; i < t.N; i++ {
		data, err := scanner.FindIndex(data, 0, 1)
		if err != nil {
			t.FailNow()
			return
		}

		if string(data) != `"world"` {
			t.FailNow()
			return
		}
	}
}

func TestFindIndex(t *testing.T) {
	Convey("Verify FindIndex", t, func() {
		testCases := map[string]struct {
			In       string
			Index    int
			Expected string
			HasErr   bool
		}{
			"simple": {
				In:       `["hello","world"]`,
				Index:    1,
				Expected: `"world"`,
			},
			"spaced": {
				In:       ` [ "hello" , "world" ] `,
				Index:    1,
				Expected: `"world"`,
			},
			"all types": {
				In:       ` [ "hello" , 123, {"hello":"world"} ] `,
				Index:    2,
				Expected: `{"hello":"world"}`,
			},
		}

		for label, tc := range testCases {
			Convey(label, func() {
				data, err := scanner.FindIndex([]byte(tc.In), 0, tc.Index)
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
