package jq

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRegexp(t *testing.T) {
	Convey("Verify dot selectors", t, func() {
		testCases := map[string]struct {
			In   string
			From string
			To   string
		}{
			"simple": {
				In:   `[0]`,
				From: "0",
			},
			"range": {
				In:   `[0:1]`,
				From: "0",
				To:   "1",
			},
			"space before": {
				In:   ` [0:1]`,
				From: "0",
				To:   "1",
			},
			"space after": {
				In:   `[0:1] `,
				From: "0",
				To:   "1",
			},
			"space from": {
				In:   `[ 0 :1] `,
				From: "0",
				To:   "1",
			},
			"space to": {
				In:   `[0: 1 ] `,
				From: "0",
				To:   "1",
			},
		}

		for label, tc := range testCases {
			Convey(label, func() {
				matches := reArray.FindAllStringSubmatch(tc.In, -1)
				So(len(matches), ShouldEqual, 1)
				So(len(matches[0]), ShouldEqual, 4)
				So(matches[0][1], ShouldEqual, tc.From)
				So(matches[0][3], ShouldEqual, tc.To)
			})
		}
	})
}
