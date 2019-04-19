package app

import (
	"os"

	"github.com/evenfound/even-go/node/cmd/evec/tool"
	"github.com/evenfound/even-go/node/cmd/evenctl/config"
	"github.com/evenfound/even-go/node/cmd/evenctl/rpc"
	"github.com/jawher/mow.cli"
)

const (
	walletSpec  = "--name ... --password ..."
	accountSpec = "--name ... --password ... --account ..."
)

// Init initializes the application.
func Init() {
}

// Close finalizes the application.
func Close() {
}

// Run starts the application.
func Run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	a := cli.App("evenctl", "Even Network control tool.")

	config.Debug = *a.BoolOpt("d debug", false, "show additional information")

	a.Command("test", "test", func(config *cli.Cmd) {
		config.Command("call", "call a smart contract", cmdTestCall)
	})

	a.Command("wallet", "manage wallets", func(config *cli.Cmd) {
		config.Command("generate", "create new unique wallet", cmdWalletGenerate)
		config.Command("create", "(re)create a wallet with known seed", cmdWalletCreate)
		config.Command("unlock", "unlock password temporarily", cmdWalletUnlock)
		config.Command("nextaccount", "generate next account", cmdWalletNextAccount)
		config.Command("privkey", "show private key of account", cmdAccountPrivateKey)
		config.Command("pubkey", "show public key of account", cmdAccountPublicKey)
		config.Command("balance", "show current balance of account", cmdAccountBalance)
		config.Command("info", "show some information about wallet", cmdWalletInfo)
	})

	return a.Run(os.Args)
}

func cmdTestCall(c *cli.Cmd) {
	var (
		file  = c.StringOpt("f file", "", "name of smart contract file")
		entry = c.StringOpt("e entry", config.DefaultEntryFunction, "name of entry function")
	)
	c.Spec = "--file ... [--entry ...]"
	c.Action = func() {
		tool.Must(rpc.Call(*file, *entry))
	}
}

func cmdWalletGenerate(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
	)
	c.Spec = walletSpec
	c.Action = func() {
		tool.Must(rpc.GenerateWallet(*name, *password))
	}
}

func cmdWalletCreate(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
		mnemonic = c.StringOpt("s seed", "", "mnemonic seed phrase")
	)
	c.Spec = "--name ... --password ... --seed ..."
	c.Action = func() {
		tool.Must(rpc.CreateWallet(*name, *mnemonic, *password))
	}
}

func cmdWalletUnlock(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
	)
	c.Spec = walletSpec
	c.Action = func() {
		tool.Must(rpc.UnlockWallet(*name, *password))
	}
}

func cmdWalletNextAccount(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
	)
	c.Spec = walletSpec
	c.Action = func() {
		tool.Must(rpc.WalletNextAccount(*name, *password))
	}
}

func cmdAccountPrivateKey(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
		account  = c.StringOpt("a account", "", "address of account")
	)
	c.Spec = accountSpec
	c.Action = func() {
		tool.Must(rpc.WalletAccountDumpPrivateKey(*name, *password, *account))
	}
}

func cmdAccountPublicKey(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
		account  = c.StringOpt("a account", "", "address of account")
	)
	c.Spec = accountSpec
	c.Action = func() {
		tool.Must(rpc.WalletAccountDumpPublicKey(*name, *password, *account))
	}
}

func cmdAccountBalance(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
		account  = c.StringOpt("a account", "", "address of account")
	)
	c.Spec = accountSpec
	c.Action = func() {
		tool.Must(rpc.WalletAccountShowBalance(*name, *password, *account))
	}
}

func cmdWalletInfo(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
	)
	c.Spec = walletSpec
	c.Action = func() {
		tool.Must(rpc.GetWalletInfo(*name, *password))
	}
}
