package internal

import (
	"github.com/evenfound/even-go/node/cmd/evec/compiler"
	"github.com/evenfound/even-go/node/cmd/evec/tool"
	"io/ioutil"
	"path/filepath"

	tengo "github.com/d5/tengo/compiler"
	tengoParser "github.com/d5/tengo/compiler/parser"
	tengoSource "github.com/d5/tengo/compiler/source"
)

type tengoCompiler struct {
	//
}

// Compiler translates a source code from a file into binary bytecode.
func (t tengoCompiler) Compile(filename string) (compiler.Bytecode, error) {
	src, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return nil, tool.Wrap(err, "read file")
	}

	fileSet := tengoSource.NewFileSet()
	srcFile := fileSet.AddFile(filename, -1, len(src))

	p := tengoParser.NewParser(srcFile, src, nil)
	file, err := p.ParseFile()
	if err != nil {
		return nil, tool.Wrap(err, "parsing")
	}

	c := tengo.NewCompiler(srcFile, nil, nil, nil, nil)
	if err := c.Compile(file); err != nil {
		return nil, tool.Wrap(err, "tengo compiler")
	}

	return c.Bytecode(), nil
}
