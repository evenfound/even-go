package cmd

import (
	"fmt"
	"github.com/evenfound/even-go/node/hdwallet"
)

const (
	AccountFoundResult = "Name  [%v] :  ID  [%v]"
)

type ListAccounts struct {
	hdwallet.WalletAuth
}

// Getting all accounts of wallet
func (la *ListAccounts) Execute(args []string) error {

	var wallet, err = la.Authorize()

	if err != nil {
		return err
	}

	var scope = la.GetKeyScope()

	accounts, err := wallet.Accounts(scope)

	if err != nil {
		return err
	}

	for _, acc := range accounts.Accounts {
		fmt.Println(fmt.Sprintf(AccountFoundResult, acc.AccountName, acc.AccountNumber))
	}

	return nil

}
