// Copyright 2016 Matt Ho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
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
