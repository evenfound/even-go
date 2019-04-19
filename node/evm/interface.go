package evm

// Interface is the abstract interface of the Even VM.
type Interface interface {
	// Run executes the bytecode with call of the entryFunc.
	// Returns resulting string.
	Run(bc Bytecode, entryFunc string) (string, error)
}

// New creates another instance of the EVM.
func New() Interface {
	return newTengoVM()
}
