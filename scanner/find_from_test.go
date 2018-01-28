// Copyright (c) 2016 Matt Ho <matt.ho@gmail.com>
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

package scanner_test

import (
	"testing"

	"github.com/savaki/jq/scanner"
)

func TestFindFrom(t *testing.T) {
	testCases := map[string]struct {
		In       string
		From     int
		Expected string
		HasErr   bool
	}{
		"all": {
			In:       `["a","b","c","d","e"]`,
			From:     0,
			Expected: `["a","b","c","d","e"]`,
		},
		"last": {
			In:       `["a","b","c","d","e"]`,
			From:     4,
			Expected: `["e"]`,
		},
		"middle": {
			In:       `["a","b","c","d","e"]`,
			From:     2,
			Expected: `["c","d","e"]`,
		},
		"mixed": {
			In:       `["a",{"hello":"world"},"c","d","e"]`,
			From:     0,
			Expected: `["a",{"hello":"world"},"c","d","e"]`,
		},
		"out of bounds": {
			In:     `["a",{"hello":"world"},"c","d","e"]`,
			From:   20,
			HasErr: true,
		},
	}

	for label, tc := range testCases {
		t.Run(label, func(t *testing.T) {
			data, err := scanner.FindFrom([]byte(tc.In), 0, tc.From)
			if tc.HasErr {
				if err == nil {
					t.FailNow()
				}
			} else {
				if string(data) != tc.Expected {
					t.FailNow()
				}
				if err != nil {
					t.FailNow()
				}
			}
		})
	}
}
