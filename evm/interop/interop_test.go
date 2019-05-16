package interop

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	goodScript = `even := import("even")
wallet := import("wallet")
result = ""
test := func() {
	w := even.createWallet("XXX", "novel interest powder ocean meadow act doctor toast element ability goddess april")
	result = wallet.save(w, "xxxxxx")
}
default := func() {
	w := even.createWallet("ZZZ", "quite deadening some fill others sing they sleep lie down there october")
	result = wallet.save(w, "zzzzzz")
}`

	badScript = "XXX()"

	runScript = `result = ""
default := func() {
	result = "ZZZ"
}
default()`
)

func TestConstructor(t *testing.T) {
	Convey("NewEnvironment(goodScript) should work correctly", t, func() {
		env, err := NewEnvironment([]byte(goodScript))
		So(env, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
	Convey("NewEnvironment(badScript) should return error", t, func() {
		env, err := NewEnvironment([]byte(badScript))
		So(env, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}

func TestEnvironment(t *testing.T) {
	Convey("Run(runScript) should work correctly", t, func() {
		env, err := NewEnvironment([]byte(runScript))
		So(env, ShouldNotBeNil)
		So(err, ShouldBeNil)
		err = env.Run()
		So(err, ShouldBeNil)
	})
	Convey("Get should work correctly", t, func() {
		env, err := NewEnvironment([]byte(runScript))
		So(env, ShouldNotBeNil)
		So(err, ShouldBeNil)
		err = env.Run()
		So(err, ShouldBeNil)
		res := env.Get("result").String()
		So(res, ShouldEqual, "ZZZ")
	})
}
