package interop

import "fmt"

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

// evenHashMessage hashes the given message.
func (e *Environment) evenHashMessage(msg string) string {
	return msg + "XXX"
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
