module github.com/dashevo/dashd-go/btcutil

go 1.16

require (
	github.com/aead/siphash v1.0.1
	github.com/dashevo/dashd-go v0.23.0-test.0
	github.com/dashevo/dashd-go/btcec/v2 v2.0.0-test.0
	github.com/davecgh/go-spew v1.1.1
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1
	github.com/kkdai/bstream v0.0.0-20161212061736-f391b8402d23
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
)

retract (
	v1.0.1
	v1.0.0
	v0.1.1
	v0.1.0
)
