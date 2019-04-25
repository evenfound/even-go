// Copyright (c) 2018-2019 The Even Foundation developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package mbnd

import (
	_ "github.com/evenfound/even-go/node/mbnd/btcnet"

	"github.com/evenfound/even-go/node/mbnd/common"
	"github.com/evenfound/even-go/node/server/api"
	"github.com/go-errors/errors"
)

type (
	// Define an instance struct.
	FetchBalance struct {
		Request  *api.AddressesRequest
		Response *api.BalancesResponse
	}
)

// Execute is the main entry point.
func (x *FetchBalance) Execute() error {

	if x.Request.Addresses == nil {
		return errors.New("address list mandatory parameter is empty")
	}

	chain, err := common.Init("btc")
	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Infof("Blockchain network: %v", chain.String())

	err = chain.Open()
	if err != nil {
		logger.Error(err)
		return err
	}

	defer func() {
		chain.Close()
	}()

	for _, addr := range x.Request.Addresses {

		res, err := chain.Balance(addr.Address)
		if err != nil {
			logger.Error(err)
			return err
		}

		x.Response.Balances[addr.Address] = res.Value
	}

	logger.Infof("Response: %v", x.Response)

	return nil
}
