module github.com/dashpay/dashd-go

require (
	github.com/btcsuite/btclog v1.0.0
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd
	github.com/btcsuite/goleveldb v1.0.0
	github.com/btcsuite/websocket v0.0.0-20150119174127-31079b680792
	github.com/btcsuite/winsvc v1.0.0
	github.com/dashpay/dashd-go/btcec/v2 v2.2.0
	github.com/dashpay/dashd-go/btcutil v1.3.0
	github.com/davecgh/go-spew v1.1.1
	github.com/decred/dcrd/lru v1.1.3
	github.com/jessevdk/go-flags v1.6.1
	github.com/jrick/logrotate v1.1.2
	golang.org/x/crypto v0.31.0
)

require (
	github.com/aead/siphash v1.0.1 // indirect
	github.com/btcsuite/snappy-go v1.0.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/howeyc/fsnotify v0.9.0 // indirect
	github.com/hpcloud/tail v2.10.6-bug100770-inotify-leak+incompatible // indirect
	github.com/kkdai/bstream v1.0.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	launchpad.net/tomb v0.0.0-20140529072043-000000000018 // indirect
)

replace (
	github.com/dashpay/dashd-go => ./
	github.com/dashpay/dashd-go/btcec/v2 => ./btcec
	github.com/dashpay/dashd-go/btcutil => ./btcutil
)

retract (
	v0.22.0-beta
	v0.2.0
	v0.1.4
	v0.1.3
	v0.1.2
	v0.1.1
	v0.1.0
)

go 1.23

toolchain go1.23.2
