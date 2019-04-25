// Copyright (c) 2015-2016 The btcsuite developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package btcnet

import (
	"fmt"

	"github.com/evenfound/even-go/node/mbnd/common"
)

const (
	Token = "btc"
)

func initProvider() (common.Blockchain, error) {

	bc := common.Blockchain(&bchain{})

	return bc, nil
}

func init() {
	// Register the provider.
	provider := common.Provider{
		Token: Token,
		Init:  initProvider,
	}

	if err := common.RegisterProvider(provider); err != nil {
		panic(fmt.Sprintf("Failed to regiser blockchain provider '%s': %v", Token, err))
	}
}
