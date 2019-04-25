package transaction

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMarshalTimestamp(t *testing.T) {
	Convey("timestamp.MarshalJSON should work correctly", t, func() {
		t, err := time.Parse(timeLayout, "Mon Jan 15 10:12:24 +0300 2020")
		So(err, ShouldBeNil)
		So(t.String(), ShouldEqual, "2020-01-15 10:12:24 +0300 MSK")
		ts := timestamp(t)
		b, err := ts.MarshalJSON()
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, `"Wed Jan 15 10:12:24 +0300 2020"`)
	})
}

func TestUnmarshalTimestamp(t *testing.T) {
	Convey("timestamp.UnmarshalJSON should work correctly", t, func() {
		t := new(timestamp)
		So(time.Time(*t).String(), ShouldEqual, "0001-01-01 00:00:00 +0000 UTC")
		err := t.UnmarshalJSON([]byte(`"Wed Jan 15 10:12:24 +0300 2020"`))
		So(err, ShouldBeNil)
		So(time.Time(*t).String(), ShouldEqual, "2020-01-15 10:12:24 +0300 MSK")
	})
}
