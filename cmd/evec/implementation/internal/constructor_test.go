package internal

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConstructor(t *testing.T) {
	Convey("Tengo compiler", t, func() {
		c := New()
		So(c, ShouldNotBeNil)
	})
}
