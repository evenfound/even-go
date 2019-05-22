module interop

go 1.12

require (
	github.com/btcsuite/btcd v0.0.0-20190427004231-96897255fd17 // indirect
	github.com/d5/tengo v1.24.1
	github.com/evenfound/even-go/crypto v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f // indirect
)

replace github.com/evenfound/even-go/crypto => ../../crypto
