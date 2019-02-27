package cmd

import (
	"fmt"
	"github.com/evenfound/even-go/node/hdwallet"
)

const (
	MessageAccountCreated = "Created account with name [%v] : ID [%v]"
)

type CreateAccount struct {
	hdwallet.AccountManager
}

// Creating a new account in specified wallet
func (account *CreateAccount) Execute(args []string) error {

	var wallet, err = account.Authorize()

	if err != nil {
		return err
	}

	accountId, err := account.NewAccount(*wallet)

	if err == nil {
		fmt.Println(fmt.Sprintf(MessageAccountCreated, account.AccountName, accountId))
	}

	return err
}
