package scanner

import "testing"

func TestSkipSpace(t *testing.T) {
	content := []byte(" \t\n\r!")
	end, err := skipSpace(content, 0)
	if err != nil {
		t.FailNow()
	}
	if end+1 != len(content) {
		t.FailNow()
	}
}
