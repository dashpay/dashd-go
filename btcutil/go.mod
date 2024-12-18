module github.com/dashpay/dashd-go/btcutil

go 1.23

toolchain go1.23.2

require (
	github.com/aead/siphash v1.0.1
	github.com/dashpay/dashd-go v0.24.0
	github.com/dashpay/dashd-go/btcec/v2 v2.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0
	github.com/kkdai/bstream v1.0.0
	golang.org/x/crypto v0.31.0
)

require github.com/btcsuite/btclog v1.0.0 // indirect

replace (
	github.com/dashpay/dashd-go => ../
	github.com/dashpay/dashd-go/btcec/v2 => ../btcec
)
