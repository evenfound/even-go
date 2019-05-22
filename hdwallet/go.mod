module hdwallet

go 1.12

require (
	github.com/alexmullins/zip v0.0.0-20180717182244-4affb64b04d0
	github.com/btcsuite/btcd v0.0.0-20190427004231-96897255fd17
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/edunuzzi/go-bip44 v0.0.0-20190109211530-eb6b7decf5cc
	github.com/evenfound/even-go/core v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/ipfs v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/repo v0.0.0-00010101000000-000000000000 // indirect
	github.com/evenfound/even-go/schema v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/transaction v0.0.0-00010101000000-000000000000
	github.com/ipfs/go-cid v0.0.2
	github.com/smartystreets/goconvey v0.0.0-20190222223459-a17d461953aa
)

replace github.com/evenfound/even-go/core => ../core

replace github.com/evenfound/even-go/pb => ../pb

replace github.com/evenfound/even-go/net => ../net

replace github.com/evenfound/even-go/ipfs => ../ipfs

replace github.com/evenfound/even-go/repo => ../repo

replace github.com/evenfound/even-go/server/api => ../server/api

replace github.com/evenfound/even-go/schema => ../schema

replace github.com/evenfound/even-go/transaction => ../transaction
