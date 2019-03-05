package wallet

import "github.com/evenfound/even-go/node/hdwallet"

type Address int

func (_ *Address) Create(manage hdwallet.AddressManager, response *JsonResponse) error {

	var wallet, err = manage.Authorize()

	if err != nil {
		response.Error(err)
		return err
	}

	manage.SetWallet(wallet)

	addresses, err := manage.NewAddress()

	if err != nil {
		response.Error(err)
		return err
	}

	response.Render(addresses)

	return nil
}
