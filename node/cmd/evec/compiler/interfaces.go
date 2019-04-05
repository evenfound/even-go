package compiler

import (
	"io"

	"github.com/d5/tengo/objects"
)

// Bytecode represents a binary byte code for a VM.
type Bytecode interface {
	Decode(r io.Reader, m *objects.ModuleMap) error
	Encode(w io.Writer) error
	CountObjects() int
	FormatInstructions() []string
	FormatConstants() (output []string)
}

// Interface represents a compiler.
type Interface interface {
	Compile(filename string) (Bytecode, error)
	TryCompile(filename string) ([]byte, error)
}
