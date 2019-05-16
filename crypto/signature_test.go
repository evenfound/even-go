package crypto

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	message   = "Hello"
	privkey   = "a51193ec0ee7f4734ec054aa86328fc575a8eb418b5712d42b872eb14a22ca2e"
	pubkey    = "024d08969cb5dd02b4374412fa6cebfcbcf8e796d5e82f875fd61ede5978b23f75"
	signature = "30440220790fc006bb8321354456b89e2c6a855e4f77047704b14d799faa73363ee3e182022060b6273f6857154cc26cf7c1b84f79db98b3c014d9869dbbbe2e5c2b57ba06cb"
)

func TestSign(t *testing.T) {
	Convey("Sign should work correctly", t, func() {
		So(len(privkey), ShouldEqual, 64)
		So(len(pubkey), ShouldEqual, 66)
		So(len(signature), ShouldEqual, 140)

		s, err := Sign(message, privkey)
		So(err, ShouldBeNil)
		So(s, ShouldEqual, signature)

		s, err = Sign(message, privkey+"0000")
		So(err, ShouldBeNil)
		So(s, ShouldEqual, "3045022100a876f387adcae5a24d0049591de4d61dffc8e4834980e9febfd1166af8b5161a02202b132a30cee1b75f3de36cdf88acf0281e0f7c1fdb247e98e0471da3653007d0")

		s, err = Sign(message, privkey[:20])
		So(err, ShouldBeNil)
		So(s, ShouldEqual, "3044022020583eaaa0c7198a69cca52e2d9bec195009be6cb4ba20d9a0a8e9e3eb2a719102201908b57876ff3381167ffad34144116ee81cc7c19e46ab5a1cc220a574996030")
	})
}

func TestSignBad(t *testing.T) {
	Convey("Sign should handle errors correctly", t, func() {
		s, err := Sign(message, privkey+"GOGO")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "encoding/hex: invalid byte: U+0047 'G'")
		So(s, ShouldBeEmpty)
	})
}

func TestVerify(t *testing.T) {
	Convey("Verify should work correctly", t, func() {
		So(len(privkey), ShouldEqual, 64)
		So(len(pubkey), ShouldEqual, 66)
		So(len(signature), ShouldEqual, 140)

		ok, err := Verify(message, signature, pubkey)
		So(err, ShouldBeNil)
		So(ok, ShouldBeTrue)

		ok, err = Verify(message+"2", signature, pubkey)
		So(err, ShouldBeNil)
		So(ok, ShouldBeFalse)

		ok, err = Verify(message, signature+"00", pubkey)
		So(err, ShouldBeNil)
		So(ok, ShouldBeTrue)
	})
}

func TestVerifyBad(t *testing.T) {
	Convey("Verify should handle errors correctly", t, func() {
		ok, err := Verify(message, "00"+signature, pubkey)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "malformed signature: no header magic")
		So(ok, ShouldBeFalse)

		ok, err = Verify(message, signature, "00"+pubkey)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "invalid pub key length 34")
		So(ok, ShouldBeFalse)

		ok, err = Verify(message, signature, pubkey+"FFFF")
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "invalid pub key length 35")
		So(ok, ShouldBeFalse)
	})
}
