package app

import (
	"bytes"
	"compress/gzip"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/evenfound/even-go/node/cmd/evec/config"
	"github.com/evenfound/even-go/node/cmd/evec/implementation"
	"github.com/evenfound/even-go/node/cmd/evec/tool"
	"github.com/urfave/cli"
)

const (
	header = "EVEN"
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
	err := compile(filename, outputFilename)
	if err == nil {
		log.Printf("Built %s\n", outputFilename)
	}
	return err
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
			if err := os.Remove(filepath.Join(dir, name)); err != nil {
				return tool.Wrap(err, "file removal")
			}
		}
	}

	return nil
}

// compile selects a concrete compiler and performs the compilation.
func compile(inName, outName string) error {
	compiler := implementation.New(filepath.Ext(inName))
	if compiler == nil {
		return tool.NewError("unknown format of file " + inName)
	}

	src, err := compiler.TryCompile(inName)
	if err != nil {
		return err // no need to wrap
	}

	var binary bytes.Buffer
	zipper := gzip.NewWriter(&binary)

	if _, err := zipper.Write(src); err != nil {
		return tool.Wrap(err, "compress")
	}
	if err := zipper.Flush(); err != nil {
		return tool.Wrap(err, "flush")
	}
	if err := zipper.Close(); err != nil {
		return tool.Wrap(err, "close")
	}

	out, err := os.OpenFile(outName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return tool.Wrap(err, "create file")
	}
	defer func() { tool.Must(out.Close()) }()

	if _, err := out.Write([]byte(header)); err != nil {
		return tool.Wrap(err, "write to file")
	}
	if _, err := out.Write(binary.Bytes()); err != nil {
		return tool.Wrap(err, "write to file")
	}

	return nil
}
