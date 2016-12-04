package scanner

func Array(in []byte, pos int) (int, error) {
	pos, err := skipSpace(in, pos)
	if err != nil {
		return 0, err
	}

	if v := in[pos]; v != '[' {
		return 0, newError(pos, v)
	}
	pos++

	for {
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
		case ']':
			return pos + 1, nil
		}
	}
}
