module github.com/dashpay/dashd-go/btcec/v2

go 1.23

require (
	github.com/dashpay/dashd-go v0.26.0
	github.com/davecgh/go-spew v1.1.1
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0
)

require github.com/decred/dcrd/crypto/blake256 v1.0.1 // indirect

replace github.com/dashpay/dashd-go => ../
