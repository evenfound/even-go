package rpc

import (
	"context"
	"log"
	"time"

	"github.com/evenfound/even-go/cmd/evenctl/config"
	"github.com/evenfound/even-go/cmd/evenctl/tool"
	pb "github.com/evenfound/even-go/server/api"
	"google.golang.org/grpc"
)

// Call performs a gRPC call implementing a smart contract call.
func Call(filename, entryFunc string) error {
	if !isCorrectFilename(filename) {
		return tool.NewError("filename '" + filename +
			"' is incorrect. It should correspond to the file://<host>/<path> or /ipfs/<hash> scheme")
	}
	if !isCorrectFunction(entryFunc) {
		return tool.NewError("name '" + entryFunc +
			"' is incorrect. It should start with a letter")
	}

	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return tool.Wrap(err, "RPC connect")
	}
	defer func() { tool.Must(conn.Close()) }()

	// Create a client
	scc := pb.NewSmartContractClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Make the call
	input := pb.SmartContractInput{
		Uri:       filename,
		EntryFunc: entryFunc,
	}
	r, err := scc.Call(ctx, &input)
	if err != nil {
		return tool.Wrap(err, "SmartContract.Call")
	}

	if r.Ok {
		log.Printf("Call succeeded with '%s'", r.Result)
	} else {
		log.Printf("Call failed with '%s'", r.Result)
	}

	return nil
}
