package ipfs

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/ipfs/go-ipfs-config"
	libp2p "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"

	//"github.com/ipfs/go-ipfs/repo/fsrepo"
)

func IdentityFromKey(privkey []byte) (config.Identity, error) {

	ident := config.Identity{}
	sk, err := libp2p.UnmarshalPrivateKey(privkey)
	if err != nil {
		return ident, err
	}
	skbytes, err := sk.Bytes()
	if err != nil {
		return ident, err
	}
	ident.PrivKey = base64.StdEncoding.EncodeToString(skbytes)

	id, err := peer.IDFromPublicKey(sk.GetPublic())
	if err != nil {
		return ident, err
	}
	ident.PeerID = id.Pretty()
	return ident, nil
}

func IdentityKeyFromSeed(seed []byte, bits int) ([]byte, error) {
	hmac := hmac.New(sha256.New, []byte("EvenNetwork seed"))
	hmac.Write(seed)
	reader := bytes.NewReader(hmac.Sum(nil))
	sk, _, err := libp2p.GenerateKeyPairWithReader(libp2p.Ed25519, bits, reader)
	if err != nil {
		return nil, err
	}
	encodedKey, err := sk.Bytes()
	if err != nil {
		return nil, err
	}
	return encodedKey, nil
}
