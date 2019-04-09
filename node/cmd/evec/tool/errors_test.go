package tool

import (
	"testing"

	"github.com/evenfound/even-go/node/cmd/evec/config"

	"github.com/ztrue/tracerr"

	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMust(t *testing.T) {
	Convey("tool.Must(nil) should quietly return", t, func() {
		m := func() { Must(nil) }
		So(m, ShouldNotPanic)
	})
	Convey("tool.Must(non-nil) should panic", t, func() {
		m := func() { Must(NewError("Test")) }
		So(m, ShouldPanic)
	})
}

func TestIgnore(t *testing.T) {
	Convey("tool.Ignore(nil) should quietly return", t, func() {
		i := func() { Ignore(nil) }
		So(i, ShouldNotPanic)
	})
	Convey("tool.Ignore(non-nil) should quietly return", t, func() {
		i := func() { Ignore(nil) }
		So(i, ShouldNotPanic)
	})
}

func TestNewError(t *testing.T) {
	Convey("tool.NewError(NonDebug) should return correct error", t, func() {
		config.Debug = false
		text := "NonDebug"
		got := NewError(text)
		So(got.Error(), ShouldEqual, text)
		So(got, ShouldHaveSameTypeAs, errors.New(""))
	})
	Convey("tool.NewError(Debug) should return correct error", t, func() {
		config.Debug = true
		text := "Debug"
		got := NewError(text)
		So(got.Error(), ShouldEqual, text)
		So(got, ShouldHaveSameTypeAs, tracerr.New(""))
	})
}

func TestWrap(t *testing.T) {
	err := errors.New("Head")
	Convey("tool.Wrap(NonDebug) should return correct error", t, func() {
		config.Debug = false
		text := "NonDebug"
		got := Wrap(err, text)
		expected := "NonDebug: Head"
		So(got.Error(), ShouldEqual, expected)
		So(got, ShouldHaveSameTypeAs, errors.WithMessage(err, ""))
	})
	Convey("tool.Wrap(Debug) should return correct error", t, func() {
		config.Debug = true
		text := "Debug"
		got := Wrap(err, text)
		expected := "Debug: Head"
		So(got.Error(), ShouldEqual, expected)
		So(got, ShouldHaveSameTypeAs, tracerr.New(""))
	})
}

func TestWrapf(t *testing.T) {
	err := errors.New("Head")
	Convey("tool.Wrapf(NonDebug) should return correct error", t, func() {
		config.Debug = false
		got := Wrapf(err, "config.Debug = %t", config.Debug)
		expected := "config.Debug = false: Head"
		So(got.Error(), ShouldEqual, expected)
		So(got, ShouldHaveSameTypeAs, errors.WithMessage(err, ""))
	})
	Convey("tool.Wrap(Debug) should return correct error", t, func() {
		config.Debug = true
		got := Wrapf(err, "config.Debug = %t", config.Debug)
		expected := "config.Debug = true: Head"
		So(got.Error(), ShouldEqual, expected)
		So(got, ShouldHaveSameTypeAs, tracerr.New(""))
	})
}
