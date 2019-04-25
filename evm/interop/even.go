package interop

import (
	"fmt"

	"github.com/evenfound/even-go/crypto"
)

// This file contains implementation of the Smart Contract API: namespace even.
// See the Even Network Smart Contract Specification.

func (e *Environment) evenPrintln(msg string) int {
	n, _ := fmt.Println(msg)
	return n
}

func (e *Environment) addString(str string) int {
	e.strings = append(e.strings, str)
	return len(e.strings) - 1
}

// evenHash hashes the given message.
func (e *Environment) evenHash(msg string) string {
	return string(crypto.Hash([]byte(msg)))
}

// evenSign signs a message and returns the signature.
func (e *Environment) evenSign(msg, privkey string) (string, error) {
	return crypto.Sign(msg, privkey)
}

// evenVerify checks if the signature has been generated from paired key.
func (e *Environment) evenVerify(msg, signature, pubkey string) (bool, error) {
	return crypto.Verify(msg, signature, pubkey)
}

// evenCreateWallet creates a wallet object.
func (e *Environment) evenCreateWallet(name, seed string) (handle, error) {
	w := wallet{}
	if err := w.create(name, seed); err != nil {
		return -1, err
	}
	h := e.addWallet(&w)
	return h, nil
}
