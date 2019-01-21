package main

import (
	"fmt"
	"hdgen"
	"net"
	"net/http"
	"net/rpc"
)

type RPCPool struct{}

type RPCResponse struct {
	Status bool
	Data   []string
}

var pool = new(RPCPool)

func RPCConnect() {
	err := rpc.Register(pool)

	if err != nil {
		panic(err)
	}

	rpc.HandleHTTP()

	listener, e := net.Listen("tcp", ":"+config.rpcPort)
	if e != nil {
		panic(e)
	}
	err = http.Serve(listener, nil)
	if err != nil {
		panic(err)
	}
}

func (rpc *RPCPool) Create(phrase hdgen.WalletGenerator, response *RPCResponse) error {
	fmt.Println("RPC:Create")
	return nil
}

func (rpc *RPCPool) Generate(data hdgen.AddressGenerator, response *RPCResponse) error {
	fmt.Println("RPC:Generate")
	return nil
}

func (rpc RPCPool) List(_ string, response *RPCResponse) error {
	fmt.Println("RPC:List")
	return nil
}

func (rpc *RPCPool) Delete(hash string, response *RPCResponse) error {
	fmt.Println("RPC:Delete")
	return nil
}
