package evm

// Interface is the abstract interface of the Even VM.
type Interface interface {
	Run(bc Bytecode) error
	Interpret(sc string) error
}
