package server

//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api transaction.proto
//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api files.proto
//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api peers.proto
//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api smartcontract.proto

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
		transactionHandler = handlers.Transaction{}
		smartHandler       = handlers.SmartContract{}
		filesHandler       = handlers.FilesHandler{}
		peerHandler        = handlers.PeersHandler{}
		grpcServer         = grpc.NewServer()
	)

	api.RegisterTransactionServer(grpcServer, &transactionHandler)
	api.RegisterSmartContractServer(grpcServer, &smartHandler)
	api.RegisterFileServiceServer(grpcServer, &filesHandler)
	api.RegisterPeersServer(grpcServer, &peerHandler)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	fmt.Println(fmt.Sprintf("Listening gRPC server on %v port", port))
}
