package hdwallet

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("New should return non-nil", t, func() {
		w := New("name", "password")
		So(w, ShouldNotBeNil)
	})
}
