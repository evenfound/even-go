package implementation

import (
	"testing"

	"github.com/evenfound/even-go/node/cmd/evec/config"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConstructionByFilenameSuffix(t *testing.T) {
	Convey("Tengo compiler", t, func() {
		c := New(".tgo")
		So(c, ShouldNotBeNil)
	})
	Convey("Evelyn compiler", t, func() {
		c := New(".evl")
		So(c, ShouldNotBeNil)
	})
	Convey("Vyper compiler is not yet implemented", t, func() {
		c := New(".vy")
		So(c, ShouldBeNil)
	})
	Convey("Solidity compiler is not yet implemented", t, func() {
		c := New(".sol")
		So(c, ShouldBeNil)
	})
	Convey("Unused suffix", t, func() {
		c := New("XXX")
		So(c, ShouldBeNil)
	})
}

func TestConstructionByGlobalFlag(t *testing.T) {
	config.BuildTengo = false
	config.BuildEvelyn = false
	config.BuildVyper = false
	config.BuildSolidity = false
	Convey("Tengo compiler", t, func() {
		config.BuildTengo = true
		c := New("XXX")
		So(c, ShouldNotBeNil)
		config.BuildTengo = false
	})
	Convey("Evelyn compiler", t, func() {
		config.BuildEvelyn = true
		c := New("XXX")
		So(c, ShouldNotBeNil)
		config.BuildEvelyn = false
	})
	Convey("Vyper compiler is not yet implemented", t, func() {
		config.BuildVyper = true
		c := New("XXX")
		So(c, ShouldBeNil)
		config.BuildVyper = false
	})
	Convey("Solidity compiler is not yet implemented", t, func() {
		config.BuildSolidity = true
		c := New("XXX")
		So(c, ShouldBeNil)
		config.BuildSolidity = false
	})
}
