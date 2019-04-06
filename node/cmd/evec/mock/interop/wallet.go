package interop

// This file contains implementation of the Smart Contract API: Wallet object.
// See the Even Network Smart Contract Specification.

import (
	"fmt"
)

type wallet struct {
	//
}

// walletSave stores encrypted wallet locally.
func (e *Environment) walletSave(h handle, password string) error {
	w := e.wallet(h)
	return w.save(password)
}

func (w *wallet) create(name, seed string) error {
	fmt.Printf("Wallet created with name '%s'\n", name)
	return nil
}

func (w *wallet) save(password string) error {
	fmt.Printf("Wallet saved with password '%s'\n", password)
	return nil
}
