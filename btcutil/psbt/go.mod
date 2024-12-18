module github.com/dashpay/dashd-go/btcutil/psbt

go 1.23

toolchain go1.23.2

require (
	github.com/dashpay/dashd-go v0.26.0
	github.com/dashpay/dashd-go/btcec/v2 v2.2.0
	github.com/dashpay/dashd-go/btcutil v1.3.0
	github.com/davecgh/go-spew v1.1.1
)

require (
	github.com/btcsuite/btclog v1.0.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
)

replace github.com/dashpay/dashd-go/btcec/v2 => ../../btcec

replace github.com/dashpay/dashd-go/btcutil => ../

replace github.com/dashpay/dashd-go => ../..
