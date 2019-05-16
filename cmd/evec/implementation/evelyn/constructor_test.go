package evelyn

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConstructor(t *testing.T) {
	Convey("Evelyn compiler", t, func() {
		c := New()
		So(c, ShouldNotBeNil)
	})
}
