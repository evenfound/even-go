package vyper

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConstructor(t *testing.T) {
	Convey("Vyper compiler is not yet implemented", t, func() {
		c := New()
		So(c, ShouldBeNil)
	})
}
