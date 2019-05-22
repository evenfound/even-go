module core

go 1.12

require (
	github.com/OpenBazaar/bitcoind-wallet v0.0.0-20180924194541-6c59e2405456 // indirect
	github.com/OpenBazaar/jsonpb v0.0.0-20171123000858-37d32ddf4eef // indirect
	github.com/OpenBazaar/spvwallet v0.0.0-20190417151124-49419d61fdff // indirect
	github.com/OpenBazaar/zcashd-wallet v0.0.0-20180924204619-2b76590b8874 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/btcsuite/btcwallet v0.0.0-20190424224017-9d95f76e99a7 // indirect
	github.com/ccding/go-stun v0.0.0-20180726100737-be486d185f3d // indirect
	github.com/cevaris/ordered_map v0.0.0-20190319150403-3adeae072e73 // indirect
	github.com/cpacia/bchutil v0.0.0-20181003130114-b126f6a35b6c // indirect
	github.com/evenfound/even-go/ipfs v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/net v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/pb v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/schema v0.0.0-00010101000000-000000000000 // indirect
	github.com/evenfound/even-go/server/api v0.0.0-00010101000000-000000000000
	github.com/ipfs/go-cid v0.0.2
	github.com/ipfs/go-ipfs v0.4.20
	github.com/ipfs/go-ipfs-ds-help v0.0.1
	github.com/ipfs/go-merkledag v0.0.3
	github.com/ipfs/go-mfs v0.0.4
	github.com/ipfs/go-unixfs v0.0.4
	github.com/ipfs/interface-go-ipfs-core v0.0.6
	github.com/kennygrant/sanitize v1.2.4
	github.com/libp2p/go-libp2p-crypto v0.0.1
	github.com/libp2p/go-libp2p-peer v0.1.0
	github.com/libp2p/go-libp2p-routing v0.0.1
	github.com/multiformats/go-multiaddr v0.0.3
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/yawning/bulb v0.0.0-20170405033506-85d80d893c3d // indirect
	golang.org/x/net v0.0.0-20190514140710-3ec191127204
	golang.org/x/text v0.3.1-0.20180807135948-17ff2d5776d2 // indirect
)

replace github.com/evenfound/even-go/net => ../net

replace github.com/evenfound/even-go/ipfs => ../ipfs

replace github.com/evenfound/even-go/server/api => ../server/api

replace github.com/evenfound/even-go/pb => ../pb

replace github.com/evenfound/even-go/repo => ../repo

replace github.com/evenfound/even-go/schema => ../schema
