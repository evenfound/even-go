package handlers

import (
	"github.com/evenfound/even-go/mbnd"
	"github.com/evenfound/even-go/server/api"
	"golang.org/x/net/context"
)

// Multichain is a handler.
type Multichain struct{}

// FetchBalance calls a handler.
func (x *Multichain) FetchBalance(ct context.Context, rq *api.BalancesRequest) (rs *api.BalancesResponse, err error) {

	rs = &api.BalancesResponse{
		Balances: make(map[string]float64, len(rq.Addrlist.Addresses)),
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
