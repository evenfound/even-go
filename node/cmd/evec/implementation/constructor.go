package implementation

import (
	"github.com/evenfound/even-go/node/cmd/evec/compiler"
	"github.com/evenfound/even-go/node/cmd/evec/config"
	"github.com/evenfound/even-go/node/cmd/evec/implementation/evelyn"
	"github.com/evenfound/even-go/node/cmd/evec/implementation/internal"
	"github.com/evenfound/even-go/node/cmd/evec/implementation/solidity"
	"github.com/evenfound/even-go/node/cmd/evec/implementation/vyper"
)

// New creates another instance of the compiler.
func New(ext string) compiler.Interface {
	switch {
	case config.BuildTengo:
		return internal.New()
	case config.BuildEvelyn:
		return evelyn.New()
	case config.BuildVyper:
		return vyper.New()
	case config.BuildSolidity:
		return solidity.New()
	}

	switch ext {
	case config.TengoExt:
		return internal.New()
	case config.EvelynExt:
		return evelyn.New()
	case config.VyperExt:
		return vyper.New()
	case config.SolidityExt:
		return solidity.New()
	}

	return nil
}
