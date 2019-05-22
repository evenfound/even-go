module api

go 1.12

require (
	github.com/OpenBazaar/jsonpb v0.0.0-20171123000858-37d32ddf4eef
	github.com/OpenBazaar/spvwallet v0.0.0-20190417151124-49419d61fdff
	github.com/OpenBazaar/wallet-interface v0.0.0-20190411204206-5b458c29c191
	github.com/btcsuite/btcd v0.0.0-20190427004231-96897255fd17
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/evenfound/even-go/core v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/ipfs v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/pb v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/repo v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/schema v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.3.1
	github.com/gorilla/websocket v1.4.0
	github.com/ipfs/go-cid v0.0.2
	github.com/ipfs/go-ipfs v0.4.20
	github.com/ipfs/go-ipfs-ds-help v0.0.1
	github.com/ipfs/go-path v0.0.3
	github.com/libp2p/go-libp2p-crypto v0.0.2
	github.com/libp2p/go-libp2p-kad-dht v0.0.8
	github.com/libp2p/go-libp2p-peer v0.1.1
	github.com/libp2p/go-libp2p-peerstore v0.0.6
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/multiformats/go-multihash v0.0.5
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
)

replace github.com/evenfound/even-go/core => ../core

replace github.com/evenfound/even-go/ipfs => ../ipfs

replace github.com/evenfound/even-go/pb => ../pb

replace github.com/evenfound/even-go/net => ../net

replace github.com/evenfound/even-go/repo => ../repo

replace github.com/evenfound/even-go/schema => ../schema
