module transaction

go 1.12

require (
	github.com/cesanta/ubjson v0.0.0-20160505143622-3758bbe9c1d8
	github.com/evenfound/even-go/core v0.0.0-00010101000000-000000000000
	github.com/evenfound/even-go/ipfs v0.0.0-00010101000000-000000000000
	github.com/smartystreets/goconvey v0.0.0-20190222223459-a17d461953aa
)

replace github.com/evenfound/even-go/core => ../core

replace github.com/evenfound/even-go/ipfs => ../ipfs

replace github.com/evenfound/even-go/pb => ../pb

replace github.com/evenfound/even-go/repo => ../repo

replace github.com/evenfound/even-go/schema => ../schema

replace github.com/evenfound/even-go/server/api => ../server/api

replace github.com/evenfound/even-go/net => ../net
