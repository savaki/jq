package scanner_test

import (
	"testing"

	"github.com/savaki/jq/scanner"
	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkAny(t *testing.B) {
	data := []byte(`"Hello, 世界 - 生日快乐"`)

	for i := 0; i < t.N; i++ {
		end, err := scanner.Any(data, 0)
		if err != nil {
			t.FailNow()
			return
		}

		if end == 0 {
			t.FailNow()
			return
		}
	}
}

func TestAny(t *testing.T) {
	text := `  "hello"`
	Convey("Verify trims space", t, func() {
		end, err := scanner.Any([]byte(text), 0)
		So(err, ShouldBeNil)

		any := text[0:end]
		So(string(any), ShouldEqual, text)
	})
}
