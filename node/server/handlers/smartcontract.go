package handlers

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/evenfound/even-go/node/evm"

	"github.com/evenfound/even-go/node/server/api"
	"golang.org/x/net/context"
)

const filePrefix = "file://"

// SmartContract is a smart contract handler.
type SmartContract struct{}

// Call creates a VM instance and makes a call of a smart contract.
func (sc *SmartContract) Call(ctx context.Context, in *api.SmartContractInput) (*api.SmartContractResult, error) {
	filename := in.Uri
	if strings.HasPrefix(filename, filePrefix) {
		filename = filename[len(filePrefix):]
	}

	bytecode, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return nil, err
	}

	vm := evm.New()
	err = vm.Run(bytecode)
	if err != nil {
		return nil, err
	}

	res := &api.SmartContractResult{Result: "OK"}

	return res, nil
}
