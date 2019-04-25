package util

import (
	"errors"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func DecodeAddress(address string, params *chaincfg.Params) (btcutil.Address, error) {
	return nil, errors.New("unknown address")
}
