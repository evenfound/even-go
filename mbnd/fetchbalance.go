// Copyright (c) 2018-2019 The Even Foundation developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package mbnd

import (
	"fmt"

	"github.com/evenfound/even-go/mbnd/common"
	"github.com/evenfound/even-go/server/api"
	"github.com/go-errors/errors"

	// add here a new supported blockchain network in alphabetical order
	_ "github.com/evenfound/even-go/mbnd/btcnet"
	_ "github.com/evenfound/even-go/mbnd/ltcnet"
)

type (
	// Define an instance struct.
	FetchBalance struct {
		Request  *api.BalancesRequest
		Response *api.BalancesResponse
	}
)

// Execute is the main entry point.
func (x *FetchBalance) Execute() error {

	if x.Request.Token == "" || common.SupportedProvider(x.Request.Token) == false {
		return errors.New(fmt.Sprintf("only supported providers values: %v", common.SupportingProviders()))
	}

	if len(x.Request.Addrlist.Addresses) == 0 {
		return errors.New("address list mandatory parameter is empty")
	}

	chain, err := common.Init(x.Request.Token, x.Request.Network.String())
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

	for _, addr := range x.Request.Addrlist.Addresses {

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
