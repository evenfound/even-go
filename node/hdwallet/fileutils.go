package hdwallet

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/alexmullins/zip"

	"github.com/evenfound/even-go/node/core"
	"github.com/evenfound/even-go/node/schema"
)

func writeEncrypted(filename, password string, data []byte) error {
	fzip, err := os.Create(filename)
	if err != nil {
		return err
	}
	zipw := zip.NewWriter(fzip)
	defer func() { must(zipw.Close()) }()
	w, err := zipw.Encrypt(walletContainer, password)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, bytes.NewReader(data))
	if err != nil {
		return err
	}
	err = zipw.Flush()
	if err != nil {
		return err
	}
	return nil
}

func readEncrypted(filename, password string) ([]byte, error) {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer func() { must(reader.Close()) }()

	var buf bytes.Buffer

	// Iterate through the files in the archive
	count := 0
	for _, f := range reader.File {
		if count > 0 {
			return nil, errors.New("invalid wallet file format (too many containers)")
		}
		if f.Name != walletContainer {
			return nil, errors.New("invalid wallet file format (unknown container)")
		}

		f.SetPassword(password)
		r, err := f.Open()
		if err != nil {
			if err == zip.ErrPassword {
				return nil, errors.New("invalid wallet password")
			}
			return nil, err
		}

		_, err = buf.ReadFrom(r)
		if err != nil {
			must(r.Close())
			return nil, err
		}

		must(r.Close())
		count++
	}
	if count < 1 {
		return nil, errors.New("invalid wallet file format (empty)")
	}

	return buf.Bytes(), nil
}

func absoluteFilename(name string) (string, error) {
	confDir, err := schema.OpenbazaarPathTransform(userHomeDir(), core.Node.TestnetEnable)
	if err != nil {
		return "", err
	}
	absoluteWalletDir := filepath.Join(confDir, walletDir)
	if _, err := os.Stat(absoluteWalletDir); err != nil {
		err = os.Mkdir(absoluteWalletDir, 0755)
		if err != nil {
			return "", err
		}
	}
	return filepath.Join(absoluteWalletDir, name+walletExt), nil
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

/** unused
func enumerateFilesInDir(dir string) ([]string, error) {
	files := make([]string, 0)
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, f := range entries {
		if !f.IsDir() {
			files = append(files, f.Name())
		}
	}
	return files, nil
}*/
