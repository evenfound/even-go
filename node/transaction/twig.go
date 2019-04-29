package transaction

// twig represents candidate transaction in a branch.
type twig struct {
	Hash  Hash `json:"hash"`
	Proof Hash `json:"proof"`
}

type twigs []twig

func newTwig(h Hash) twig {
	return twig{
		Hash:  h,
		Proof: "Proof",
	}
}
