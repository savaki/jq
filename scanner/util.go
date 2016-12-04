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
package scanner

import (
	"errors"
	"fmt"
	"unicode"
	"unicode/utf8"
)

var (
	errUnexpectedEOF    = errors.New("unexpected EOF")
	errKeyNotFound      = errors.New("key not found")
	errIndexOutOfBounds = errors.New("index out of bounds")
	errToLessThanFrom   = errors.New("to index less than from index")
)

func skipSpace(in []byte, pos int) (int, error) {
	for {
		r, size := utf8.DecodeRune(in[pos:])
		if size == 0 {
			return 0, errUnexpectedEOF
		}
		if !unicode.IsSpace(r) {
			break
		}
		pos += size
	}

	return pos, nil
}

func expectByte(in []byte, pos int, expected byte) (int, error) {
	pos, err := skipSpace(in, pos)
	if err != nil {
		return 0, err
	}

	if v := in[pos]; v != expected {
		return 0, newError(pos, v)
	}

	return pos + 1, nil
}

func newError(pos int, b byte) error {
	return fmt.Errorf("invalid character at position, %v; %v", pos, string([]byte{b}))
}
