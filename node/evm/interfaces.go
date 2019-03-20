package evm

// Interface is the abstract interface of the EVM.
type Interface interface {
	Run(bc Bytecode) error
	Interpret(sc string) error
}
