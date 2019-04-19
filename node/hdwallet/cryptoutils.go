package hdwallet

import (
	"crypto/ecdsa"
	"errors"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/edunuzzi/go-bip44"
)

type Network int16

const (
	bitSize                 = 128 // 12 words in mnemonic seed
	BitcoinTestnet3 Network = iota + 1
	BitcoinMainnet
	EvenTestnet = BitcoinTestnet3
	EvenMainnet = BitcoinMainnet
)

func generateMnemonic() string {
	m, _ := bip44.NewMnemonic(bitSize)
	return m.Value
}

func generateAddress(mnemonic, password string, index uint32) (string, error) {
	xKey, err := generateExtendedKey(mnemonic, password, index)
	if err != nil {
		return "", err
	}

	chainCfg, err := networkToChainConfig(EvenTestnet)
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
	privKeyECDSA := privKey.ToECDSA()

	pubKey := privKeyECDSA.Public()
	puk, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("failed to get public key")
	}

	return publicKeyToString(puk), nil
}

func generateExtendedKey(mnemonic, password string, index uint32) (*hdkeychain.ExtendedKey, error) {
	m := bip44.ParseMnemonic(mnemonic)
	seed, err := m.NewSeed(password)
	if err != nil {
		return nil, err
	}

	chainCfg, err := networkToChainConfig(EvenTestnet)
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
	return pk.D.Text(16)
}

func publicKeyToString(pk *ecdsa.PublicKey) string {
	return pk.X.Text(16) + pk.Y.Text(16)
}

func networkToChainConfig(net Network) (*chaincfg.Params, error) {
	switch net {
	case BitcoinTestnet3:
		return &chaincfg.TestNet3Params, nil

	case BitcoinMainnet:
		return &chaincfg.MainNetParams, nil
	}

	return nil, errors.New("invalid network")
}
