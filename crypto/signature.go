package crypto

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec"
)

// Sign signs a message and returns the signature.
func Sign(message, privkey string) (string, error) {
	digest := Hash([]byte(message))

	privKeySer, err := hex.DecodeString(privkey)
	if err != nil {
		return "", err
	}

	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeySer)
	signature, err := privKey.Sign(digest)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(signature.Serialize()), nil
}

// Verify checks if the signature has been generated from paired key.
func Verify(message, signatureStr, pubkey string) (bool, error) {
	pubKeySer, err := hex.DecodeString(pubkey)
	if err != nil {
		return false, err
	}

	pubKey, err := btcec.ParsePubKey(pubKeySer, btcec.S256())
	if err != nil {
		return false, err
	}

	digest := Hash([]byte(message))
	signatureSer, err := hex.DecodeString(signatureStr)
	if err != nil {
		return false, err
	}

	signature, err := btcec.ParseSignature(signatureSer, btcec.S256())
	if err != nil {
		return false, err
	}

	valid := signature.Verify(digest, pubKey)
	return valid, nil
}
