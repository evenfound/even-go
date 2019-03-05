package wallet

import "github.com/evenfound/even-go/node/hdwallet"

type Account int

type CreatedAccountResponse struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type ListAccountResponse []CreatedAccountResponse

// Creating a new account
func (account *Account) Create(manager hdwallet.AccountManager, response *JsonResponse) error {

	var wallet, err = manager.Authorize()

	if err != nil {
		response.Error(err)
		return err
	}

	accountId, err := manager.NewAccount(*wallet)

	if err != nil {
		response.Error(err)
		return err
	}

	response.Render(CreatedAccountResponse{
		Id:   accountId,
		Name: manager.AccountName,
	})

	return nil
}

// Getting list accounts
func (account *Account) List(wallet hdwallet.HDWallet, response *JsonResponse) error {

	var btcWallet, err = wallet.Authorize()

	if err != nil {
		response.Error(err)
		return err
	}

	accounts, err := btcWallet.Accounts(wallet.GetKeyScope())

	if err != nil {
		response.Error(err)
		return err
	}

	var list ListAccountResponse

	for _, account := range accounts.Accounts {
		list = append(list, CreatedAccountResponse{
			Id:   account.AccountNumber,
			Name: account.AccountName,
		})
	}

	response.Render(list)

	return nil
}
