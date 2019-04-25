package wallet

import (
	"time"

	"github.com/OpenBazaar/multiwallet"
	"github.com/OpenBazaar/wallet-interface"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/evenfound/even-go/node/repo"
	"github.com/evenfound/even-go/node/repo/db"
	"github.com/evenfound/even-go/node/schema"
	"github.com/op/go-logging"
	"golang.org/x/net/proxy"
)

// WalletConfig describes the options needed to create a MultiWallet
type WalletConfig struct {
	// ConfigFile contains the options of each native wallet
	ConfigFile *schema.WalletsConfig
	// RepoPath is the base path which contains the nodes data directory
	RepoPath string
	// Logger is an interface to support internal wallet logging
	Logger logging.Backend
	// DB is an interface to support internal transaction persistance
	DB *db.DB
	// Mnemonic is the string entropy used to generate the wallet's BIP39-compliant seed
	Mnemonic string
	// WalletCreationDate represents the time when new transactions were added by this wallet
	WalletCreationDate time.Time
	// Params describe the desired blockchain params to enforce on joining the network
	Params *chaincfg.Params
	// Proxy is an interface which allows traffic for the wallet to proxied
	Proxy proxy.Dialer
	// DisableExchangeRates will disable usage of the internal exchange rate API
	DisableExchangeRates bool
}

// NewMultiWallet returns a functional set of wallets using the provided WalletConfig.
// The value of schema.WalletsConfig.<COIN>.Type must be "API" or will be ignored. BTC
// and BCH can also use the "SPV" Type.
func NewMultiWallet() (multiwallet.MultiWallet, error) {
	return make(multiwallet.MultiWallet), nil
}

type WalletDatastore struct {
	keys           repo.KeyStore
	stxos          repo.SpentTransactionOutputStore
	txns           repo.TransactionStore
	utxos          repo.UnspentTransactionOutputStore
	watchedScripts repo.WatchedScriptStore
}

func (d *WalletDatastore) Keys() wallet.Keys {
	return d.keys
}
func (d *WalletDatastore) Stxos() wallet.Stxos {
	return d.stxos
}
func (d *WalletDatastore) Txns() wallet.Txns {
	return d.txns
}
func (d *WalletDatastore) Utxos() wallet.Utxos {
	return d.utxos
}
func (d *WalletDatastore) WatchedScripts() wallet.WatchedScripts {
	return d.watchedScripts
}
