package scanner

var (
	n = []byte("null")
)

// Null verifies the contents of bytes provided is a null starting as pos
func Null(in []byte, pos int) (int, error) {
	switch in[pos] {
	case 'n':
		return expect(in, pos, n...)
		return pos + 4, nil
	default:
		return 0, errUnexpectedValue
	}
}
