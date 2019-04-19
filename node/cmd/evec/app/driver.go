package app

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/evenfound/even-go/node/cmd/evec/config"
	"github.com/evenfound/even-go/node/cmd/evec/implementation"
	"github.com/evenfound/even-go/node/cmd/evec/tool"

	shell "github.com/ipfs/go-ipfs-api"
)

const (
	header = "EVEN"
	ipfs   = "ipfs"
)

func clean() error {
	return cleanDir(config.WorkDir, config.CompiledExt)
}

func buildFiles(files []string, of string) error {
	var err error
	for _, f := range files {
		isIPFS := of == ipfs
		if isIPFS {
			of = generateOutputFilename(ipfs)
		}
		if of == "" {
			of = generateOutputFilename(f)
		}
		if err = compile(f, of); err != nil {
			return tool.Wrap(err, "build")
		}
		log.Printf("Built %s\n", of)
		if isIPFS {
			if of, err = storeToIPFS(of); err != nil {
				return tool.Wrap(err, "store to IPFS")
			}
			log.Printf("Stored %s\n", of)
		}
	}
	return nil
}

func generateOutputFilename(filename string) string {
	if filename == ipfs {
		return generateTempFilename(filename)
	}
	basename := filepath.Base(filename)
	basename = strings.TrimSuffix(basename, filepath.Ext(basename))
	return basename + config.CompiledExt
}

func generateTempFilename(prefix string) string {
	tmpfile, err := ioutil.TempFile("", prefix+"*"+config.CompiledExt)
	tool.Must(err)
	tool.Must(tmpfile.Close())
	return tmpfile.Name()
}

func storeToIPFS(filename string) (string, error) {
	sh := shell.NewLocalShell()
	if sh == nil {
		return filename, tool.NewError("IPFS daemon is not running")
	}
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return filename, err
	}
	cid, err := sh.Add(file)
	if err != nil {
		return filename, err
	}
	tool.Must(os.Remove(filename))
	return "/" + ipfs + "/" + cid, nil
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
		return err
	}

	binary, err := compress(src)
	if err != nil {
		return err
	}

	binary, err = encrypt(binary)
	if err != nil {
		return err
	}

	if err := saveToFile(outName, binary); err != nil {
		return err
	}

	return nil
}

func compress(input []byte) ([]byte, error) {
	var result bytes.Buffer

	zipper := gzip.NewWriter(&result)
	if _, err := zipper.Write(input); err != nil {
		return nil, tool.Wrap(err, "compress")
	}
	if err := zipper.Flush(); err != nil {
		return nil, tool.Wrap(err, "flush")
	}
	if err := zipper.Close(); err != nil {
		return nil, tool.Wrap(err, "close")
	}

	return result.Bytes(), nil
}

func encrypt(stream []byte) ([]byte, error) {
	return stream, nil
}

/*
func encrypt(stream []byte) ([]byte, error) {
	const keyLengthBytes = 10
	//fmt.Println("input:", stream)
	key := make([]byte, keyLengthBytes)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	//fmt.Println("key:", key)
	c, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	c.XORKeyStream(stream, stream)
	//fmt.Println("output:", key, stream)
	return append(key, stream...), nil
}*/

func saveToFile(filename string, data []byte) error {
	out, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return tool.Wrap(err, "create file")
	}
	defer func() { tool.Must(out.Close()) }()

	if _, err := out.Write([]byte(header)); err != nil {
		return tool.Wrap(err, "write to file")
	}
	if _, err := out.Write(data); err != nil {
		return tool.Wrap(err, "write to file")
	}

	return nil
}
