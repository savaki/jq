package scanner_test

import (
	"testing"

	"github.com/savaki/jq/scanner"
	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkObject(t *testing.B) {
	data := []byte(`{"hello":"world"}`)

	for i := 0; i < t.N; i++ {
		end, err := scanner.Object(data, 0)
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

func TestObject(t *testing.T) {
	Convey("Verify Object", t, func() {
		testCases := map[string]struct {
			In     string
			Out    string
			HasErr bool
		}{
			"simple": {
				In:  `{"hello":"world"}`,
				Out: `{"hello":"world"}`,
			},
			"spaced": {
				In:  ` { "hello" : "world" } `,
				Out: ` { "hello" : "world" }`,
			},
		}

		for label, tc := range testCases {
			Convey(label, func() {
				end, err := scanner.Object([]byte(tc.In), 0)
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
