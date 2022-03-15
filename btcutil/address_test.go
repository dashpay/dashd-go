// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcutil_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/dashevo/dashd-go/btcutil"
	"github.com/dashevo/dashd-go/chaincfg"
	"github.com/dashevo/dashd-go/wire"
	"golang.org/x/crypto/ripemd160"
)

type CustomParamStruct struct {
	Net              wire.BitcoinNet
	PubKeyHashAddrID byte
	ScriptHashAddrID byte
	Bech32HRPSegwit  string
}

var CustomParams = CustomParamStruct{
	Net:              0xdbb6c0fb, // litecoin mainnet HD version bytes
	PubKeyHashAddrID: 0x30,       // starts with L
	ScriptHashAddrID: 0x32,       // starts with M
	Bech32HRPSegwit:  "ltc",      // starts with ltc
}

// We use this function to be able to test functionality in DecodeAddress for
// defaultNet addresses
func applyCustomParams(params chaincfg.Params, customParams CustomParamStruct) chaincfg.Params {
	params.Net = customParams.Net
	params.PubKeyHashAddrID = customParams.PubKeyHashAddrID
	params.ScriptHashAddrID = customParams.ScriptHashAddrID
	params.Bech32HRPSegwit = customParams.Bech32HRPSegwit
	return params
}

var customParams = applyCustomParams(chaincfg.MainNetParams, CustomParams)

func TestAddresses(t *testing.T) {
	tests := []struct {
		name    string
		addr    string
		encoded string
		valid   bool
		result  btcutil.Address
		f       func() (btcutil.Address, error)
		net     *chaincfg.Params
	}{
		// Positive P2PKH tests.
		{
			name:    "mainnet p2pkh",
			addr:    "Xbp12jNUxd588MppSqcmvb5zfgJNjtLi1Q",
			encoded: "Xbp12jNUxd588MppSqcmvb5zfgJNjtLi1Q",
			valid:   true,
			result: btcutil.TstAddressPubKeyHash(
				[ripemd160.Size]byte{
					0x0c, 0x52, 0xb7, 0x2d, 0x1a, 0x00, 0xbe, 0x72, 0x42, 0x43,
					0xc6, 0x24, 0xee, 0xa9, 0x09, 0xd5, 0xde, 0x6d, 0x2a, 0x09},
				chaincfg.MainNetParams.PubKeyHashAddrID),
			f: func() (btcutil.Address, error) {
				pkHash := []byte{
					0x0c, 0x52, 0xb7, 0x2d, 0x1a, 0x00, 0xbe, 0x72, 0x42, 0x43,
					0xc6, 0x24, 0xee, 0xa9, 0x09, 0xd5, 0xde, 0x6d, 0x2a, 0x09}
				return btcutil.NewAddressPubKeyHash(pkHash, &chaincfg.MainNetParams)
			},
			net: &chaincfg.MainNetParams,
		},
		{
			name:    "mainnet p2pkh 2",
			addr:    "XrfVnPa5LNzHddRc6npe1YLJE4WsV3WLjJ",
			encoded: "XrfVnPa5LNzHddRc6npe1YLJE4WsV3WLjJ",
			valid:   true,
			result: btcutil.TstAddressPubKeyHash(
				[ripemd160.Size]byte{
					0xaf, 0x40, 0xdc, 0x36, 0x83, 0x6c, 0xa4, 0x11, 0xaf, 0x3c,
					0x05, 0xfc, 0x06, 0x32, 0xbb, 0x49, 0x3d, 0xab, 0xc1, 0x8c},
				chaincfg.MainNetParams.PubKeyHashAddrID),
			f: func() (btcutil.Address, error) {
				pkHash := []byte{
					0xaf, 0x40, 0xdc, 0x36, 0x83, 0x6c, 0xa4, 0x11, 0xaf, 0x3c,
					0x05, 0xfc, 0x06, 0x32, 0xbb, 0x49, 0x3d, 0xab, 0xc1, 0x8c}
				return btcutil.NewAddressPubKeyHash(pkHash, &chaincfg.MainNetParams)
			},
			net: &chaincfg.MainNetParams,
		},
		{
			name:    "litecoin mainnet p2pkh",
			addr:    "LM2WMpR1Rp6j3Sa59cMXMs1SPzj9eXpGc1",
			encoded: "LM2WMpR1Rp6j3Sa59cMXMs1SPzj9eXpGc1",
			valid:   true,
			result: btcutil.TstAddressPubKeyHash(
				[ripemd160.Size]byte{
					0x13, 0xc6, 0x0d, 0x8e, 0x68, 0xd7, 0x34, 0x9f, 0x5b, 0x4c,
					0xa3, 0x62, 0xc3, 0x95, 0x4b, 0x15, 0x04, 0x50, 0x61, 0xb1},
				CustomParams.PubKeyHashAddrID),
			f: func() (btcutil.Address, error) {
				pkHash := []byte{
					0x13, 0xc6, 0x0d, 0x8e, 0x68, 0xd7, 0x34, 0x9f, 0x5b, 0x4c,
					0xa3, 0x62, 0xc3, 0x95, 0x4b, 0x15, 0x04, 0x50, 0x61, 0xb1}
				return btcutil.NewAddressPubKeyHash(pkHash, &customParams)
			},
			net: &customParams,
		},
		{
			name:    "testnet p2pkh",
			addr:    "yLdJV43FeP3zAZMBrUdtygj3viNBnfcGbG",
			encoded: "yLdJV43FeP3zAZMBrUdtygj3viNBnfcGbG",
			valid:   true,
			result: btcutil.TstAddressPubKeyHash(
				[ripemd160.Size]byte{
					0x03, 0x60, 0x81, 0x3d, 0xf9, 0xb3, 0x32, 0x5e, 0x7a, 0x21,
					0x97, 0x20, 0x29, 0x94, 0xcb, 0x2a, 0x0a, 0xc2, 0xd3, 0xb3},
				chaincfg.TestNet3Params.PubKeyHashAddrID),
			f: func() (btcutil.Address, error) {
				pkHash := []byte{
					0x03, 0x60, 0x81, 0x3d, 0xf9, 0xb3, 0x32, 0x5e, 0x7a, 0x21,
					0x97, 0x20, 0x29, 0x94, 0xcb, 0x2a, 0x0a, 0xc2, 0xd3, 0xb3}
				return btcutil.NewAddressPubKeyHash(pkHash, &chaincfg.TestNet3Params)
			},
			net: &chaincfg.TestNet3Params,
		},

		// Negative P2PKH tests.
		{
			name:  "p2pkh wrong hash length",
			addr:  "",
			valid: false,
			f: func() (btcutil.Address, error) {
				pkHash := []byte{
					0x00, 0x0e, 0xf0, 0x30, 0x10, 0x7f, 0xd2, 0x6e, 0x0b, 0x6b,
					0xf4, 0x05, 0x12, 0xbc, 0xa2, 0xce, 0xb1, 0xdd, 0x80, 0xad,
					0xaa}
				return btcutil.NewAddressPubKeyHash(pkHash, &chaincfg.MainNetParams)
			},
			net: &chaincfg.MainNetParams,
		},
		{
			name:  "p2pkh bad checksum",
			addr:  "1MirQ9bwyQcGVJPwKUgapu5ouK2E2Ey4gY",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},

		// Positive P2SH tests.
		{
			// Taken from transaction:
			// 10b2099648910f6a6bc944a7efa528a9e48dd357a2886a4e40a2b6d42f10662e.
			name:    "mainnet p2sh",
			addr:    "7iGYJ49wQWsyreaaaH1zQt9qF2eLTYhwdQ",
			encoded: "7iGYJ49wQWsyreaaaH1zQt9qF2eLTYhwdQ",
			valid:   true,
			result: btcutil.TstAddressScriptHash(
				[ripemd160.Size]byte{
					0xad, 0xf9, 0x4a, 0xd1, 0xf6, 0x33, 0xc9, 0x3b, 0x32, 0x7c,
					0x00, 0x19, 0xe1, 0x96, 0x6d, 0x3b, 0x57, 0x1b, 0x7e, 0x88},
				chaincfg.MainNetParams.ScriptHashAddrID),
			f: func() (btcutil.Address, error) {
				script := []byte{
					0x52, 0x21, 0x02, 0xe6, 0x52, 0xb7, 0xd4, 0xf3,
					0x9c, 0xc4, 0x61, 0x50, 0xd7, 0x39, 0x87, 0xe1,
					0x4f, 0x10, 0x67, 0xa6, 0x95, 0x22, 0x8e, 0x36,
					0x68, 0xfd, 0x6e, 0xb8, 0x60, 0xfd, 0x84, 0x94,
					0x61, 0xeb, 0xa1, 0x21, 0x03, 0x42, 0xd7, 0x9b,
					0xcc, 0xbf, 0x02, 0x16, 0x31, 0x9d, 0x51, 0xe5,
					0x83, 0x14, 0xfc, 0x1e, 0x66, 0xed, 0x47, 0xc7,
					0x97, 0xaf, 0xe5, 0x8f, 0x90, 0x57, 0xe6, 0x68,
					0x25, 0x40, 0xdf, 0x95, 0x11, 0x21, 0x03, 0xfd,
					0x0a, 0xc2, 0xfb, 0xd0, 0xf8, 0xb1, 0x82, 0xaa,
					0xfc, 0x99, 0xbf, 0x77, 0x56, 0x45, 0x62, 0xe4,
					0x3e, 0x99, 0x9c, 0xc3, 0xef, 0x34, 0xf6, 0x71,
					0xce, 0x1f, 0xcc, 0x11, 0x2a, 0xc3, 0x9c, 0x53,
					0xae}
				return btcutil.NewAddressScriptHash(script, &chaincfg.MainNetParams)
			},
			net: &chaincfg.MainNetParams,
		},
		{
			name:    "litecoin mainnet P2SH ",
			addr:    "MVcg9uEvtWuP5N6V48EHfEtbz48qR8TKZ9",
			encoded: "MVcg9uEvtWuP5N6V48EHfEtbz48qR8TKZ9",
			valid:   true,
			result: btcutil.TstAddressScriptHash(
				[ripemd160.Size]byte{
					0xee, 0x34, 0xac, 0x67, 0x6b, 0xda, 0xf6, 0xe3, 0x70, 0xc8,
					0xc8, 0x20, 0xb9, 0x48, 0xed, 0xfa, 0xd3, 0xa8, 0x73, 0xd8},
				CustomParams.ScriptHashAddrID),
			f: func() (btcutil.Address, error) {
				pkHash := []byte{
					0xEE, 0x34, 0xAC, 0x67, 0x6B, 0xDA, 0xF6, 0xE3, 0x70, 0xC8,
					0xC8, 0x20, 0xB9, 0x48, 0xED, 0xFA, 0xD3, 0xA8, 0x73, 0xD8}
				return btcutil.NewAddressScriptHashFromHash(pkHash, &customParams)
			},
			net: &customParams,
		},
		{
			// Taken from transaction:
			// 4604215e1e9089eb46e407daa07fa30b0296bf8896cbfff13843ba81624cb910
			name:    "mainnet p2sh 2",
			addr:    "7pdLznbbC1zQyZaM1H2iNeiPM4GBPsPE4y",
			encoded: "7pdLznbbC1zQyZaM1H2iNeiPM4GBPsPE4y",
			valid:   true,
			result: btcutil.TstAddressScriptHash(
				[ripemd160.Size]byte{
					0xf3, 0xb9, 0x59, 0xee, 0xa9, 0x10, 0xc3, 0xbb, 0x69, 0x4b,
					0x93, 0xd0, 0xc3, 0xb3, 0xf5, 0x95, 0x8a, 0x9a, 0x2e, 0x23},
				chaincfg.MainNetParams.ScriptHashAddrID),
			f: func() (btcutil.Address, error) {
				hash := []byte{
					0xf3, 0xb9, 0x59, 0xee, 0xa9, 0x10, 0xc3, 0xbb, 0x69, 0x4b,
					0x93, 0xd0, 0xc3, 0xb3, 0xf5, 0x95, 0x8a, 0x9a, 0x2e, 0x23}
				return btcutil.NewAddressScriptHashFromHash(hash, &chaincfg.MainNetParams)
			},
			net: &chaincfg.MainNetParams,
		},
		{
			name:    "testnet p2sh",
			addr:    "92e9x7VTKZP3Rrzc5Y2fq2XkEa31U9njQD",
			encoded: "92e9x7VTKZP3Rrzc5Y2fq2XkEa31U9njQD",
			valid:   true,
			result: btcutil.TstAddressScriptHash(
				[ripemd160.Size]byte{
					0xf3, 0xb9, 0x59, 0xee, 0xa9, 0x10, 0xc3, 0xbb, 0x69, 0x4b,
					0x93, 0xd0, 0xc3, 0xb3, 0xf5, 0x95, 0x8a, 0x9a, 0x2e, 0x23},
				chaincfg.TestNet3Params.ScriptHashAddrID),
			f: func() (btcutil.Address, error) {
				hash := []byte{
					0xf3, 0xb9, 0x59, 0xee, 0xa9, 0x10, 0xc3, 0xbb, 0x69, 0x4b,
					0x93, 0xd0, 0xc3, 0xb3, 0xf5, 0x95, 0x8a, 0x9a, 0x2e, 0x23}
				return btcutil.NewAddressScriptHashFromHash(hash, &chaincfg.TestNet3Params)
			},
			net: &chaincfg.TestNet3Params,
		},

		// Negative P2SH tests.
		{
			name:  "p2sh wrong hash length",
			addr:  "",
			valid: false,
			f: func() (btcutil.Address, error) {
				hash := []byte{
					0x00, 0xf8, 0x15, 0xb0, 0x36, 0xd9, 0xbb, 0xbc, 0xe5, 0xe9,
					0xf2, 0xa0, 0x0a, 0xbd, 0x1b, 0xf3, 0xdc, 0x91, 0xe9, 0x55,
					0x10}
				return btcutil.NewAddressScriptHashFromHash(hash, &chaincfg.MainNetParams)
			},
			net: &chaincfg.MainNetParams,
		},

		// Positive P2PK tests.
		{
			name:    "mainnet p2pk compressed (0x02)",
			addr:    "02192d74d0cb94344c9569c2e77901573d8d7903c3ebec3a957724895dca52c6b4",
			encoded: "Xct6vgwwvzh7wzoRtJrHJp289Bwx2A6x5S",
			valid:   true,
			result: btcutil.TstAddressPubKey(
				[]byte{
					0x02, 0x19, 0x2d, 0x74, 0xd0, 0xcb, 0x94, 0x34, 0x4c, 0x95,
					0x69, 0xc2, 0xe7, 0x79, 0x01, 0x57, 0x3d, 0x8d, 0x79, 0x03,
					0xc3, 0xeb, 0xec, 0x3a, 0x95, 0x77, 0x24, 0x89, 0x5d, 0xca,
					0x52, 0xc6, 0xb4},
				btcutil.PKFCompressed, chaincfg.MainNetParams.PubKeyHashAddrID),
			f: func() (btcutil.Address, error) {
				serializedPubKey := []byte{
					0x02, 0x19, 0x2d, 0x74, 0xd0, 0xcb, 0x94, 0x34, 0x4c, 0x95,
					0x69, 0xc2, 0xe7, 0x79, 0x01, 0x57, 0x3d, 0x8d, 0x79, 0x03,
					0xc3, 0xeb, 0xec, 0x3a, 0x95, 0x77, 0x24, 0x89, 0x5d, 0xca,
					0x52, 0xc6, 0xb4}
				return btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.MainNetParams)
			},
			net: &chaincfg.MainNetParams,
		},
		{
			name:    "mainnet p2pk compressed (0x03)",
			addr:    "03b0bd634234abbb1ba1e986e884185c61cf43e001f9137f23c2c409273eb16e65",
			encoded: "XfZ7zd2N99ugwAQiPdXccwJHw7cvKr9Wqq",
			valid:   true,
			result: btcutil.TstAddressPubKey(
				[]byte{
					0x03, 0xb0, 0xbd, 0x63, 0x42, 0x34, 0xab, 0xbb, 0x1b, 0xa1,
					0xe9, 0x86, 0xe8, 0x84, 0x18, 0x5c, 0x61, 0xcf, 0x43, 0xe0,
					0x01, 0xf9, 0x13, 0x7f, 0x23, 0xc2, 0xc4, 0x09, 0x27, 0x3e,
					0xb1, 0x6e, 0x65},
				btcutil.PKFCompressed, chaincfg.MainNetParams.PubKeyHashAddrID),
			f: func() (btcutil.Address, error) {
				serializedPubKey := []byte{
					0x03, 0xb0, 0xbd, 0x63, 0x42, 0x34, 0xab, 0xbb, 0x1b, 0xa1,
					0xe9, 0x86, 0xe8, 0x84, 0x18, 0x5c, 0x61, 0xcf, 0x43, 0xe0,
					0x01, 0xf9, 0x13, 0x7f, 0x23, 0xc2, 0xc4, 0x09, 0x27, 0x3e,
					0xb1, 0x6e, 0x65}
				return btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.MainNetParams)
			},
			net: &chaincfg.MainNetParams,
		},
		{
			name: "mainnet p2pk uncompressed (0x04)",
			addr: "0411db93e1dcdb8a016b49840f8c53bc1eb68a382e97b1482ecad7b148a6909a5cb2" +
				"e0eaddfb84ccf9744464f82e160bfa9b8b64f9d4c03f999b8643f656b412a3",
			encoded: "XcJSEb79KEeNbwMU7eE27aL5dhDwvjgBuH",
			valid:   true,
			result: btcutil.TstAddressPubKey(
				[]byte{
					0x04, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a, 0x01, 0x6b,
					0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e, 0xb6, 0x8a, 0x38,
					0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca, 0xd7, 0xb1, 0x48, 0xa6,
					0x90, 0x9a, 0x5c, 0xb2, 0xe0, 0xea, 0xdd, 0xfb, 0x84, 0xcc,
					0xf9, 0x74, 0x44, 0x64, 0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b,
					0x8b, 0x64, 0xf9, 0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43,
					0xf6, 0x56, 0xb4, 0x12, 0xa3},
				btcutil.PKFUncompressed, chaincfg.MainNetParams.PubKeyHashAddrID),
			f: func() (btcutil.Address, error) {
				serializedPubKey := []byte{
					0x04, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a, 0x01, 0x6b,
					0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e, 0xb6, 0x8a, 0x38,
					0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca, 0xd7, 0xb1, 0x48, 0xa6,
					0x90, 0x9a, 0x5c, 0xb2, 0xe0, 0xea, 0xdd, 0xfb, 0x84, 0xcc,
					0xf9, 0x74, 0x44, 0x64, 0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b,
					0x8b, 0x64, 0xf9, 0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43,
					0xf6, 0x56, 0xb4, 0x12, 0xa3}
				return btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.MainNetParams)
			},
			net: &chaincfg.MainNetParams,
		},
		{
			name:    "testnet p2pk compressed (0x02)",
			addr:    "02192d74d0cb94344c9569c2e77901573d8d7903c3ebec3a957724895dca52c6b4",
			encoded: "yNWhwe2PNYMCHjiyTAAgLqSURUSKSPuei9",
			valid:   true,
			result: btcutil.TstAddressPubKey(
				[]byte{
					0x02, 0x19, 0x2d, 0x74, 0xd0, 0xcb, 0x94, 0x34, 0x4c, 0x95,
					0x69, 0xc2, 0xe7, 0x79, 0x01, 0x57, 0x3d, 0x8d, 0x79, 0x03,
					0xc3, 0xeb, 0xec, 0x3a, 0x95, 0x77, 0x24, 0x89, 0x5d, 0xca,
					0x52, 0xc6, 0xb4},
				btcutil.PKFCompressed, chaincfg.TestNet3Params.PubKeyHashAddrID),
			f: func() (btcutil.Address, error) {
				serializedPubKey := []byte{
					0x02, 0x19, 0x2d, 0x74, 0xd0, 0xcb, 0x94, 0x34, 0x4c, 0x95,
					0x69, 0xc2, 0xe7, 0x79, 0x01, 0x57, 0x3d, 0x8d, 0x79, 0x03,
					0xc3, 0xeb, 0xec, 0x3a, 0x95, 0x77, 0x24, 0x89, 0x5d, 0xca,
					0x52, 0xc6, 0xb4}
				return btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.TestNet3Params)
			},
			net: &chaincfg.TestNet3Params,
		},
		{
			name:    "testnet p2pk compressed (0x03)",
			addr:    "03b0bd634234abbb1ba1e986e884185c61cf43e001f9137f23c2c409273eb16e65",
			encoded: "yRBj1a6oahZmGuLFxUr1exieDQ7HkviwPq",
			valid:   true,
			result: btcutil.TstAddressPubKey(
				[]byte{
					0x03, 0xb0, 0xbd, 0x63, 0x42, 0x34, 0xab, 0xbb, 0x1b, 0xa1,
					0xe9, 0x86, 0xe8, 0x84, 0x18, 0x5c, 0x61, 0xcf, 0x43, 0xe0,
					0x01, 0xf9, 0x13, 0x7f, 0x23, 0xc2, 0xc4, 0x09, 0x27, 0x3e,
					0xb1, 0x6e, 0x65},
				btcutil.PKFCompressed, chaincfg.TestNet3Params.PubKeyHashAddrID),
			f: func() (btcutil.Address, error) {
				serializedPubKey := []byte{
					0x03, 0xb0, 0xbd, 0x63, 0x42, 0x34, 0xab, 0xbb, 0x1b, 0xa1,
					0xe9, 0x86, 0xe8, 0x84, 0x18, 0x5c, 0x61, 0xcf, 0x43, 0xe0,
					0x01, 0xf9, 0x13, 0x7f, 0x23, 0xc2, 0xc4, 0x09, 0x27, 0x3e,
					0xb1, 0x6e, 0x65}
				return btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.TestNet3Params)
			},
			net: &chaincfg.TestNet3Params,
		},
		{
			name: "testnet p2pk uncompressed (0x04)",
			addr: "0411db93e1dcdb8a016b49840f8c53bc1eb68a382e97b1482ecad7b148a6909a5" +
				"cb2e0eaddfb84ccf9744464f82e160bfa9b8b64f9d4c03f999b8643f656b412a3",
			encoded: "yMw3FYBaknJSwgH1gVYR9bkRuyiKPxp1ya",
			valid:   true,
			result: btcutil.TstAddressPubKey(
				[]byte{
					0x04, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a, 0x01, 0x6b,
					0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e, 0xb6, 0x8a, 0x38,
					0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca, 0xd7, 0xb1, 0x48, 0xa6,
					0x90, 0x9a, 0x5c, 0xb2, 0xe0, 0xea, 0xdd, 0xfb, 0x84, 0xcc,
					0xf9, 0x74, 0x44, 0x64, 0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b,
					0x8b, 0x64, 0xf9, 0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43,
					0xf6, 0x56, 0xb4, 0x12, 0xa3},
				btcutil.PKFUncompressed, chaincfg.TestNet3Params.PubKeyHashAddrID),
			f: func() (btcutil.Address, error) {
				serializedPubKey := []byte{
					0x04, 0x11, 0xdb, 0x93, 0xe1, 0xdc, 0xdb, 0x8a, 0x01, 0x6b,
					0x49, 0x84, 0x0f, 0x8c, 0x53, 0xbc, 0x1e, 0xb6, 0x8a, 0x38,
					0x2e, 0x97, 0xb1, 0x48, 0x2e, 0xca, 0xd7, 0xb1, 0x48, 0xa6,
					0x90, 0x9a, 0x5c, 0xb2, 0xe0, 0xea, 0xdd, 0xfb, 0x84, 0xcc,
					0xf9, 0x74, 0x44, 0x64, 0xf8, 0x2e, 0x16, 0x0b, 0xfa, 0x9b,
					0x8b, 0x64, 0xf9, 0xd4, 0xc0, 0x3f, 0x99, 0x9b, 0x86, 0x43,
					0xf6, 0x56, 0xb4, 0x12, 0xa3}
				return btcutil.NewAddressPubKey(serializedPubKey, &chaincfg.TestNet3Params)
			},
			net: &chaincfg.TestNet3Params,
		},
		// Segwit address tests.
		{
			name:    "segwit mainnet p2wpkh v0",
			addr:    "BC1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4",
			encoded: "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
			valid:   true,
			result: btcutil.TstAddressWitnessPubKeyHash(
				0,
				[20]byte{
					0x75, 0x1e, 0x76, 0xe8, 0x19, 0x91, 0x96, 0xd4, 0x54, 0x94,
					0x1c, 0x45, 0xd1, 0xb3, 0xa3, 0x23, 0xf1, 0x43, 0x3b, 0xd6},
				chaincfg.MainNetParams.Bech32HRPSegwit),
			f: func() (btcutil.Address, error) {
				pkHash := []byte{
					0x75, 0x1e, 0x76, 0xe8, 0x19, 0x91, 0x96, 0xd4, 0x54, 0x94,
					0x1c, 0x45, 0xd1, 0xb3, 0xa3, 0x23, 0xf1, 0x43, 0x3b, 0xd6}
				return btcutil.NewAddressWitnessPubKeyHash(pkHash, &chaincfg.MainNetParams)
			},
			net: &chaincfg.MainNetParams,
		},
		{
			name:    "segwit mainnet p2wsh v0",
			addr:    "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3",
			encoded: "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3",
			valid:   true,
			result: btcutil.TstAddressWitnessScriptHash(
				0,
				[32]byte{
					0x18, 0x63, 0x14, 0x3c, 0x14, 0xc5, 0x16, 0x68,
					0x04, 0xbd, 0x19, 0x20, 0x33, 0x56, 0xda, 0x13,
					0x6c, 0x98, 0x56, 0x78, 0xcd, 0x4d, 0x27, 0xa1,
					0xb8, 0xc6, 0x32, 0x96, 0x04, 0x90, 0x32, 0x62},
				chaincfg.MainNetParams.Bech32HRPSegwit),
			f: func() (btcutil.Address, error) {
				scriptHash := []byte{
					0x18, 0x63, 0x14, 0x3c, 0x14, 0xc5, 0x16, 0x68,
					0x04, 0xbd, 0x19, 0x20, 0x33, 0x56, 0xda, 0x13,
					0x6c, 0x98, 0x56, 0x78, 0xcd, 0x4d, 0x27, 0xa1,
					0xb8, 0xc6, 0x32, 0x96, 0x04, 0x90, 0x32, 0x62}
				return btcutil.NewAddressWitnessScriptHash(scriptHash, &chaincfg.MainNetParams)
			},
			net: &chaincfg.MainNetParams,
		},
		{
			name:    "segwit testnet p2wpkh v0",
			addr:    "tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
			encoded: "tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
			valid:   true,
			result: btcutil.TstAddressWitnessPubKeyHash(
				0,
				[20]byte{
					0x75, 0x1e, 0x76, 0xe8, 0x19, 0x91, 0x96, 0xd4, 0x54, 0x94,
					0x1c, 0x45, 0xd1, 0xb3, 0xa3, 0x23, 0xf1, 0x43, 0x3b, 0xd6},
				chaincfg.TestNet3Params.Bech32HRPSegwit),
			f: func() (btcutil.Address, error) {
				pkHash := []byte{
					0x75, 0x1e, 0x76, 0xe8, 0x19, 0x91, 0x96, 0xd4, 0x54, 0x94,
					0x1c, 0x45, 0xd1, 0xb3, 0xa3, 0x23, 0xf1, 0x43, 0x3b, 0xd6}
				return btcutil.NewAddressWitnessPubKeyHash(pkHash, &chaincfg.TestNet3Params)
			},
			net: &chaincfg.TestNet3Params,
		},
		{
			name:    "segwit testnet p2wsh v0",
			addr:    "tb1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q0sl5k7",
			encoded: "tb1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q0sl5k7",
			valid:   true,
			result: btcutil.TstAddressWitnessScriptHash(
				0,
				[32]byte{
					0x18, 0x63, 0x14, 0x3c, 0x14, 0xc5, 0x16, 0x68,
					0x04, 0xbd, 0x19, 0x20, 0x33, 0x56, 0xda, 0x13,
					0x6c, 0x98, 0x56, 0x78, 0xcd, 0x4d, 0x27, 0xa1,
					0xb8, 0xc6, 0x32, 0x96, 0x04, 0x90, 0x32, 0x62},
				chaincfg.TestNet3Params.Bech32HRPSegwit),
			f: func() (btcutil.Address, error) {
				scriptHash := []byte{
					0x18, 0x63, 0x14, 0x3c, 0x14, 0xc5, 0x16, 0x68,
					0x04, 0xbd, 0x19, 0x20, 0x33, 0x56, 0xda, 0x13,
					0x6c, 0x98, 0x56, 0x78, 0xcd, 0x4d, 0x27, 0xa1,
					0xb8, 0xc6, 0x32, 0x96, 0x04, 0x90, 0x32, 0x62}
				return btcutil.NewAddressWitnessScriptHash(scriptHash, &chaincfg.TestNet3Params)
			},
			net: &chaincfg.TestNet3Params,
		},
		{
			name:    "segwit testnet p2wsh witness v0",
			addr:    "tb1qqqqqp399et2xygdj5xreqhjjvcmzhxw4aywxecjdzew6hylgvsesrxh6hy",
			encoded: "tb1qqqqqp399et2xygdj5xreqhjjvcmzhxw4aywxecjdzew6hylgvsesrxh6hy",
			valid:   true,
			result: btcutil.TstAddressWitnessScriptHash(
				0,
				[32]byte{
					0x00, 0x00, 0x00, 0xc4, 0xa5, 0xca, 0xd4, 0x62,
					0x21, 0xb2, 0xa1, 0x87, 0x90, 0x5e, 0x52, 0x66,
					0x36, 0x2b, 0x99, 0xd5, 0xe9, 0x1c, 0x6c, 0xe2,
					0x4d, 0x16, 0x5d, 0xab, 0x93, 0xe8, 0x64, 0x33},
				chaincfg.TestNet3Params.Bech32HRPSegwit),
			f: func() (btcutil.Address, error) {
				scriptHash := []byte{
					0x00, 0x00, 0x00, 0xc4, 0xa5, 0xca, 0xd4, 0x62,
					0x21, 0xb2, 0xa1, 0x87, 0x90, 0x5e, 0x52, 0x66,
					0x36, 0x2b, 0x99, 0xd5, 0xe9, 0x1c, 0x6c, 0xe2,
					0x4d, 0x16, 0x5d, 0xab, 0x93, 0xe8, 0x64, 0x33}
				return btcutil.NewAddressWitnessScriptHash(scriptHash, &chaincfg.TestNet3Params)
			},
			net: &chaincfg.TestNet3Params,
		},
		{
			name:    "segwit litecoin mainnet p2wpkh v0",
			addr:    "LTC1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KGMN4N9",
			encoded: "ltc1qw508d6qejxtdg4y5r3zarvary0c5xw7kgmn4n9",
			valid:   true,
			result: btcutil.TstAddressWitnessPubKeyHash(
				0,
				[20]byte{
					0x75, 0x1e, 0x76, 0xe8, 0x19, 0x91, 0x96, 0xd4, 0x54, 0x94,
					0x1c, 0x45, 0xd1, 0xb3, 0xa3, 0x23, 0xf1, 0x43, 0x3b, 0xd6},
				CustomParams.Bech32HRPSegwit,
			),
			f: func() (btcutil.Address, error) {
				pkHash := []byte{
					0x75, 0x1e, 0x76, 0xe8, 0x19, 0x91, 0x96, 0xd4, 0x54, 0x94,
					0x1c, 0x45, 0xd1, 0xb3, 0xa3, 0x23, 0xf1, 0x43, 0x3b, 0xd6}
				return btcutil.NewAddressWitnessPubKeyHash(pkHash, &customParams)
			},
			net: &customParams,
		},

		// P2TR address tests.
		{
			name:    "segwit v1 mainnet p2tr",
			addr:    "bc1paardr2nczq0rx5rqpfwnvpzm497zvux64y0f7wjgcs7xuuuh2nnqwr2d5c",
			encoded: "bc1paardr2nczq0rx5rqpfwnvpzm497zvux64y0f7wjgcs7xuuuh2nnqwr2d5c",
			valid:   true,
			result: btcutil.TstAddressTaproot(
				1, [32]byte{
					0xef, 0x46, 0xd1, 0xaa, 0x78, 0x10, 0x1e, 0x33,
					0x50, 0x60, 0x0a, 0x5d, 0x36, 0x04, 0x5b, 0xa9,
					0x7c, 0x26, 0x70, 0xda, 0xa9, 0x1e, 0x9f, 0x3a,
					0x48, 0xc4, 0x3c, 0x6e, 0x73, 0x97, 0x54, 0xe6,
				}, chaincfg.MainNetParams.Bech32HRPSegwit,
			),
			f: func() (btcutil.Address, error) {
				scriptHash := []byte{
					0xef, 0x46, 0xd1, 0xaa, 0x78, 0x10, 0x1e, 0x33,
					0x50, 0x60, 0x0a, 0x5d, 0x36, 0x04, 0x5b, 0xa9,
					0x7c, 0x26, 0x70, 0xda, 0xa9, 0x1e, 0x9f, 0x3a,
					0x48, 0xc4, 0x3c, 0x6e, 0x73, 0x97, 0x54, 0xe6,
				}
				return btcutil.NewAddressTaproot(
					scriptHash, &chaincfg.MainNetParams,
				)
			},
			net: &chaincfg.MainNetParams,
		},

		// Invalid bech32m tests. Source:
		// https://github.com/bitcoin/bips/blob/master/bip-0350.mediawiki
		{
			name:  "segwit v1 invalid human-readable part",
			addr:  "tc1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vq5zuyut",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit v1 mainnet bech32 instead of bech32m",
			addr:  "bc1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vqh2y7hd",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit v1 testnet bech32 instead of bech32m",
			addr:  "tb1z0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vqglt7rf",
			valid: false,
			net:   &chaincfg.TestNet3Params,
		},
		{
			name:  "segwit v1 mainnet bech32 instead of bech32m upper case",
			addr:  "BC1S0XLXVLHEMJA6C4DQV22UAPCTQUPFHLXM9H8Z3K2E72Q4K9HCZ7VQ54WELL",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit v0 mainnet bech32m instead of bech32",
			addr:  "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kemeawh",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit v1 testnet bech32 instead of bech32m second test",
			addr:  "tb1q0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vq24jc47",
			valid: false,
			net:   &chaincfg.TestNet3Params,
		},
		{
			name:  "segwit v1 mainnet bech32m invalid character in checksum",
			addr:  "bc1p38j9r5y49hruaue7wxjce0updqjuyyx0kh56v8s25huc6995vvpql3jow4",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit mainnet witness v17",
			addr:  "BC130XLXVLHEMJA6C4DQV22UAPCTQUPFHLXM9H8Z3K2E72Q4K9HCZ7VQ7ZWS8R",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit v1 mainnet bech32m invalid program length (1 byte)",
			addr:  "bc1pw5dgrnzv",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit v1 mainnet bech32m invalid program length (41 bytes)",
			addr:  "bc1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7v8n0nx0muaewav253zgeav",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit v1 testnet bech32m mixed case",
			addr:  "tb1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vq47Zagq",
			valid: false,
			net:   &chaincfg.TestNet3Params,
		},
		{
			name:  "segwit v1 mainnet bech32m zero padding of more than 4 bits",
			addr:  "bc1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7v07qwwzcrf",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit v1 mainnet bech32m non-zero padding in 8-to-5-conversion",
			addr:  "tb1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vpggkg4j",
			valid: false,
			net:   &chaincfg.TestNet3Params,
		},
		{
			name:  "segwit v1 mainnet bech32m empty data section",
			addr:  "bc1gmk9yu",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},

		// Unsupported witness versions (version 0 and 1 only supported at this point)
		{
			name:  "segwit mainnet witness v16",
			addr:  "BC1SW50QA3JX3S",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit mainnet witness v2",
			addr:  "bc1zw508d6qejxtdg4y5r3zarvaryvg6kdaj",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		// Invalid segwit addresses
		{
			name:  "segwit invalid hrp",
			addr:  "tc1qw508d6qejxtdg4y5r3zarvary0c5xw7kg3g4ty",
			valid: false,
			net:   &chaincfg.TestNet3Params,
		},
		{
			name:  "segwit invalid checksum",
			addr:  "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t5",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit invalid witness version",
			addr:  "BC13W508D6QEJXTDG4Y5R3ZARVARY0C5XW7KN40WF2",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit invalid program length",
			addr:  "bc1rw5uspcuh",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit invalid program length",
			addr:  "bc10w508d6qejxtdg4y5r3zarvary0c5xw7kw508d6qejxtdg4y5r3zarvary0c5xw7kw5rljs90",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit invalid program length for witness version 0 (per BIP141)",
			addr:  "BC1QR508D6QEJXTDG4Y5R3ZARVARYV98GJ9P",
			valid: false,
			net:   &chaincfg.MainNetParams,
		},
		{
			name:  "segwit mixed case",
			addr:  "tb1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3q0sL5k7",
			valid: false,
			net:   &chaincfg.TestNet3Params,
		},
		{
			name:  "segwit zero padding of more than 4 bits",
			addr:  "tb1pw508d6qejxtdg4y5r3zarqfsj6c3",
			valid: false,
			net:   &chaincfg.TestNet3Params,
		},
		{
			name:  "segwit non-zero padding in 8-to-5 conversion",
			addr:  "tb1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3pjxtptv",
			valid: false,
			net:   &chaincfg.TestNet3Params,
		},
	}

	if err := chaincfg.Register(&customParams); err != nil {
		panic(err)
	}

	for _, test := range tests {
		// Decode addr and compare error against valid.
		decoded, err := btcutil.DecodeAddress(test.addr, test.net)
		if (err == nil) != test.valid {
			t.Errorf("%v: decoding test failed: %v", test.name, err)
			return
		}

		if err == nil {
			// Ensure the stringer returns the same address as the
			// original.
			if decodedStringer, ok := decoded.(fmt.Stringer); ok {
				addr := test.addr

				// For Segwit addresses the string representation
				// will always be lower case, so in that case we
				// convert the original to lower case first.
				if strings.Contains(test.name, "segwit") {
					addr = strings.ToLower(addr)
				}

				if addr != decodedStringer.String() {
					t.Errorf("%v: String on decoded value does not match expected value: %v != %v",
						test.name, test.addr, decodedStringer.String())
					return
				}
			}

			// Encode again and compare against the original.
			encoded := decoded.EncodeAddress()
			if test.encoded != encoded {
				t.Errorf("%v: decoding and encoding produced different addresses: %v != %v",
					test.name, test.encoded, encoded)
				return
			}

			// Perform type-specific calculations.
			var saddr []byte
			switch d := decoded.(type) {
			case *btcutil.AddressPubKeyHash:
				saddr = btcutil.TstAddressSAddr(encoded)

			case *btcutil.AddressScriptHash:
				saddr = btcutil.TstAddressSAddr(encoded)

			case *btcutil.AddressPubKey:
				// Ignore the error here since the script
				// address is checked below.
				saddr, _ = hex.DecodeString(d.String())
			case *btcutil.AddressWitnessPubKeyHash:
				saddr = btcutil.TstAddressSegwitSAddr(encoded)
			case *btcutil.AddressWitnessScriptHash:
				saddr = btcutil.TstAddressSegwitSAddr(encoded)
			case *btcutil.AddressTaproot:
				saddr = btcutil.TstAddressTaprootSAddr(encoded)
			}

			// Check script address, as well as the Hash160 method for P2PKH and
			// P2SH addresses.
			if !bytes.Equal(saddr, decoded.ScriptAddress()) {
				t.Errorf("%v: script addresses do not match:\n%x != \n%x",
					test.name, saddr, decoded.ScriptAddress())
				return
			}
			switch a := decoded.(type) {
			case *btcutil.AddressPubKeyHash:
				if h := a.Hash160()[:]; !bytes.Equal(saddr, h) {
					t.Errorf("%v: hashes do not match:\n%x != \n%x",
						test.name, saddr, h)
					return
				}

			case *btcutil.AddressScriptHash:
				if h := a.Hash160()[:]; !bytes.Equal(saddr, h) {
					t.Errorf("%v: hashes do not match:\n%x != \n%x",
						test.name, saddr, h)
					return
				}

			case *btcutil.AddressWitnessPubKeyHash:
				if hrp := a.Hrp(); test.net.Bech32HRPSegwit != hrp {
					t.Errorf("%v: hrps do not match:\n%x != \n%x",
						test.name, test.net.Bech32HRPSegwit, hrp)
					return
				}

				expVer := test.result.(*btcutil.AddressWitnessPubKeyHash).WitnessVersion()
				if v := a.WitnessVersion(); v != expVer {
					t.Errorf("%v: witness versions do not match:\n%x != \n%x",
						test.name, expVer, v)
					return
				}

				if p := a.WitnessProgram(); !bytes.Equal(saddr, p) {
					t.Errorf("%v: witness programs do not match:\n%x != \n%x",
						test.name, saddr, p)
					return
				}

			case *btcutil.AddressWitnessScriptHash:
				if hrp := a.Hrp(); test.net.Bech32HRPSegwit != hrp {
					t.Errorf("%v: hrps do not match:\n%x != \n%x",
						test.name, test.net.Bech32HRPSegwit, hrp)
					return
				}

				expVer := test.result.(*btcutil.AddressWitnessScriptHash).WitnessVersion()
				if v := a.WitnessVersion(); v != expVer {
					t.Errorf("%v: witness versions do not match:\n%x != \n%x",
						test.name, expVer, v)
					return
				}

				if p := a.WitnessProgram(); !bytes.Equal(saddr, p) {
					t.Errorf("%v: witness programs do not match:\n%x != \n%x",
						test.name, saddr, p)
					return
				}
			}

			// Ensure the address is for the expected network.
			if !decoded.IsForNet(test.net) {
				t.Errorf("%v: calculated network does not match expected",
					test.name)
				return
			}
		} else {
			// If there is an error, make sure we can print it
			// correctly.
			errStr := err.Error()
			if errStr == "" {
				t.Errorf("%v: error was non-nil but message is"+
					"empty: %v", test.name, err)
			}
		}

		if !test.valid {
			// If address is invalid, but a creation function exists,
			// verify that it returns a nil addr and non-nil error.
			if test.f != nil {
				_, err := test.f()
				if err == nil {
					t.Errorf("%v: address is invalid but creating new address succeeded",
						test.name)
					return
				}
			}
			continue
		}

		// Valid test, compare address created with f against expected result.
		addr, err := test.f()
		if err != nil {
			t.Errorf("%v: address is valid but creating new address failed with error %v",
				test.name, err)
			return
		}

		if !reflect.DeepEqual(addr, test.result) {
			t.Errorf("%v: created address does not match expected result",
				test.name)
			return
		}
	}
}
