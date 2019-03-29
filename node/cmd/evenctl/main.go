package main

//go:generate protoc --proto_path=../../server/proto --go_out=plugins=grpc:rpc/api smartcontract.proto

import (
	"log"

	"github.com/evenfound/even-go/node/cmd/evenctl/app"
	"github.com/evenfound/even-go/node/cmd/evenctl/config"

	"github.com/ztrue/tracerr"
)

func main() {
	app.Init()
	defer app.Close()
	err := app.Run()
	if err != nil {
		if config.Debug {
			tracerr.PrintSourceColor(err)
		}
		log.Fatal("Fatal error: ", err)
	}
}
