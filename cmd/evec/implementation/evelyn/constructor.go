package evelyn

import (
	"github.com/evenfound/even-go/cmd/evec/compiler"
)

// New creates another instance of the Evelyn compiler.
func New() compiler.Interface {
	return evelynCompiler{}
}
