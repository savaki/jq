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

package jq

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	reArray = regexp.MustCompile(`^\s*\[\s*(?:(\d+))?\s*(?:(:))?\s*(?:(\d+))?\s*\]\s*$`)
)

// Must is a convenience method similar to template.Must
func Must(op Op, err error) Op {
	if err != nil {
		panic(fmt.Errorf("unable to parse selector; %v", err.Error()))
	}

	return op
}

// Parse takes a string representation of a selector and returns the corresponding Op definition
func Parse(selector string) (Op, error) {
	segments := strings.Split(selector, ".")

	ops := make([]Op, 0, len(segments))
	for _, segment := range segments {
		key := strings.TrimSpace(segment)
		if key == "" {
			continue
		}

		if op, ok := parseArray(key); ok {
			ops = append(ops, op)
			continue
		}

		ops = append(ops, Dot(key))
	}

	return Chain(ops...), nil
}

func parseArray(key string) (Op, bool) {
	match := FindIndices(key)

	if len(match) == 0 {
		return nil, false
	}

	matches := match[0]

	if matches[1]+matches[2]+matches[3] == "" {
		return From(0), true
	}

	if matches[2] == "" {
		idx, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, false
		}
		return Index(idx), true
	}

	if matches[1] == "" {
		to, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, false
		}
		return To(to), true
	}

	if matches[3] == "" {
		from, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, false
		}
		return From(from), true
	}

	from, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, false
	}

	to, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil, false
	}

	return Range(from, to), true
}

func FindIndices(key string) [][]string {
	return reArray.FindAllStringSubmatch(key, -1)
}
