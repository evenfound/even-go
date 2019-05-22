module repo

go 1.12

require (
	github.com/OpenBazaar/bitcoind-wallet v0.0.0-20180924194541-6c59e2405456 // indirect
	github.com/OpenBazaar/jsonpb v0.0.0-20171123000858-37d32ddf4eef
	github.com/OpenBazaar/spvwallet v0.0.0-20190417151124-49419d61fdff // indirect
	github.com/OpenBazaar/wallet-interface v0.0.0-20190411204206-5b458c29c191
	github.com/OpenBazaar/zcashd-wallet v0.0.0-20180924204619-2b76590b8874
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/btcsuite/btcd v0.0.0-20190427004231-96897255fd17
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/btcsuite/btcwallet v0.0.0-20190424224017-9d95f76e99a7 // indirect
	github.com/cevaris/ordered_map v0.0.0-20190319150403-3adeae072e73 // indirect
	github.com/cpacia/bchutil v0.0.0-20181003130114-b126f6a35b6c
	github.com/evenfound/even-go/ipfs v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/pb v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/schema v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.3.1
	github.com/ipfs/go-ipfs v0.4.20
	github.com/ipfs/go-path v0.0.3
	github.com/libp2p/go-libp2p-crypto v0.0.1
	github.com/libp2p/go-libp2p-peer v0.1.0
	github.com/multiformats/go-multihash v0.0.5
	github.com/mutecomm/go-sqlcipher v0.0.0-20190227152316-55dbde17881f
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/tyler-smith/go-bip39 v1.0.0
	golang.org/x/text v0.3.1-0.20180807135948-17ff2d5776d2 // indirect
)

replace github.com/evenfound/even-go/ipfs => ../ipfs

replace github.com/evenfound/even-go/core => ../core

replace github.com/evenfound/even-go/pb => ../pb

replace github.com/evenfound/even-go/net => ../net

replace github.com/evenfound/even-go/schema => ../schema
