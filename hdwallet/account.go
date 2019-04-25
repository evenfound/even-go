package hdwallet

type account struct {
	Index   uint32 `json:"index"`
	Address string `json:"address"`
}

// newAccount creates another instance of account.
func newAccount(index uint32, address string) *account {
	return &account{
		Index:   index,
		Address: address,
	}
}
