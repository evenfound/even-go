package internal

import (
	"io/ioutil"
	"path/filepath"

	"github.com/evenfound/even-go/node/cmd/evec/compiler"
	"github.com/evenfound/even-go/node/evm/interop"
	"github.com/pkg/errors"

	tengo "github.com/d5/tengo/compiler"
	tengoParser "github.com/d5/tengo/compiler/parser"
	tengoSource "github.com/d5/tengo/compiler/source"
)

type tengoCompiler struct {
}

// Compile translates a source code from a file into binary bytecode.
func (t tengoCompiler) Compile(filename string) (compiler.Bytecode, error) {
	src, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return nil, errors.Wrap(err, "read file")
	}

	fileSet := tengoSource.NewFileSet()
	srcFile := fileSet.AddFile(filename, -1, len(src))

	p := tengoParser.NewParser(srcFile, src, nil)
	file, err := p.ParseFile()
	if err != nil {
		return nil, errors.Wrap(err, "parsing")
	}

	c := tengo.NewCompiler(srcFile, nil, nil, nil, nil)
	if err := c.Compile(file); err != nil {
		return nil, errors.Wrap(err, "Tengo compiler")
	}

	return c.Bytecode(), nil
}

// TryCompile checks if source code from a file is compilable and returns the source code.
func (t tengoCompiler) TryCompile(filename string) ([]byte, error) {
	src, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return nil, errors.Wrap(err, "read file")
	}

	_, err = interop.NewEnvironment(src)
	if err != nil {
		return nil, err
	}

	return src, nil
}
