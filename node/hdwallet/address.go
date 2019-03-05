package hdwallet

import (
	"github.com/btcsuite/btcwallet/waddrmgr"
	btcdWallet "github.com/btcsuite/btcwallet/wallet"
	"github.com/btcsuite/btcwallet/walletdb"
)

type AddressManager struct {
	Account string `short:"a" long:"account" description:"Account name" json:"account"` 

	WalletAuth

	Level int `short:"l" long:"level" description:"Address generation level"`
	Index int `short:"i" long:"index" description:"Address generation index"`

	wallet *btcdWallet.Wallet
}

// This function will set the wallet in AddressManager
// If the wallet doesn't set than will be empty wallet instead
func (am *AddressManager) SetWallet(wallet *btcdWallet.Wallet) {
	am.wallet = wallet
}

// This function generates a new address based an account and the wallet
// Also considering coin type
func (am *AddressManager) NewAddress() (addresses []string, err error) {

	var (
		address string
		w       = am.wallet
	)

	err = walletdb.Update(w.Database(), func(tx walletdb.ReadWriteTx) error {
		addrmgrNs := tx.ReadWriteBucket(AccountNamespaceKey)
		address, _, err = am.address(addrmgrNs)
		addresses = append(addresses, address)
		return err
	})

	if err != nil {
		return
	}

	return
}

// This function generates an address.
// To generate address this function will use account
// and wallet credentials. Than using scope parametrs generate addresses based on coin type
func (am *AddressManager) address(addrmgrNs walletdb.ReadWriteBucket) (string, *waddrmgr.AccountProperties, error) {

	var (
		w     = am.wallet
		scope = am.GetKeyScope()
	)

	var account, err = w.AccountNumber(scope, am.Account)

	if err != nil {
		return "", nil, err
	}

	manager, err := w.Manager.FetchScopedKeyManager(scope)

	if err != nil {
		return "", nil, err
	}

	// Get next address from wallet.
	addrs, err := manager.NextExternalAddresses(addrmgrNs, account, 1)

	if err != nil {
		return "", nil, err
	}

	props, err := manager.AccountProperties(addrmgrNs, account)
	if err != nil {
		return "", nil, err
	}

	return addrs[0].Address().String(), props, nil
}
