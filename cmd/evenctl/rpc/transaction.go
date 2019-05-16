package rpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/evenfound/even-go/cmd/evenctl/config"
	pb "github.com/evenfound/even-go/server/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func formatCode(format string) (int32, error) {
	switch format {
	case "json":
		return 1, nil
	case "zjson":
		return 2, nil
	case "ubjson":
		return 3, nil
	case "gob":
		return 4, nil
	}
	msg := fmt.Sprintf("'%s' unknown file format (expected json | zjson | ubjson | gob)", format)
	return 0, errors.New(msg)
}

// CreateTransaction creates new transaction.
func CreateTransaction(format string) error {
	fc, err := formatCode(format)
	if err != nil {
		return err
	}

	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "RPC connect")
	}
	defer func() { must(conn.Close()) }()

	// Create a client
	scc := pb.NewEvenTransactionClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Make the call
	input := pb.EvenTransactionCreateInput{
		Format: fc,
		Tag:    "TEST",
	}
	r, err := scc.Create(ctx, &input)
	if err != nil {
		return errors.Wrap(err, "EvenTransaction.Create")
	}

	if !r.Ok {
		return errors.New(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// ShowTransaction reads and shows transaction.
func ShowTransaction(filename string) error {
	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "RPC connect")
	}
	defer func() { must(conn.Close()) }()

	// Create a client
	scc := pb.NewEvenTransactionClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Make the call
	input := pb.EvenTransactionInput{
		Filename: filename,
	}
	r, err := scc.Show(ctx, &input)
	if err != nil {
		return errors.Wrap(err, "EvenTransaction.Show")
	}

	if !r.Ok {
		return errors.New(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// AnalyzeTransaction reads and analyzes transaction.
func AnalyzeTransaction(filename string) error {
	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "RPC connect")
	}
	defer func() { must(conn.Close()) }()

	// Create a client
	scc := pb.NewEvenTransactionClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Make the call
	input := pb.EvenTransactionInput{
		Filename: filename,
	}
	r, err := scc.Analyze(ctx, &input)
	if err != nil {
		return errors.Wrap(err, "EvenTransaction.Analyze")
	}

	if !r.Ok {
		return errors.New(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// VerifyTransaction reads and validates transaction.
func VerifyTransaction(filename string) error {
	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "RPC connect")
	}
	defer func() { must(conn.Close()) }()

	// Create a client
	scc := pb.NewEvenTransactionClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Make the call
	input := pb.EvenTransactionInput{
		Filename: filename,
	}
	r, err := scc.Verify(ctx, &input)
	if err != nil {
		return errors.Wrap(err, "EvenTransaction.Verify")
	}

	if !r.Ok {
		return errors.New(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// WalletAccountTxNewReg creates new initial transaction.
func WalletAccountTxNewReg(name, password, account string) error {
	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "RPC connect")
	}
	defer func() { must(conn.Close()) }()

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
	r, err := scc.WalletAccountTxNewReg(ctx, &input)
	if err != nil {
		return errors.Wrap(err, "Wallet.WalletAccountTxNewReg")
	}

	if !r.Ok {
		return errors.New(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// WalletAccountTxContract creates contract-deploy transaction.
func WalletAccountTxContract(name, password, account, contract string) error {
	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "RPC connect")
	}
	defer func() { must(conn.Close()) }()

	// Create a client
	scc := pb.NewWalletClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Make the call
	input := pb.WalletContractInput{
		Name:     name,
		Password: password,
		Account:  account,
		Contract: contract,
	}
	r, err := scc.WalletAccountTxContract(ctx, &input)
	if err != nil {
		return errors.Wrap(err, "Wallet.WalletAccountTxContract")
	}

	if !r.Ok {
		return errors.New(r.Result)
	}
	log.Println(r.Result)

	return nil
}

// WalletAccountTxContractInvoke creates contract-invoking transaction.
func WalletAccountTxContractInvoke(name, password, account, contract, function string) error {
	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "RPC connect")
	}
	defer func() { must(conn.Close()) }()

	// Create a client
	scc := pb.NewWalletClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Make the call
	input := pb.WalletContractInput{
		Name:     name,
		Password: password,
		Account:  account,
		Contract: contract,
		Function: function,
	}
	r, err := scc.WalletAccountTxInvoke(ctx, &input)
	if err != nil {
		return errors.Wrap(err, "Wallet.WalletAccountTxInvoke")
	}

	if !r.Ok {
		return errors.New(r.Result)
	}
	log.Println(r.Result)

	return nil
}
