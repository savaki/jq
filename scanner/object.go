package scanner

// Object returns the position of the end of the object that begins at the specified pos
func Object(in []byte, pos int) (int, error) {
	pos, err := skipSpace(in, pos)
	if err != nil {
		return 0, err
	}

	if v := in[pos]; v != '{' {
		return 0, newError(pos, v)
	}
	pos++

	for {
		// key
		pos, err = String(in, pos)
		if err != nil {
			return 0, err
		}

		// colon
		pos, err = expectByte(in, pos, ':')
		if err != nil {
			return 0, err
		}

		// data
		pos, err = Any(in, pos)
		if err != nil {
			return 0, err
		}

		pos, err = skipSpace(in, pos)
		if err != nil {
			return 0, err
		}

		switch in[pos] {
		case ',':
			pos++
		case '}':
			return pos + 1, nil
		}
	}
}
