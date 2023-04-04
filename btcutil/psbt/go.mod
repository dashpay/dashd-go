module github.com/dashpay/dashd-go/btcutil/psbt

go 1.17

require (
	github.com/dashpay/dashd-go s-dashevo-dashpay
	github.com/dashpay/dashd-go/btcec s-dashevo-dashpay
	github.com/dashpay/dashd-go/btcutil s-dashevo-dashpay
	github.com/davecgh/go-spew v1.1.1
)

require (
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
)

replace github.com/dashpay/dashd-go/btcec => ../../btcec

replace github.com/dashpay/dashd-go/btcutil => ../

replace github.com/dashpay/dashd-go => ../..
