// Copyright (c) 2013 - 2020 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcutil_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/dashevo/dashd-go/btcec/v2"
	. "github.com/dashevo/dashd-go/btcutil"
	"github.com/dashevo/dashd-go/chaincfg"
)

func TestEncodeDecodeWIF(t *testing.T) {
	validEncodeCases := []struct {
		privateKey []byte           // input
		net        *chaincfg.Params // input
		compress   bool             // input
		wif        string           // output
		publicKey  []byte           // output
		name       string           // name of subtest
	}{
		{
			privateKey: []byte{
				0xf1, 0xdf, 0x40, 0x0b, 0x33, 0x56, 0x60, 0x44,
				0x14, 0xb0, 0x75, 0x59, 0x63, 0x48, 0xca, 0x09,
				0x85, 0x62, 0x7e, 0xb2, 0xf3, 0xb3, 0x83, 0x32,
				0xc6, 0x51, 0x08, 0x20, 0xa2, 0x14, 0x3d, 0xc8},
			net:      &chaincfg.MainNetParams,
			compress: false,
			wif:      "7sPPHZ7127xo1NAGtoMHGeoBHWTTBTVKmawchyLMXPEpRbrhbGN",
			publicKey: []byte{
				0x04, 0x38, 0x3b, 0x40, 0x04, 0xe9, 0x7f, 0xaa,
				0x23, 0xb4, 0x76, 0x60, 0x7e, 0x22, 0x99, 0x96,
				0x07, 0x36, 0xfd, 0x96, 0xeb, 0x2a, 0x8b, 0x8a,
				0x93, 0x13, 0x0a, 0xd1, 0xba, 0xbc, 0xe8, 0xeb,
				0xdf, 0x8a, 0x67, 0x91, 0x8a, 0xe8, 0xb6, 0xfa,
				0x3f, 0x00, 0x1b, 0x70, 0xf6, 0xe6, 0x45, 0xc5,
				0x8e, 0xee, 0xfe, 0x60, 0x90, 0xe2, 0xfe, 0xa1,
				0x2d, 0x4e, 0x64, 0x76, 0x84, 0x53, 0x24, 0x22,
				0xd7},
			name: "encodeValidUncompressedMainNetWif",
		},
		{
			privateKey: []byte{
				0x10, 0x3e, 0x05, 0x45, 0x01, 0xfc, 0xb2, 0x86,
				0xc2, 0x02, 0x95, 0xba, 0xa7, 0xa5, 0x4a, 0x1e,
				0x03, 0x65, 0x2e, 0xfa, 0x47, 0x21, 0x94, 0xad,
				0xce, 0xc4, 0x24, 0x8a, 0x3d, 0xf0, 0x84, 0xeb},
			net:      &chaincfg.TestNet3Params,
			compress: true,
			wif:      "cN8GsdKbCifhAFzPsucjWHPvKH7isup3WReCHGgYznAmWiBKyzv3",
			publicKey: []byte{
				0x02, 0xac, 0x92, 0xa9, 0x65, 0x32, 0xe6, 0x70,
				0x0b, 0xb7, 0x8e, 0x3c, 0xb7, 0xe6, 0x85, 0x72,
				0xae, 0x73, 0xc9, 0xb2, 0xf9, 0xfb, 0x9d, 0x33,
				0x47, 0xb8, 0x3a, 0x66, 0xa7, 0x78, 0x85, 0x88,
				0x1d},
			name: "encodeValidCompressedTestNet3Wif",
		},
	}

	for _, validCase := range validEncodeCases {
		validCase := validCase

		t.Run(validCase.name, func(t *testing.T) {
			priv, _ := btcec.PrivKeyFromBytes(validCase.privateKey)
			wif, err := NewWIF(priv, validCase.net, validCase.compress)
			if err != nil {
				t.Fatalf("NewWIF failed: expected no error, got '%v'", err)
			}

			if !wif.IsForNet(validCase.net) {
				t.Fatal("IsForNet failed: got 'false', want 'true'")
			}

			if gotPubKey := wif.SerializePubKey(); !bytes.Equal(gotPubKey, validCase.publicKey) {
				t.Fatalf("SerializePubKey failed: got '%s', want '%s'",
					hex.EncodeToString(gotPubKey), hex.EncodeToString(validCase.publicKey))
			}

			// Test that encoding the WIF structure matches the expected string.
			got := wif.String()
			if got != validCase.wif {
				t.Fatalf("NewWIF failed: want '%s', got '%s'",
					validCase.wif, got)
			}

			// Test that decoding the expected string results in the original WIF
			// structure.
			decodedWif, err := DecodeWIF(got)
			if err != nil {
				t.Fatalf("DecodeWIF failed: expected no error, got '%v'", err)
			}
			if decodedWifString := decodedWif.String(); decodedWifString != validCase.wif {
				t.Fatalf("NewWIF failed: want '%v', got '%v'", validCase.wif, decodedWifString)
			}
		})
	}

	invalidDecodeCases := []struct {
		name string
		wif  string
		err  error
	}{
		{
			name: "decodeInvalidLengthWif",
			wif:  "deadbeef",
			err:  ErrMalformedPrivateKey,
		},
		{
			name: "decodeInvalidCompressMagicWif",
			wif:  "KwDiBf89QgGbjEhKnhXJuH7LrciVrZi3qYjgd9M7rFU73sfZr2ym",
			err:  ErrMalformedPrivateKey,
		},
		{
			name: "decodeInvalidChecksumWif",
			wif:  "5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTj",
			err:  ErrChecksumMismatch,
		},
	}

	for _, invalidCase := range invalidDecodeCases {
		invalidCase := invalidCase

		t.Run(invalidCase.name, func(t *testing.T) {
			decodedWif, err := DecodeWIF(invalidCase.wif)
			if decodedWif != nil {
				t.Fatalf("DecodeWIF: unexpectedly succeeded - got '%v', want '%v'",
					decodedWif, nil)
			}
			if err != invalidCase.err {
				t.Fatalf("DecodeWIF: expected error '%v', got '%v'",
					invalidCase.err, err)
			}
		})
	}

	t.Run("encodeInvalidNetworkWif", func(t *testing.T) {
		privateKey := []byte{
			0x0c, 0x28, 0xfc, 0xa3, 0x86, 0xc7, 0xa2, 0x27,
			0x60, 0x0b, 0x2f, 0xe5, 0x0b, 0x7c, 0xae, 0x11,
			0xec, 0x86, 0xd3, 0xbf, 0x1f, 0xbe, 0x47, 0x1b,
			0xe8, 0x98, 0x27, 0xe1, 0x9d, 0x72, 0xaa, 0x1d}
		priv, _ := btcec.PrivKeyFromBytes(privateKey)

		wif, err := NewWIF(priv, nil, true)

		if wif != nil {
			t.Fatalf("NewWIF: unexpectedly succeeded - got '%v', want '%v'",
				wif, nil)
		}
		if err == nil || err.Error() != "no network" {
			t.Fatalf("NewWIF: expected error 'no network', got '%v'", err)
		}
	})
}
