package app

//go:generate protoc --proto_path=../../server/proto --go_out=plugins=grpc:rpc/api smartcontract.proto

import (
	"context"
	"log"
	"time"

	"github.com/evenfound/even-go/node/cmd/evenctl/tool"

	pb "github.com/evenfound/even-go/node/cmd/evenctl/rpc/api"

	"google.golang.org/grpc"
)

const (
	rpcAddress = "localhost:8090"
)

func call(filename string) error {
	// Set up a connection to the server
	conn, err := grpc.Dial(rpcAddress, grpc.WithInsecure())
	if err != nil {
		return tool.Wrap(err, "connect")
	}
	defer func() { tool.Must(conn.Close()) }()

	// Create a client and make a call
	scc := pb.NewSmartContractClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := scc.Call(ctx, &pb.ContractUri{Uri: filename})
	if err != nil {
		return tool.Wrap(err, "call")
	}

	log.Printf("Result: %s", r.Result)
	return nil
}
