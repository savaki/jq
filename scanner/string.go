package scanner

import "errors"

// String returns the position of the string that begins at the specified pos
func String(in []byte, pos int) (int, error) {
	pos, err := skipSpace(in, pos)
	if err != nil {
		return 0, err
	}

	max := len(in)

	if v := in[pos]; v != '"' {
		return 0, newError(pos, v)
	}
	pos++

	for {
		switch in[pos] {
		case '\\':
			if in[pos+1] == '"' {
				pos++
			}
		case '"':
			return pos + 1, nil
		}
		pos++

		if pos >= max {
			break
		}
	}

	return 0, errors.New("unclosed string")
}
