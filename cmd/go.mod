module command

go 1.12

require (
	github.com/OpenBazaar/go-blockstackclient v0.0.0-20170922215143-28fae9038857
	github.com/OpenBazaar/jsonpb v0.0.0-20171123000858-37d32ddf4eef
	github.com/OpenBazaar/wallet-interface v0.0.0-20190411204206-5b458c29c191
	github.com/antlr/antlr4 v0.0.0-20190518164840-edae2a1c9b4b
	github.com/btcsuite/btcd v0.0.0-20190427004231-96897255fd17
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/d5/tengo v1.24.1
	github.com/evenfound/even-go/core v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/ipfs v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/net v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/pb v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.7.0
	github.com/golang/protobuf v1.3.1
	github.com/ipfs/go-ipfs v0.4.20
	github.com/ipfs/go-ipfs-api v0.0.1
	github.com/jawher/mow.cli v1.1.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/pkg/errors v0.8.1
	github.com/smartystreets/goconvey v0.0.0-20190222223459-a17d461953aa
	github.com/tyler-smith/go-bip39 v1.0.0
	github.com/urfave/cli v1.20.0
	github.com/ztrue/tracerr v0.3.0
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f
	golang.org/x/net v0.0.0-20190520210107-018c4d40a106
	google.golang.org/grpc v1.20.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace github.com/evenfound/even-go/ipfs => ../ipfs

replace github.com/evenfound/even-go/core => ../core

replace github.com/evenfound/even-go/pb => ../pb

replace github.com/evenfound/even-go/net => ../net

replace github.com/evenfound/even-go/repo => ../repo
