package hdwallet

import (
	"encoding/json"
	"errors"
	"time"
)

// wallet implements Interface.
type wallet struct {
	Name     string `json:"name"`
	Mnemonic string `json:"mnemonic"`
	password string
	Accounts []account `json:"accounts"`
}

// newWallet creates another instance of wallet.
func newWallet(name, pass string) *wallet {
	return &wallet{
		Name:     name,
		password: pass,
		Accounts: make([]account, 0),
	}
}

// Generate implements Interface.Generate.
func (w *wallet) Generate() (string, error) {
	mnemonic := generateMnemonic()
	if err := w.Create(mnemonic); err != nil {
		return "", err
	}
	return mnemonic, nil
}

// Create implements Interface.Create.
func (w *wallet) Create(mnemonic string) error {
	w.Mnemonic = mnemonic
	err := w.save()
	if err != nil {
		return err
	}
	return nil
}

// Unlock implements Interface.Unlock.
func (w *wallet) Unlock(duration time.Duration) error {
	return errors.New("not yet implemented")
}

// NextAccount implements Interface.NextAccount.
func (w *wallet) NextAccount() (string, error) {
	if err := w.load(); err != nil {
		return "", err
	}

	index := w.nextAccountIndex()
	address, err := generateAddress(w.Mnemonic, w.password, index)
	if err != nil {
		return "", err
	}

	w.addAccount(newAccount(index, address))
	if err = w.save(); err != nil {
		return "", err
	}

	return address, nil
}

// PrivateKey implements Interface.PrivateKey.
func (w *wallet) PrivateKey(address string) (string, error) {
	if err := w.load(); err != nil {
		return "", err
	}

	acc := w.findAccount(address)
	if acc == nil {
		return "", errors.New("account not found")
	}

	return generatePrivateKey(w.Mnemonic, w.password, acc.Index)
}

// PublicKey implements Interface.PublicKey.
func (w *wallet) PublicKey(address string) (string, error) {
	if err := w.load(); err != nil {
		return "", err
	}

	acc := w.findAccount(address)
	if acc == nil {
		return "", errors.New("account not found")
	}

	return generatePublicKey(w.Mnemonic, w.password, acc.Index)
}

// Balance implements Interface.Balance.
func (w *wallet) Balance(address string) (string, error) {
	return "0", nil
}

// GetInfo implements Interface.GetWalletInfo.
func (w *wallet) GetInfo() (string, error) {
	filename, err := absoluteFilename(w.Name)
	if err != nil {
		return "", err
	}

	js, err := readEncrypted(filename, w.password)
	if err != nil {
		return "", err
	}

	return string(js), nil
}

// save saves wallet as single binary encrypted file.
func (w *wallet) save() error {
	filename, err := absoluteFilename(w.Name)
	if err != nil {
		return err
	}

	js, err := json.Marshal(w)
	if err != nil {
		return err
	}

	err = writeEncrypted(filename, w.password, js)
	if err != nil {
		return err
	}

	return nil
}

// load restores wallet from a file.
func (w *wallet) load() error {
	filename, err := absoluteFilename(w.Name)
	if err != nil {
		return err
	}

	js, err := readEncrypted(filename, w.password)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, w)
	if err != nil {
		return err
	}

	return nil
}

func (w *wallet) nextAccountIndex() uint32 {
	return uint32(len(w.Accounts))
}

func (w *wallet) addAccount(a *account) {
	w.Accounts = append(w.Accounts, *a)
}

func (w *wallet) findAccount(address string) *account {
	for _, a := range w.Accounts {
		if a.Address == address {
			return &a
		}
	}
	return nil
}
