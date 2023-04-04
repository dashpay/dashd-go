module github.com/dashpay/dashd-go/btcutil

go 1.17

require (
	github.com/aead/siphash v1.0.1
	github.com/dashpay/dashd-go s-dashevo-dashpay
	github.com/dashpay/dashd-go/btcec s-dashevo-dashpay
	github.com/davecgh/go-spew v1.1.1
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1
	github.com/kkdai/bstream v1.0.0
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292
)

replace (
	github.com/dashpay/dashd-go => ../
	github.com/dashpay/dashd-go/btcec => ../btcec
)
