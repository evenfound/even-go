package hdwallet

import (
	"errors"
	"fmt"
	btcdwallet "github.com/btcsuite/btcwallet/wallet"
	_ "github.com/btcsuite/btcwallet/walletdb/bdb"
	"os"
	"path/filepath"
	"time"
)

const (
	DefaultWalletDirectory = "wallets"
	DefaultPrivateDataName = "private.key"

	ErrorPrivateKeyCreating = "Can't create private data in %v"

	StatusSuccessfullyCreated = "Wallet created successfully."
)

var AvailableCoinTypes = []string{"BTC", "LTC", "ETH"}

type HDWallet struct {
	SeedPhrase string `json:"seed" short:"s" long:"seed" description:"Seed phrase of the  wallet"`
	WalletAuth
}

// Create function receives a data to create and initialize wallet
// and return instance of HDWallet and/or error
func (wallet *HDWallet) Create() (*btcdwallet.Wallet, error) {

	var loader, err = GetLoader(wallet.TestNet, wallet.WalletName)

	if err != nil {
		return nil, err
	}

	newWallet, err := loader.CreateNewWallet([]byte(wallet.Password), []byte(wallet.SeedPhrase), nil, time.Now())

	defer newWallet.Manager.Close()

	if err != nil {
		return nil, err
	}

	var dbpath, _ = GeneratePath(wallet.TestNet, filepath.Join(wallet.WalletName, DefaultPrivateDataName))

	if dbpath == "" {
		return newWallet, errors.New(fmt.Sprintf(ErrorPrivateKeyCreating, dbpath))
	}

	f, _ := os.Create(dbpath)

	defer f.Close()

	f.Write(Encrypt([]byte(wallet.SeedPhrase), wallet.Password))

	return newWallet, nil
}
