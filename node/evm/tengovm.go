package evm

import (
	"bytes"
	"errors"

	"github.com/d5/tengo/script"

	"github.com/d5/tengo/compiler"
	"github.com/d5/tengo/runtime"
)

var _ Interface = tengoVM{}

// newTengoVM creates new instance of the tengoVM.
func newTengoVM() Interface {
	vm := tengoVM{}
	return vm
}

// tengoVM represents the Tengo VM.
type tengoVM struct {
}

// Run implements corresponding method of the EVM interface.
func (tengoVM) Run(bc Bytecode) error {
	return runCompiled(bc)
}

// Interpret implements corresponding method of the EVM interface.
func (tengoVM) Interpret(sc string) error {
	s := script.New([]byte(sc))
	if _, err := s.Run(); err != nil {
		return err
	}
	return nil
}

func runCompiled(data Bytecode) error {
	bc := &compiler.Bytecode{}
	err := bc.Decode(bytes.NewReader([]byte(data)))
	if err != nil {
		return errors.New("invalid bytecode")
	}

	machine := runtime.NewVM(bc, nil, nil, nil)

	err = machine.Run()
	if err != nil {
		return err
	}

	return nil
}
