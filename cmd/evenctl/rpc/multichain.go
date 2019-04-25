package rpc

import (
	"context"
	"log"
	"time"

	"github.com/evenfound/even-go/cmd/evenctl/config"
	"github.com/evenfound/even-go/cmd/evenctl/tool"
	"github.com/evenfound/even-go/server/api"
	"google.golang.org/grpc"
)

// Call performs a gRPC call.
func AddrBalance(token string, network string, addrList []string) error {

	// Set up a connection to the server
	conn, err := grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	if err != nil {
		return tool.Wrap(err, "RPC connect")
	}

	defer func() {
		tool.Must(conn.Close())
	}()

	// Create a context
	timeout := 60 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	// Make a request input data
	input := api.BalancesRequest{
		Token:   token,
		Network: api.BalancesRequest_Network(api.BalancesRequest_Network_value[network]),
		Addrlist: &api.Addrlist{
			Addresses: make([]*api.Addrlist_Address, len(addrList)),
		},
	}

	for idx, addr := range addrList {
		input.Addrlist.Addresses[idx] = &api.Addrlist_Address{
			Address: addr,
		}
	}

	// Create a client
	client := api.NewBalanceClient(conn)
	response, err := client.FetchBalance(ctx, &input)
	if err != nil {
		return tool.Wrap(err, "Addr.FetchBalance")
	}

	if len(response.Balances) > 0 {
		log.Printf("Call succeeded with: %v", response.GetBalances())
	} else {
		log.Print("Call failed with empty result")
	}

	return nil
}
