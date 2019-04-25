package rpc

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPreconditions(t *testing.T) {
	Convey("Precondition checks should work correctly", t, func() {
		So(isCorrectFilename("XXX"), ShouldBeFalse)
		So(isCorrectFilename("file:///tmp/xxx"), ShouldBeTrue)
		So(isCorrectFilename("/ipfs/xxx"), ShouldBeTrue)
		So(isCorrectFunction("123"), ShouldBeFalse)
		So(isCorrectFunction("5_func"), ShouldBeFalse)
		So(isCorrectFunction("func"), ShouldBeTrue)
		So(isCorrectFunction("func123"), ShouldBeTrue)
		So(isCorrectFunction("func_789"), ShouldBeTrue)
	})
}
