package scanner

import "bytes"

func FindKey(in []byte, pos int, k []byte) ([]byte, error) {
	pos, err := skipSpace(in, pos)
	if err != nil {
		return nil, err
	}

	if v := in[pos]; v != '{' {
		return nil, newError(pos, v)
	}
	pos++

	for {
		pos, err = skipSpace(in, pos)
		if err != nil {
			return nil, err
		}

		keyStart := pos
		// key
		pos, err = String(in, pos)
		if err != nil {
			return nil, err
		}
		key := in[keyStart+1 : pos-1]
		match := bytes.Equal(k, key)

		// colon
		pos, err = expectByte(in, pos, ':')
		if err != nil {
			return nil, err
		}

		pos, err = skipSpace(in, pos)
		if err != nil {
			return nil, err
		}

		valueStart := pos
		// data
		pos, err = Any(in, pos)
		if err != nil {
			return nil, err
		}

		if match {
			return in[valueStart:pos], nil
		}

		pos, err = skipSpace(in, pos)
		if err != nil {
			return nil, err
		}

		switch in[pos] {
		case ',':
			pos++
		case '}':
			return nil, errKeyNotFound
		}
	}
}
