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

// WalletAccountDumpPrivateKey performs a gRPC call implementing dump of the account private key.
func WalletAccountDumpPrivateKey(name, password, account string) error {
	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return tool.Wrap(err, "RPC connect")
	}
	defer func() { tool.Must(conn.Close()) }()

	// Create a client
	scc := pb.NewWalletClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Make the call
	input := pb.WalletAccountInput{
		Name:     name,
		Password: password,
		Account:  account,
	}
	r, err := scc.WalletAccountDumpPrivateKey(ctx, &input)
	if err != nil {
		return tool.Wrap(err, "Wallet.WalletAccountDumpPrivateKey")
	}

	if !r.Ok {
		return tool.NewError(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// WalletAccountDumpPublicKey performs a gRPC call implementing dump of the account public key.
func WalletAccountDumpPublicKey(name, password, account string) error {
	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return tool.Wrap(err, "RPC connect")
	}
	defer func() { tool.Must(conn.Close()) }()

	// Create a client
	scc := pb.NewWalletClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Make the call
	input := pb.WalletAccountInput{
		Name:     name,
		Password: password,
		Account:  account,
	}
	r, err := scc.WalletAccountDumpPublicKey(ctx, &input)
	if err != nil {
		return tool.Wrap(err, "Wallet.WalletAccountDumpPublicKey")
	}

	if !r.Ok {
		return tool.NewError(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// WalletAccountShowBalance performs a gRPC call implementing view of the account balance.
func WalletAccountShowBalance(name, password, account string) error {
	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return tool.Wrap(err, "RPC connect")
	}
	defer func() { tool.Must(conn.Close()) }()

	// Create a client
	scc := pb.NewWalletClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Make the call
	input := pb.WalletAccountInput{
		Name:     name,
		Password: password,
		Account:  account,
	}
	r, err := scc.WalletAccountShowBalance(ctx, &input)
	if err != nil {
		return tool.Wrap(err, "Wallet.WalletAccountShowBalance")
	}

	if !r.Ok {
		return tool.NewError(r.Result)
	}
	log.Println(r.Result)

	return nil
}
