package cmd

import (
	"github.com/evenfound/even-go/node/hdwallet"
)

type NewAddress struct {
	hdwallet.AddressManager
}

// Generating a new address in specified account
func (generator *NewAddress) Execute(args []string) error {

	var wallet, err = generator.Authorize()

	if err != nil {
		return err
	}

	generator.SetWallet(wallet)

	generator.NewAddress()

	return nil
}
