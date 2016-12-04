package scanner

import "errors"

func Any(in []byte, pos int) (int, error) {
	pos, err := skipSpace(in, pos)
	if err != nil {
		return 0, err
	}

	switch in[pos] {
	case '"':
		return String(in, pos)
	case '{':
		return Object(in, pos)
	case '.', '-', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
		return Number(in, pos)
	default:
		return 0, errors.New("invalid object")
	}
}
