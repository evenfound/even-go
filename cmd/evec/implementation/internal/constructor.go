package internal

import (
	"github.com/evenfound/even-go/cmd/evec/compiler"
)

// New creates another instance of internal (Tengo) compiler.
func New() compiler.Interface {
	return tengoCompiler{}
}
