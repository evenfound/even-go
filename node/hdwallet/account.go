package hdwallet

import (
	"errors"
	"fmt"
	"github.com/btcsuite/btcwallet/waddrmgr"
	btcWallet "github.com/btcsuite/btcwallet/wallet"
	"github.com/btcsuite/btcwallet/walletdb"
	"io/ioutil"
	"path/filepath"
)

type AccountManager struct {
	WalletAuth
	AccountName string `short:"a" long:"account" description:"Account name" json:"account"`
}

var (
	AccountNamespaceKey = []byte("waddrmgr")
)

// This function creates a new account based on wallet and already specified name
// For first you nedd to unlock the wallet with passpriv
// If any error caused during this process , the function will return 0 and an error
func (generator AccountManager) NewAccount(wallet btcWallet.Wallet) (uint32, error) {

	var (
		account uint32
		props   *waddrmgr.AccountProperties
	)

	dberr := walletdb.Update(wallet.Database(), func(tx walletdb.ReadWriteTx) error {

		var scope = generator.GetKeyScope()

		addrmgrNs := tx.ReadWriteBucket(AccountNamespaceKey)

		var dbpath, _ = GeneratePath(generator.TestNet, filepath.Join(generator.WalletName, DefaultPrivateDataName))

		data, _ := ioutil.ReadFile(dbpath)

		var privatePassword = Decrypt(data, generator.Password)

		err := wallet.Manager.Unlock(addrmgrNs, []byte(privatePassword))

		if err != nil {
			return err
		}

		manager, err := wallet.Manager.FetchScopedKeyManager(scope)

		defer manager.Close()

		if err != nil {
			return err
		}

		account, err = manager.NewAccount(addrmgrNs, generator.AccountName)

		if err != nil {
			return err
		}

		props, err = manager.AccountProperties(addrmgrNs, account)

		return err
	})

	if dberr != nil {
		return 0, errors.New(fmt.Sprintf("Cannot fetch new account properties for notification "+
			"after account creation: %v", dberr))
	} else {
		return props.AccountNumber, nil
	}
}
