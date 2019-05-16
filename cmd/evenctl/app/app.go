package app

import (
	"os"

	"github.com/evenfound/even-go/node/cmd/evenctl/config"
	"github.com/evenfound/even-go/node/cmd/evenctl/rpc"
	"github.com/jawher/mow.cli"
)

const (
	walletSpec  = "--name ... --password ..."
	accountSpec = "--name ... --password ... --account ..."
)

var fcmd, _ = rpc.NewFileCMD()
var pcmd, _ = rpc.NewPeerCMD()

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

	a.BoolOptPtr(&config.Debug, "d debug", false, "show additional information")

	a.Command("test", "test", func(config *cli.Cmd) {
		config.Command("call", "call smart contract", cmdTestCall)
		config.Command("sign", "sign message", cmdTestSign)
		config.Command("verify", "verify signed message", cmdTestVerify)
		config.Command("tx", "transactions", func(config *cli.Cmd) {
			config.Command("create", "create new transaction", cmdTestCreateTx)
			config.Command("show", "read and show transaction", cmdTestShowTx)
			config.Command("analyze", "read and analyze transaction", cmdTestAnalyzeTx)
			config.Command("verify", "read and verify transaction", cmdTestVerifyTx)
		})
	})

	a.Command("wallet", "manage wallets", func(config *cli.Cmd) {
		config.Command("generate", "create new unique wallet", cmdWalletGenerate)
		config.Command("create", "(re)create a wallet with known seed", cmdWalletCreate)
		config.Command("unlock", "unlock wallet temporarily", cmdWalletUnlock)
		config.Command("nextaccount", "generate next account", cmdWalletNextAccount)
		config.Command("privkey", "show private key of account", cmdAccountPrivateKey)
		config.Command("pubkey", "show public key of account", cmdAccountPublicKey)
		config.Command("balance", "show current balance of account", cmdAccountBalance)
		config.Command("info", "show some information about wallet", cmdWalletInfo)
		config.Command("tx", "transactions", func(config *cli.Cmd) {
			config.Command("newreg", "create initial transaction", cmdWalletTxNewReg)
			config.Command("contract", "create contract-deploy transaction", cmdWalletTxContract)
			config.Command("invoke", "create contract-invoke transaction", cmdWalletTxContractInvoke)
		})
	})

	a.Command("file", "manage files", func(cmd *cli.Cmd) {
		cmd.Command("create", "create new file", cmdCreateFile)
		cmd.Command("mkdir", "create new directory", cmdFilesMkdir)
		cmd.Command("find", "find file by hash", cmdFileFind)
		cmd.Command("stat", "file stat information", cmdFileStat)
	})

	a.Command("peer", "manage peers", func(cmd *cli.Cmd) {
		cmd.Command("list", "peer list", cmdPeerList)
		cmd.Command("send", "send store to peers", cmdPeerSend)
	})

	//TODO create cmd command for files and peers

	return a.Run(os.Args)
}

func cmdTestCall(c *cli.Cmd) {
	var (
		file  = c.StringOpt("f file", "", "name of smart contract file")
		entry = c.StringOpt("e entry", config.DefaultEntryFunction, "name of entry function")
	)
	c.Spec = "--file ... [--entry ...]"
	c.Action = func() {
		must(rpc.Call(*file, *entry))
	}
}

func cmdTestSign(c *cli.Cmd) {
	var (
		message = c.StringArg("MESSAGE", "", "arbitrary message")
		privkey = c.StringOpt("k privkey", "", "private key")
	)
	c.Spec = "MESSAGE --privkey ..."
	c.Action = func() {
		must(rpc.Sign(*message, *privkey))
	}
}

func cmdTestVerify(c *cli.Cmd) {
	var (
		message   = c.StringArg("MESSAGE", "", "message")
		signature = c.StringOpt("s signature", "", "signature")
		pubkey    = c.StringOpt("k pubkey", "", "public key")
	)
	c.Spec = "MESSAGE --signature ... --pubkey ..."
	c.Action = func() {
		must(rpc.Verify(*message, *signature, *pubkey))
	}
}

func cmdTestCreateTx(c *cli.Cmd) {
	var (
		format = c.StringOpt("f format", "", "file format")
	)
	c.Spec = "--format ..."
	c.Action = func() {
		must(rpc.CreateTransaction(*format))
	}
}

func cmdTestShowTx(c *cli.Cmd) {
	var (
		file = c.StringArg("FILE", "", "filename")
	)
	c.Spec = "FILE"
	c.Action = func() {
		must(rpc.ShowTransaction(*file))
	}
}

func cmdTestAnalyzeTx(c *cli.Cmd) {
	var (
		file = c.StringArg("FILE", "", "filename")
	)
	c.Spec = ""
	c.Action = func() {
		must(rpc.AnalyzeTransaction(*file))
	}
}

func cmdTestVerifyTx(c *cli.Cmd) {
	var (
		file = c.StringArg("FILE", "", "filename")
	)
	c.Spec = ""
	c.Action = func() {
		must(rpc.VerifyTransaction(*file))
	}
}

func cmdWalletGenerate(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
	)
	c.Spec = walletSpec
	c.Action = func() {
		must(rpc.GenerateWallet(*name, *password))
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
		must(rpc.CreateWallet(*name, *mnemonic, *password))
	}
}

func cmdWalletUnlock(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
	)
	c.Spec = walletSpec
	c.Action = func() {
		must(rpc.UnlockWallet(*name, *password))
	}
}

func cmdWalletNextAccount(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
	)
	c.Spec = walletSpec
	c.Action = func() {
		must(rpc.WalletNextAccount(*name, *password))
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
		must(rpc.WalletAccountDumpPrivateKey(*name, *password, *account))
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
		must(rpc.WalletAccountDumpPublicKey(*name, *password, *account))
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
		must(rpc.WalletAccountShowBalance(*name, *password, *account))
	}
}

func cmdWalletInfo(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
	)
	c.Spec = walletSpec
	c.Action = func() {
		must(rpc.GetWalletInfo(*name, *password))
	}
}

func cmdWalletTxNewReg(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
		account  = c.StringOpt("a account", "", "address of account")
	)
	c.Spec = accountSpec
	c.Action = func() {
		must(rpc.WalletAccountTxNewReg(*name, *password, *account))
	}
}

func cmdWalletTxContract(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
		account  = c.StringOpt("a account", "", "address of account")
		contract = c.StringOpt("c contract", "", "IPFS-hash of contract")
	)
	c.Spec = "--name ... --password ... --account ... --contract ..."
	c.Action = func() {
		must(rpc.WalletAccountTxContract(*name, *password, *account, *contract))
	}
}

func cmdWalletTxContractInvoke(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of wallet")
		password = c.StringOpt("p password", "", "password")
		account  = c.StringOpt("a account", "", "address of account")
		contract = c.StringOpt("c contract", "", "IPFS-hash of contract")
		function = c.StringOpt("f function", "", "name of function to invoke (default if omitted)")
	)
	c.Spec = "--name ... --password ... --account ... --contract ... [--function ...]"
	c.Action = func() {
		must(rpc.WalletAccountTxContractInvoke(*name, *password, *account,
			*contract, *function))
	}
}

func cmdCreateFile(c *cli.Cmd) {
	var (
		name     = c.StringOpt("n name", "", "name of file")
		password = c.StringOpt("s source", "", "source file")
	)

	c.Spec = "--name ... --source ..."

	c.Action = func() {

		must(fcmd.Create(*name, *password))
	}
}

func cmdFileStat(c *cli.Cmd) {
	var (
		name = c.StringOpt("n name", "", "name of file")
	)

	c.Spec = "--name"

	c.Action = func() {

		must(fcmd.Stat(*name))
	}
}

func cmdFilesMkdir(c *cli.Cmd) {
	var (
		name = c.StringOpt("n name", "", "name of file")
	)

	c.Spec = "--name"

	c.Action = func() {

		must(fcmd.Mkdir(*name))
	}
}

func cmdFileFind(c *cli.Cmd) {
	var (
		hash   = c.StringOpt("h hash", "", "hash of the file")
		output = c.StringOpt("o output", "", "full name of the output file")
	)

	c.Spec = "--hash ... -output ..."

	c.Action = func() {

		must(fcmd.GetFileByHash(*hash, *output))
	}
}

func cmdPeerList(c *cli.Cmd) {
	c.Spec = ""

	c.Action = func() {

		must(pcmd.List())
	}
}

func cmdPeerSend(c *cli.Cmd) {

	var (
		hash = c.StringOpt("h hash", "", "hash of the file")
	)

	c.Spec = "--hash ..."

	c.Action = func() {

		must(pcmd.SendStore(*hash))
	}
}
