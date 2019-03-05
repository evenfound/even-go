package wallet

import (
	"fmt"
	"github.com/evenfound/even-go/node/hdwallet"
)

type Wallet int

// Creating a new wallet
func (_ *Wallet) Create(w hdwallet.HDWallet, reply *JsonResponse) error {

	fmt.Println("Creating wallet")

	var _, err = w.Create()

	if err != nil {
		reply.Error(err)
	} else {
		reply.Render(w)
	}

	return nil
}
