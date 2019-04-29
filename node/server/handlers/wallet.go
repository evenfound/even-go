package handlers

import (
	"time"

	"github.com/evenfound/even-go/node/hdwallet"
	"github.com/evenfound/even-go/node/server/api"
	"golang.org/x/net/context"
)

// Wallet is a wallet service handler.
type Wallet struct{}

// GenerateWallet generates new wallet.
func (sc *Wallet) GenerateWallet(ctx context.Context, in *api.WalletInput) (*api.WalletResult, error) {
	wallet := hdwallet.New(in.Name, in.Password)
	mnemonic, err := wallet.Generate()
	if err != nil {
		return newWalletResult(false, err.Error()), nil
	}
	return newWalletResult(true, mnemonic), nil
}

// CreateWallet creates wallet.
func (sc *Wallet) CreateWallet(ctx context.Context, in *api.CreateWalletInput) (*api.WalletResult, error) {
	wallet := hdwallet.New(in.Name, in.Password)
	err := wallet.Create(in.Mnemonic)
	if err != nil {
		return newWalletResult(false, err.Error()), nil
	}
	return newWalletResult(true, in.Name), nil
}

// UnlockWallet allows wallet operations without password.
func (sc *Wallet) UnlockWallet(ctx context.Context, in *api.WalletInput) (*api.WalletResult, error) {
	wallet := hdwallet.New(in.Name, in.Password)
	err := wallet.Unlock(600 * time.Second)
	if err != nil {
		return newWalletResult(false, err.Error()), nil
	}
	return newWalletResult(true, "600s"), nil
}

// WalletNextAccount generates next deterministic account.
func (sc *Wallet) WalletNextAccount(ctx context.Context, in *api.WalletInput) (*api.WalletResult, error) {
	wallet := hdwallet.New(in.Name, in.Password)
	address, err := wallet.NextAccount()
	if err != nil {
		return newWalletResult(false, err.Error()), nil
	}
	return newWalletResult(true, address), nil
}

// WalletAccountDumpPrivateKey outputs the private key of account.
func (sc *Wallet) WalletAccountDumpPrivateKey(ctx context.Context, in *api.WalletAccountInput) (*api.WalletResult, error) {
	wallet := hdwallet.New(in.Name, in.Password)
	privKey, err := wallet.PrivateKey(in.Account)
	if err != nil {
		return newWalletResult(false, err.Error()), nil
	}
	return newWalletResult(true, privKey), nil
}

// WalletAccountDumpPublicKey outputs the public key of account.
func (sc *Wallet) WalletAccountDumpPublicKey(ctx context.Context, in *api.WalletAccountInput) (*api.WalletResult, error) {
	wallet := hdwallet.New(in.Name, in.Password)
	pubKey, err := wallet.PublicKey(in.Account)
	if err != nil {
		return newWalletResult(false, err.Error()), nil
	}
	return newWalletResult(true, pubKey), nil
}

// WalletAccountShowBalance outputs the current balance of account.
func (sc *Wallet) WalletAccountShowBalance(ctx context.Context, in *api.WalletAccountInput) (*api.WalletResult, error) {
	wallet := hdwallet.New(in.Name, in.Password)
	balance, err := wallet.Balance(in.Account)
	if err != nil {
		return newWalletResult(false, err.Error()), nil
	}
	return newWalletResult(true, balance), nil
}

// GetWalletInfo retrieves some information about wallet.
func (sc *Wallet) GetWalletInfo(ctx context.Context, in *api.WalletInput) (*api.WalletResult, error) {
	wallet := hdwallet.New(in.Name, in.Password)
	info, err := wallet.GetInfo()
	if err != nil {
		return newWalletResult(false, err.Error()), nil
	}
	return newWalletResult(true, info), nil
}

// WalletAccountTxNewReg creates initial transaction.
func (sc *Wallet) WalletAccountTxNewReg(ctx context.Context, in *api.WalletAccountInput) (*api.WalletResult, error) {
	wallet := hdwallet.New(in.Name, in.Password)
	hash, err := wallet.TxNewReg(in.Account)
	if err != nil {
		return newWalletResult(false, err.Error()), nil
	}
	return newWalletResult(true, hash), nil
}

func newWalletResult(ok bool, msg string) *api.WalletResult {
	return &api.WalletResult{
		Ok:     ok,
		Result: msg,
	}
}
