package scanner

// FindIndex accepts a JSON array and return the value of the element at the specified index
func FindIndex(in []byte, pos, index int) ([]byte, error) {
	pos, err := skipSpace(in, pos)
	if err != nil {
		return nil, err
	}

	if v := in[pos]; v != '[' {
		return nil, newError(pos, v)
	}
	pos++

	idx := 0
	for {
		pos, err = skipSpace(in, pos)
		if err != nil {
			return nil, err
		}

		itemStart := pos
		// data
		pos, err = Any(in, pos)
		if err != nil {
			return nil, err
		}
		if index == idx {
			return in[itemStart:pos], nil
		}

		pos, err = skipSpace(in, pos)
		if err != nil {
			return nil, err
		}

		switch in[pos] {
		case ',':
			pos++
		case ']':
			return nil, errIndexOutOfBounds
		}

		idx++
	}
}
