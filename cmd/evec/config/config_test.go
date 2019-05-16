package config

import (
	"flag"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/urfave/cli"
)

func TestOk(t *testing.T) {
	Convey("config.Ok() should return true if 0 languages", t, func() {
		BuildTengo = false
		BuildEvelyn = false
		BuildVyper = false
		ok, msg := Ok()
		So(ok, ShouldBeTrue)
		So(msg, ShouldBeBlank)
	})
	Convey("config.Ok() should return true if 1 language", t, func() {
		BuildTengo = true
		BuildEvelyn = false
		BuildVyper = false
		ok, msg := Ok()
		So(ok, ShouldBeTrue)
		So(msg, ShouldBeBlank)
	})
	Convey("config.Ok() should return false if >1 language", t, func() {
		BuildTengo = true
		BuildEvelyn = true
		BuildVyper = false
		ok, msg := Ok()
		So(ok, ShouldBeFalse)
		So(msg, ShouldEqual, "Command error: only one explicit language is allowed")
	})
	Convey("config.Ok() should return false if no arguments", t, func() {
		BuildTengo = false
		BuildEvelyn = false
		BuildVyper = false
		ok, msg := Ok()
		So(ok, ShouldBeFalse)
		So(msg, ShouldEqual, "no input files")
	})
}

func TestLooksLikeSourceFile(t *testing.T) {
	Convey("config.LooksLikeSourceFile() should return true if Tengo", t, func() {
		ok := LooksLikeSourceFile("filename.tgo")
		So(ok, ShouldBeTrue)
	})
	Convey("config.LooksLikeSourceFile() should return true if Evelyn", t, func() {
		ok := LooksLikeSourceFile("filename.evl")
		So(ok, ShouldBeTrue)
	})
	Convey("config.LooksLikeSourceFile() should return true if Vyper", t, func() {
		ok := LooksLikeSourceFile("filename.vy")
		So(ok, ShouldBeTrue)
	})
	Convey("config.LooksLikeSourceFile() should return false if other", t, func() {
		ok := LooksLikeSourceFile("filename.xxx")
		So(ok, ShouldBeFalse)
	})
}

func mockContext(args ...string) *cli.Context {
	app := cli.NewApp()
	flagset := flag.FlagSet{}
	if len(args) > 0 && args[0] != "" {
		if err := flagset.Parse(args); err != nil {
			panic(err)
		}
	}
	return cli.NewContext(app, &flagset, nil)
}
