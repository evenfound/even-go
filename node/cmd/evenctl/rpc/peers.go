package rpc

import (
	"context"
	"fmt"
	"github.com/evenfound/even-go/node/cmd/evenctl/config"
	"github.com/evenfound/even-go/node/server/api"
	"google.golang.org/grpc"
)

type PeerCMD struct {
	ctx  context.Context
	stub api.PeersClient
}

func NewPeerCMD(ctx context.Context) (*PeerCMD, error) {

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
