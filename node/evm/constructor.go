package evm

// New creates another instance of the EVM.
func New() Interface {
	return newTengoVM()
}
