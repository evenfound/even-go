package server

// please add in alphabetical order
//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api crypto.proto
//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api eventransaction.proto
//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api files.proto
//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api multichain.proto
//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api peers.proto
//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api smartcontract.proto
//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api transaction.proto
//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api wallet.proto

import (
	"fmt"
	"log"
	"net"

	"github.com/evenfound/even-go/node/server/api"
	"github.com/evenfound/even-go/node/server/handlers"
	"google.golang.org/grpc"
)

// Run function listens gRPC server on specific port
func Run(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var (
		grpcServer = grpc.NewServer()

		// please add in alphabetical order
		cryptoHandler      = handlers.Crypto{}
		eventrHandler      = handlers.EvenTransaction{}
		filesHandler       = handlers.FilesHandler{}
		multichainHandler  = handlers.Multichain{}
		peerHandler        = handlers.PeersHandler{}
		smartHandler       = handlers.SmartContract{}
		transactionHandler = handlers.Transaction{}
		walletHandler      = handlers.Wallet{}
	)

	// please add in alphabetical order
	api.RegisterBalanceServer(grpcServer, &multichainHandler)
	api.RegisterCryptoServer(grpcServer, &cryptoHandler)
	api.RegisterEvenTransactionServer(grpcServer, &eventrHandler)
	api.RegisterSmartContractServer(grpcServer, &smartHandler)
	api.RegisterFileServiceServer(grpcServer, &filesHandler)
	api.RegisterPeersServer(grpcServer, &peerHandler)
	api.RegisterTransactionServer(grpcServer, &transactionHandler)
	api.RegisterWalletServer(grpcServer, &walletHandler)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	fmt.Println(fmt.Sprintf("Listening gRPC server on %v port", port))
}
