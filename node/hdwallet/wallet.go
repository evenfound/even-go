package hdwallet

import (
	"github.com/btcsuite/btcwallet/waddrmgr"
	btcwallet "github.com/btcsuite/btcwallet/wallet"
)

type WalletAuth struct {
	WalletName string `short:"n" long:"name" description:"Name of the wallet"`
	Password   string `short:"p" long:"password" description:"Password of the wallet"`
	Coin       uint32 `short:"c" long:"coin" description:"Coin type"`
	TestNet    bool   `short:"t" long:"testnet" description:"TesNet network"`
}

// This function will authorize wallet based on name and password
// If the loader can't find the wallet will be returned an error
// Else the function will be return an existing wallet
func (wallet *WalletAuth) Authorize() (*btcwallet.Wallet, error) {

	var loader, err = GetLoader(wallet.TestNet, wallet.WalletName)

	if err != nil {
		return nil, err
	}

	return loader.OpenExistingWallet([]byte(wallet.Password), false)
}

// This function return kyscope data based on BIP and coin id
// ( we are always uisng BIP44 )
func (wallet *WalletAuth) GetKeyScope() waddrmgr.KeyScope {
	return waddrmgr.KeyScope{
		Purpose: 44,
		Coin:    wallet.Coin,
	}
}

func (wallet WalletAuth) Path() {
	// TODO  - this function will return full path to data directory
}
