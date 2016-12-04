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
