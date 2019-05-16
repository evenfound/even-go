package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/evenfound/even-go/core"
	"github.com/evenfound/even-go/pb"
	"github.com/evenfound/even-go/server/api"
	"github.com/golang/protobuf/ptypes/any"
	peer "gx/ipfs/QmZoWKhxUmZ2seW4BzX6fJkNR8hh9PsGModr7q171yq2SS/go-libp2p-peer"
	"gx/ipfs/QmcZfnkapfECQGcLZaf9B79NRg7cRa9EnZh4LSbkCzwNvY/go-cid"
)

type PeersHandler struct {
}

func (handler *PeersHandler) List(ctx context.Context, req *api.PeerEmptyRequest) (*api.PeerListResponse, error) {

	if core.Node != nil {

		peerStore := core.Node.IpfsNode.Peerstore.Peers()

		peers := []string{}

		for _, p := range peerStore {
			peers = append(peers, p.Pretty())
		}

		return &api.PeerListResponse{Peers: peers}, nil

	}

	return &api.PeerListResponse{}, errors.New("IPFS not is not s")
}

func (handler *PeersHandler) Message(hash, peerId string) error {

	p, err := peer.IDB58Decode(peerId)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), core.Node.OfflineMessageFailoverTimeout)

	defer cancel()

	message := pb.Message{
		MessageType: pb.Message_CHAT,
		Payload: &any.Any{
			Value: []byte(hash),
		},
	}

	err = core.Node.Service.SendMessage(ctx, p, &message)

	return nil
}

func (handler *PeersHandler) SendStore(ctx context.Context, req *api.PeerSendStoreRequest) (*api.PeerListResponse, error) {

	var hash = req.GetCid()

	id, err := cid.Decode(hash)

	if err != nil {
		return &api.PeerListResponse{}, err
	}

	var graph = []cid.Cid{*id}

	for _, p := range core.Node.IpfsNode.Peerstore.Peers() {

		pId := p.Pretty()

		err = core.Node.SendStore(pId, graph)

		fmt.Printf("Sending store to  %v \n", p.Pretty())

		if err != nil {

			fmt.Printf("Can not send store to peer with id %v \n", pId)

			fmt.Println(err)
		}

	}

	return &api.PeerListResponse{}, nil
}
