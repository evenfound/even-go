package rpc

import (
	"context"
	"log"
	"time"

	"github.com/evenfound/even-go/node/cmd/evenctl/config"
	"github.com/evenfound/even-go/node/cmd/evenctl/tool"
	pb "github.com/evenfound/even-go/node/server/api"
	"google.golang.org/grpc"
)

// Sign signs message with private key.
func Sign(message, privkey string) error {
	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return tool.Wrap(err, "RPC connect")
	}
	defer func() { tool.Must(conn.Close()) }()

	// Create a client
	scc := pb.NewCryptoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Make the call
	input := pb.SignInput{
		Message: message,
		Privkey: privkey,
	}
	r, err := scc.Sign(ctx, &input)
	if err != nil {
		return tool.Wrap(err, "Crypto.Sign")
	}

	if !r.Ok {
		return tool.NewError(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// Verify recovers account which was used to sign a message.
func Verify(message, signature, pubkey string) error {
	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return tool.Wrap(err, "RPC connect")
	}
	defer func() { tool.Must(conn.Close()) }()

	// Create a client
	scc := pb.NewCryptoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Make the call
	input := pb.VerifyInput{
		Message:   message,
		Signature: signature,
		Pubkey:    pubkey,
	}
	r, err := scc.Verify(ctx, &input)
	if err != nil {
		return tool.Wrap(err, "Crypto.Verify")
	}

	if !r.Ok {
		return tool.NewError(r.Result)
	}
	log.Println(r.Result)

	return nil
}
