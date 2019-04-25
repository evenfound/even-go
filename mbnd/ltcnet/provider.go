// Copyright (c) 2015-2016 The btcsuite developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package ltcnet

import (
	"fmt"

	"github.com/evenfound/even-go/mbnd/common"
)

const (
	NetToken = "ltc"
)

var useNetwork string

func initProvider(network string) (common.Blockchain, error) {

	useNetwork = network
	bc := common.Blockchain(&bchain{})

	return bc, nil
}

func init() {
	// Register the provider.
	provider := common.Provider{
		Token: NetToken,
		Init:  initProvider,
	}

	if err := common.RegisterProvider(provider); err != nil {
		panic(fmt.Sprintf("Failed to regiser blockchain provider '%s': %v", NetToken, err))
	}
}
