// Copyright (C) 2017-2019 The Even Network Developers

package app

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInit(t *testing.T) {
	Convey("app.Init() should work correctly", t, func() {
		So(Init, ShouldNotPanic)
	})
}
