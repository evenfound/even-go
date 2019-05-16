package transaction

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"os"
	"path/filepath"
)

func load(filename string) ([]byte, error) {
	stream, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return nil, err
	}

	header := []byte(zjsonFile.String())
	if bytes.HasPrefix(stream, header) {
		stream = stream[len(header):]
		stream, err = unpack(stream)
		if err != nil {
			return nil, err
		}
		stream = append([]byte(jsonFile.String()), stream...)
	}

	return stream, nil
}

func saveLocal(filename string, data []byte) (string, error) {
	filename = "/tmp/" + filename
	out, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}
	defer func() { must(out.Close()) }()

	if _, err := out.Write(data); err != nil {
		return "", err
	}

	return filename, nil
}

func compress(input []byte) ([]byte, error) {
	var result bytes.Buffer

	zipper := zlib.NewWriter(&result)
	if _, err := zipper.Write(input); err != nil {
		return nil, err
	}
	if err := zipper.Flush(); err != nil {
		return nil, err
	}
	if err := zipper.Close(); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

func unpack(data []byte) ([]byte, error) {
	rdata := bytes.NewReader(data)
	r, err := zlib.NewReader(rdata)
	if err != nil {
		return nil, err
	}

	src, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	must(r.Close())

	return src, nil
}
