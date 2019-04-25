package hdwallet

import "time"

// Interface represents a HD wallet.
type Interface interface {
	// Generate generates new HD wallet and returns the seed.
	Generate() (string, error)

	// Create (re)creates wallet.
	Create(mnemonic string) error

	// Unlock allows wallet operations without password for some time.
	Unlock(duration time.Duration) error

	// NextAccount generates next deterministic account.
	NextAccount() (string, error)

	// PrivateKey returns the private key of account.
	PrivateKey(address string) (string, error)

	// PublicKey returns the public key of account.
	PublicKey(address string) (string, error)

	// Balance returns the current balance of account.
	Balance(address string) (string, error)

	// GetInfo retrieves some information about wallet.
	GetInfo() (string, error)

	// TxNewReg creates initial transaction.
	TxNewReg(address string) (string, error)
}

// New creates another instance of wallet interface.
func New(name, password string) Interface {
	return newWallet(name, password)
}
