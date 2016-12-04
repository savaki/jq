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
