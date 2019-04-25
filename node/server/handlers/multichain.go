package handlers

import (
	"github.com/evenfound/even-go/node/mbnd"
	"github.com/evenfound/even-go/node/server/api"
	"golang.org/x/net/context"
)

// Multichain is a handler.
type Multichain struct{}

// FetchBalance calls a handler.
func (x *Multichain) FetchBalance(ct context.Context, rq *api.AddressesRequest) (rs *api.BalancesResponse, err error) {

	rs = &api.BalancesResponse{
		Balances: make(map[string]float64, len(rq.Addresses)),
	}

	fetcher := &mbnd.FetchBalance{
		Request:  rq,
		Response: rs,
	}

	err = fetcher.Execute()
	if err != nil {
		return nil, err
	}

	return rs, nil
}
