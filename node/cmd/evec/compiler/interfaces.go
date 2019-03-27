package compiler

import "io"

// Bytecode represents a binary byte code for a VM.
type Bytecode interface {
	Decode(r io.Reader) error
	Encode(w io.Writer) error
	CountObjects() int
	FormatInstructions() []string
	FormatConstants() (output []string)
}

// Interface represents a compiler.
type Interface interface {
	Compile(filename string) (Bytecode, error)
}
