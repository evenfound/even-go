package app

import (
	"github.com/evenfound/even-go/node/cmd/evec/config"
	"github.com/evenfound/even-go/node/cmd/evec/implementation"
	"github.com/evenfound/even-go/node/cmd/evec/tool"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

func clean() error {
	return cleanDir(config.WorkDir, config.CompiledExt)
}

func buildWorkDir() error {
	return filepath.Walk(config.WorkDir, func(path string, f os.FileInfo, _ error) error {
		if f != nil && !f.IsDir() {
			fn := f.Name()
			if config.LooksLikeSourceFile(fn) {
				if err := build(fn); err != nil {
					return tool.Wrap(err, "build")
				}
			}
		}
		return nil
	})
}

func buildFiles(filenames cli.Args) error {
	for _, f := range filenames {
		if err := build(f); err != nil {
			return tool.Wrap(err, "build")
		}
	}
	return nil
}

func build(filename string) error {
	basename := filepath.Base(filename)
	basename = strings.TrimSuffix(basename, filepath.Ext(basename))
	outputFilename := basename + config.CompiledExt
	return compile(filename, outputFilename)
}

func cleanDir(dir, suffix string) error {
	d, err := os.Open(filepath.Clean(dir))
	if err != nil {
		return tool.Wrap(err, "directory open")
	}
	defer func() { tool.Must(d.Close()) }()

	names, err := d.Readdirnames(0)
	if err != nil {
		return tool.Wrap(err, "Readdirnames")
	}

	for _, name := range names {
		if strings.HasSuffix(name, suffix) {
			err = os.Remove(filepath.Join(dir, name))
			if err != nil {
				return tool.Wrap(err, "file removal")
			}
		}
	}

	return nil
}

// compile selects a concrete compiler and performs the compilation.
func compile(inName, outName string) error {
	tool.TR("inName", inName)
	tool.TR("outName", outName)
	compiler := implementation.New(filepath.Ext(inName))
	bytecode, err := compiler.Compile(inName)
	if err != nil {
		return tool.Wrap(err, "compilation")
	}

	out, err := os.OpenFile(outName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return tool.Wrap(err, "file creation")
	}
	defer func() { tool.Must(out.Close()) }()

	err = bytecode.Encode(out)
	if err != nil {
		return tool.Wrap(err, "write bytecode to a file")
	}

	return nil
}
