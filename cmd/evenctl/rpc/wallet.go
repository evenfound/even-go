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

// GenerateWallet performs a gRPC call implementing generation of new wallet.
func GenerateWallet(name, password string) error {
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
	input := pb.WalletInput{
		Name:     name,
		Password: password,
	}
	r, err := scc.GenerateWallet(ctx, &input)
	if err != nil {
		return tool.Wrap(err, "Wallet.GenerateWallet")
	}

	if !r.Ok {
		return tool.NewError(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// CreateWallet performs a gRPC call implementing creation of wallet.
func CreateWallet(name, mnemonic, password string) error {
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
	input := pb.CreateWalletInput{
		Name:     name,
		Mnemonic: mnemonic,
		Password: password,
	}
	r, err := scc.CreateWallet(ctx, &input)
	if err != nil {
		return tool.Wrap(err, "Wallet.CreateWallet")
	}

	if !r.Ok {
		return tool.NewError(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// UnlockWallet performs a gRPC call implementing unlock a wallet.
func UnlockWallet(name, password string) error {
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
	input := pb.WalletInput{
		Name:     name,
		Password: password,
	}
	r, err := scc.UnlockWallet(ctx, &input)
	if err != nil {
		return tool.Wrap(err, "Wallet.UnlockWallet")
	}

	if !r.Ok {
		return tool.NewError(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// WalletNextAccount performs a gRPC call implementing generation of next account.
func WalletNextAccount(name, password string) error {
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
	input := pb.WalletInput{
		Name:     name,
		Password: password,
	}
	r, err := scc.WalletNextAccount(ctx, &input)
	if err != nil {
		return tool.Wrap(err, "Wallet.WalletNextAccount")
	}

	if !r.Ok {
		return tool.NewError(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// GetWalletInfo performs a gRPC call implementing wallet information retrieval.
func GetWalletInfo(name, password string) error {
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
	input := pb.WalletInput{
		Name:     name,
		Password: password,
	}
	r, err := scc.GetWalletInfo(ctx, &input)
	if err != nil {
		return tool.Wrap(err, "Wallet.GetWalletInfo")
	}

	if !r.Ok {
		return tool.NewError(r.Result)
	}
	log.Println(r.Result)

	return nil
}
