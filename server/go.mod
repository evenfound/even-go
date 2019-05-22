module server

go 1.12

require (
	github.com/evenfound/even-go/evm v0.0.0-00010101000000-000000000000 // indirect
	github.com/evenfound/even-go/hdwallet v0.0.0-00010101000000-000000000000 // indirect
	github.com/evenfound/even-go/server/api v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/server/handlers v0.0.0-00010101000000-000000000000
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/grpc v1.20.1
)

replace github.com/evenfound/even-go/core => ../core

replace github.com/evenfound/even-go/ipfs => ../ipfs

replace github.com/evenfound/even-go/pb => ../pb

replace github.com/evenfound/even-go/repo => ../repo

replace github.com/evenfound/even-go/schema => ../schema

replace github.com/evenfound/even-go/server/api => ./api

replace github.com/evenfound/even-go/server/handlers => ./handlers

replace github.com/evenfound/even-go/net => ../net

replace github.com/evenfound/even-go/crypto => ../crypto

replace github.com/evenfound/even-go/evm => ../evm

replace github.com/evenfound/even-go/evm/interop => ../evm/interop

replace github.com/evenfound/even-go/hdwallet => ../hdwallet

replace github.com/evenfound/even-go/transaction => ../transaction
