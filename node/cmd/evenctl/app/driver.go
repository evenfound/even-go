package app

//go:generate protoc --go_out=plugins=grpc:. smartcontract/smartcontract.proto

import (
	"context"
	"github.com/evenfound/even-go/node/cmd/evenctl/tool"
	"log"
	"time"

	pb "github.com/evenfound/even-go/node/cmd/evenctl/app/smartcontract"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func call() error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return tool.Wrap(err, "connect")
	}
	defer func() { tool.Must(conn.Close()) }()

	// Create a client and call.
	scc := pb.NewSmartContractClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := scc.Call(ctx, &pb.ContractUri{Uri: "xxxxxxxxxxxxxx"})
	if err != nil {
		return tool.Wrap(err, "call")
	}

	log.Printf("Result: %s", r.Result)
	return nil
}
