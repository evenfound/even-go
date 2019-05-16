package evm

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/evenfound/even-go/node/evm/interop"
)

// Ensure tengoVM satisfies evm.Interface.
var _ Interface = tengoVM{}

// tengoVM represents the Tengo VM.
type tengoVM struct {
}

// newTengoVM creates new instance of the tengoVM.
func newTengoVM() Interface {
	return tengoVM{}
}

// Run implements corresponding method of the EVM interface.
func (tengoVM) Run(bc Bytecode, entryFunc string) (string, error) {
	if entryFunc == "" {
		entryFunc = DefaultEntryFunction
	}
	entryFunc = "\n" + entryFunc + "()\n"

	bc = decypher(bc)

	src, err := unpack(bc)
	if err != nil {
		return "", err
	}

	src = instrument(src, entryFunc)

	env, err := interop.NewEnvironment(src)
	if err != nil {
		err = simplifyError(err)
		return "", err
	}

	if err := run(env); err != nil {
		return "", err
	}

	return env.Get("result").String(), nil
}

func decypher(data []byte) []byte {
	const header = "EVENtimestampinnanosecs"
	if len(data) > len(header) {
		return data[len(header):]
	}
	return data
}

func unpack(data []byte) ([]byte, error) {
	rdata := bytes.NewReader(data)
	r, err := gzip.NewReader(rdata)
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

func instrument(text []byte, entryFunc string) []byte {
	// Add call of the entry function
	text = append(text, []byte(entryFunc)...)
	//tr("Instrumented:\n", string(text))
	return text
}

// run runs the compiled script and handles panics.
func run(env *interop.Environment) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()
	err = env.Run()
	return
}

func simplifyError(err error) error {
	s := strings.ReplaceAll(err.Error(), "Compile Error:", "")
	s = strings.Trim(s, " ")
	return errors.New(s)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
