package api

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSmartContract(t *testing.T) {
	Convey("SmartContract should work correctly", t, func() {
		input := SmartContractInput{Uri: "/ipfs/zzz", EntryFunc: "default"}
		entry := input.GetEntryFunc()
		So(entry, ShouldEqual, "default")
	})
}
