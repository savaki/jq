package jq_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/savaki/jq"
)

func BenchmarkDot(t *testing.B) {
	op := jq.Dot("hello")
	data := []byte(`{"hello":"world"}`)

	for i := 0; i < t.N; i++ {
		_, err := op.Apply(data)
		if err != nil {
			t.FailNow()
			return
		}
	}
}

func TestDot(t *testing.T) {
	Convey("Verify dot selectors", t, func() {
		testCases := map[string]struct {
			In       string
			Key      string
			Expected string
			HasError bool
		}{
			"simple": {
				In:       `{"hello":"world"}`,
				Key:      "hello",
				Expected: `"world"`,
			},
			"key not found": {
				In:       `{"hello":"world"}`,
				Key:      "junk",
				HasError: true,
			},
			"unclosed value": {
				In:       `{"hello":"world`,
				Key:      "hello",
				HasError: true,
			},
		}

		for label, tc := range testCases {
			Convey(label, func() {
				op := jq.Dot(tc.Key)
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
