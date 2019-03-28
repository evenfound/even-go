package handlers

import (
	"fmt"
	"io/ioutil"

	"github.com/evenfound/even-go/node/evm"

	"github.com/evenfound/even-go/node/server/api"
	"golang.org/x/net/context"
)

// SmartContract is a smart contract handler.
type SmartContract struct{}

// Call creates a VM instance and makes a call of a smart contract.
func (sc *SmartContract) Call(ctx context.Context, in *api.ContractUri) (*api.ContractResult, error) {
	fmt.Println("SmartContract.Call " + in.Uri)
	vm := evm.New()
	bytecode, err := ioutil.ReadFile(in.Uri)
	if err != nil {
		return nil, err
	}

	err = vm.Run(bytecode)
	if err != nil {
		return nil, err
	}

	res := &api.ContractResult{Result: "OK"}

	return res, nil
}
