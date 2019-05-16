package transaction

import "github.com/evenfound/even-go/ipfs"

// twig represents candidate transaction in a branch.
type twig struct {
	Hash  ipfs.Hash `json:"hash"`
	Proof ipfs.Hash `json:"proof"`
}

type twigs []twig

func newTwig(h ipfs.Hash) twig {
	return twig{
		Hash:  h,
		Proof: "Proof",
	}
}
