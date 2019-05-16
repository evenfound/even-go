package crypto

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	hashResult = []byte{140, 166, 110, 230, 178, 254, 75, 185, 40, 168, 227, 205, 47, 80, 141, 228, 17, 156, 8, 149, 242, 46, 1, 17, 23, 226, 44, 249, 177, 61, 231, 239}
)

func TestHash(t *testing.T) {
	Convey("Hash should work correctly", t, func() {
		h := Hash([]byte("Hello"))
		So(h, ShouldResemble, hashResult)
	})
}
