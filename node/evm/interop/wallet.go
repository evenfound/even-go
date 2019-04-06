package interop

// This file contains implementation of the Smart Contract API: Wallet object.
// See the Even Network Smart Contract Specification.

import (
	"errors"
	"fmt"

	evenWallet "github.com/evenfound/even-go/node/hdwallet"
)

const (
	invalidPhraseError = "The provided seed phrase is not valid"
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
	var ewl evenWallet.HDWallet

	// Wallet name will be used to create directory,
	// where will be stored all wallet data
	ewl.WalletName = name

	// The seed phrase will be used to generate wallet
	ewl.SeedPhrase = seed

	// Validation seed phrase
	if !evenWallet.ValidatePhrase(ewl.SeedPhrase) {
		return errors.New(invalidPhraseError)
	}

	// Provide password
	ewl.Password = "zzzzzz"

	_, err := ewl.Create()
	if err != nil {
		return err
	}

	return nil
}

func (w *wallet) save(password string) error {
	fmt.Printf("Wallet saved with password '%s'\n", password)
	return nil
}
