module github.com/dashpay/dashd-go/btcec

go 1.17

require (
	github.com/dashpay/dashd-go s-dashevo-dashpay
	github.com/davecgh/go-spew v1.1.1
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1
)

require github.com/decred/dcrd/crypto/blake256 v1.0.0 // indirect

replace github.com/dashpay/dashd-go => ../
