package config

const (
	// RPCAddress is the address of a gRPC server to connect to.
	RPCAddress = "localhost:8090"

	// FilePrefix is the filepath URI prefix.
	FilePrefix = "file://"

	// IpfsPrefix is the IPFS URI prefix.
	IpfsPrefix = "/ipfs/"

	// DefaultEntryFunction is the name of SC function which is called if no entry function specified.
	DefaultEntryFunction = "default"
)
