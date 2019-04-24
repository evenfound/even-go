package interop

import (
	"github.com/d5/tengo/script"
)

// Environment represents a smart contract environment.
type Environment struct {
	runner  *script.Compiled
	wallets []wallet
	strings []string
}

// Run runs the compiled script.
func (e *Environment) Run() error {
	return e.runner.Run()
}

// Get returns a variable identified by name.
func (e *Environment) Get(name string) *script.Variable {
	return e.runner.Get(name)
}

func (e *Environment) addWallet(w *wallet) handle {
	e.wallets = append(e.wallets, *w)
	return handle(len(e.wallets) - 1)
}

func (e *Environment) wallet(h handle) *wallet {
	return &e.wallets[h]
}
