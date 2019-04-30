package rpc

import (
	"context"
	"fmt"

	"github.com/evenfound/even-go/node/cmd/evenctl/config"
	"github.com/evenfound/even-go/node/server/api"
	"google.golang.org/grpc"
)

// PeerCMD ...
type PeerCMD struct {
	ctx  context.Context
	stub api.PeersClient
}

// NewPeerCMD ...
func NewPeerCMD() (*PeerCMD, error) {

	ctx := context.Background()

	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	stub := api.NewPeersClient(conn)

	return &PeerCMD{
		stub: stub,
		ctx:  ctx,
	}, nil

}

// List ...
func (cmd *PeerCMD) List() error {

	resp, err := cmd.stub.List(cmd.ctx, &api.PeerEmptyRequest{})

	if err != nil {
		return err
	}

	for _, id := range resp.Peers {
		fmt.Println("Peer : " + id)
	}

	return nil
}

// SendStore ...
func (cmd *PeerCMD) SendStore(hash string) error {

	fmt.Printf("Sending hash %v to peers \n", hash)

	_, err := cmd.stub.SendStore(cmd.ctx, &api.PeerSendStoreRequest{
		Cid: hash,
	})

	if err != nil {
		fmt.Printf("Error : %v \n", err)
		return err
	}

	return nil
}
