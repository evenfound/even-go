package hdwallet

import (
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/edunuzzi/go-bip44"
)

// Network represents type of the Even Network.
type Network int16

const (
	bitSize                 = 128 // 12 words in mnemonic seed
	bitcoinTestnet3 Network = iota + 1
	bitcoinMainnet
	evenTestnet = bitcoinTestnet3
	evenMainnet = bitcoinMainnet
)

func generateMnemonic() (string, error) {
	m, err := bip44.NewMnemonic(bitSize)
	if err != nil {
		return "", err
	}
	return m.Value, nil
}

func generateAddress(mnemonic, password string, index uint32) (string, error) {
	xKey, err := generateExtendedKey(mnemonic, password, index)
	if err != nil {
		return "", err
	}

	chainCfg, err := networkToChainConfig(evenTestnet)
	if err != nil {
		return "", err
	}

	address, err := xKey.Address(chainCfg)
	if err != nil {
		return "", err
	}

	return address.String(), nil
}

func generatePrivateKey(mnemonic, password string, index uint32) (string, error) {
	xKey, err := generateExtendedKey(mnemonic, password, index)
	if err != nil {
		return "", err
	}

	privKey, err := xKey.ECPrivKey()
	if err != nil {
		return "", err
	}

	return privateKeyToString(privKey), nil
}

func generatePublicKey(mnemonic, password string, index uint32) (string, error) {
	xKey, err := generateExtendedKey(mnemonic, password, index)
	if err != nil {
		return "", err
	}

	privKey, err := xKey.ECPrivKey()
	if err != nil {
		return "", err
	}

	pubKey := privKey.PubKey()

	return publicKeyToString(pubKey), nil
}

func generateExtendedKey(mnemonic, password string, index uint32) (*hdkeychain.ExtendedKey, error) {
	m := bip44.ParseMnemonic(mnemonic)
	seed, err := m.NewSeed(password)
	if err != nil {
		return nil, err
	}

	chainCfg, err := networkToChainConfig(evenTestnet)
	if err != nil {
		return nil, err
	}

	xKey, err := hdkeychain.NewMaster(seed, chainCfg)
	if err != nil {
		return nil, err
	}

	xKey, err = xKey.Child(index)
	if err != nil {
		return nil, err
	}

	return xKey, nil
}

func privateKeyToString(pk *btcec.PrivateKey) string {
	return hex.EncodeToString(pk.Serialize())
}

func publicKeyToString(puk *btcec.PublicKey) string {
	return hex.EncodeToString(puk.SerializeCompressed())
}

func networkToChainConfig(net Network) (*chaincfg.Params, error) {
	switch net {
	case bitcoinTestnet3:
		return &chaincfg.TestNet3Params, nil

	case bitcoinMainnet:
		return &chaincfg.MainNetParams, nil
	}

	return nil, errors.New("invalid network")
}
