// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2021 Dash Core Group
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcjson_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/dashevo/dashd-go/btcjson"
)

func pString(s string) *string                       { return &s }
func pBool(b bool) *bool                             { return &b }
func pLLMQType(l btcjson.LLMQType) *btcjson.LLMQType { return &l }

// TestDashEvoCmds tests all of the dash evo commands marshal and unmarshal
// into valid results include handling of optional fields being omitted in the
// marshalled command, while optional fields with defaults have the default
// assigned on unmarshalled commands.
func TestDashEvoCmds(t *testing.T) {
	t.Parallel()

	testID := 1
	tests := []struct {
		name         string
		newCmd       func() (interface{}, error)
		staticCmd    func() interface{}
		marshalled   string
		unmarshalled interface{}
	}{
		{
			name: "quorum sign",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd("quorum", "sign", btcjson.LLMQType_100_67,
					"0067c4fd779a195a95b267e263c631f71f83f8d5e6191091289d114012b373a1",
					"ce490ca26cad6f1749ff9b977fe0fe4ece4391166f69be75c4619bc94b184dbc",
					"6f1018f54507606069303fd16257434073c6f374729b0090bb9dbbe629241236",
					false)
			},
			staticCmd: func() interface{} {
				return btcjson.NewQuorumSignCmd(btcjson.LLMQType_100_67,
					"0067c4fd779a195a95b267e263c631f71f83f8d5e6191091289d114012b373a1",
					"ce490ca26cad6f1749ff9b977fe0fe4ece4391166f69be75c4619bc94b184dbc",
					"6f1018f54507606069303fd16257434073c6f374729b0090bb9dbbe629241236",
					false)
			},
			marshalled: `{"jsonrpc":"1.0","method":"quorum","params":["sign",4,"0067c4fd779a195a95b267e263c631f71f83f8d5e6191091289d114012b373a1","ce490ca26cad6f1749ff9b977fe0fe4ece4391166f69be75c4619bc94b184dbc","6f1018f54507606069303fd16257434073c6f374729b0090bb9dbbe629241236",false],"id":1}`,
			unmarshalled: &btcjson.QuorumCmd{
				SubCmd:      "sign",
				LLMQType:    pLLMQType(btcjson.LLMQType_100_67),
				RequestID:   pString("0067c4fd779a195a95b267e263c631f71f83f8d5e6191091289d114012b373a1"),
				MessageHash: pString("ce490ca26cad6f1749ff9b977fe0fe4ece4391166f69be75c4619bc94b184dbc"),
				QuorumHash:  pString("6f1018f54507606069303fd16257434073c6f374729b0090bb9dbbe629241236"),
				Submit:      pBool(false),
			},
		},
		{
			name: "quorum info",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd("quorum", "info", btcjson.LLMQType_100_67,
					"0067c4fd779a195a95b267e263c631f71f83f8d5e6191091289d114012b373a1",
					false)
			},
			staticCmd: func() interface{} {
				return btcjson.NewQuorumInfoCmd(btcjson.LLMQType_100_67,
					"0067c4fd779a195a95b267e263c631f71f83f8d5e6191091289d114012b373a1",
					false)
			},
			marshalled: `{"jsonrpc":"1.0","method":"quorum","params":["info",4,"0067c4fd779a195a95b267e263c631f71f83f8d5e6191091289d114012b373a1",false],"id":1}`,
			unmarshalled: &btcjson.QuorumCmd{
				SubCmd:         "info",
				LLMQType:       pLLMQType(btcjson.LLMQType_100_67),
				QuorumHash:     pString("0067c4fd779a195a95b267e263c631f71f83f8d5e6191091289d114012b373a1"),
				IncludeSkShare: pBool(false),
			},
		},
		{
			name: "quorum",
			newCmd: func() (interface{}, error) {
				return btcjson.NewCmd("quorum",
					"verify",
					btcjson.LLMQType_100_67,
					"0067c4fd779a195a95b267e263c631f71f83f8d5e6191091289d114012b373a1",
					"ce490ca26cad6f1749ff9b977fe0fe4ece4391166f69be75c4619bc94b184dbc",
					"6f1018f54507606069303fd16257434073c6f374729b0090bb9dbbe629241236",
					"5f1018f54507606069303fd16257434073c6f374729b0090bb9dbbe629241235")
			},
			staticCmd: func() interface{} {
				return btcjson.NewQuorumVerifyCmd(btcjson.LLMQType_100_67,
					"0067c4fd779a195a95b267e263c631f71f83f8d5e6191091289d114012b373a1",
					"ce490ca26cad6f1749ff9b977fe0fe4ece4391166f69be75c4619bc94b184dbc",
					"5f1018f54507606069303fd16257434073c6f374729b0090bb9dbbe629241235",
					"6f1018f54507606069303fd16257434073c6f374729b0090bb9dbbe629241236")
			},
			marshalled: `{"jsonrpc":"1.0","method":"quorum","params":["verify",4,"0067c4fd779a195a95b267e263c631f71f83f8d5e6191091289d114012b373a1","ce490ca26cad6f1749ff9b977fe0fe4ece4391166f69be75c4619bc94b184dbc","6f1018f54507606069303fd16257434073c6f374729b0090bb9dbbe629241236","5f1018f54507606069303fd16257434073c6f374729b0090bb9dbbe629241235"],"id":1}`,
			unmarshalled: &btcjson.QuorumCmd{
				SubCmd:      "verify",
				LLMQType:    pLLMQType(btcjson.LLMQType_100_67),
				RequestID:   pString("0067c4fd779a195a95b267e263c631f71f83f8d5e6191091289d114012b373a1"),
				MessageHash: pString("ce490ca26cad6f1749ff9b977fe0fe4ece4391166f69be75c4619bc94b184dbc"),
				Signature:   pString("5f1018f54507606069303fd16257434073c6f374729b0090bb9dbbe629241235"),
				QuorumHash:  pString("6f1018f54507606069303fd16257434073c6f374729b0090bb9dbbe629241236"),
			},
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		// Marshal the command as created by the new static command
		// creation function.
		marshalled, err := btcjson.MarshalCmd(btcjson.RpcVersion1, testID, test.staticCmd())
		if err != nil {
			t.Errorf("MarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !bytes.Equal(marshalled, []byte(test.marshalled)) {
			t.Errorf("Test #%d (%s) unexpected marshalled data - "+
				"got %s, want %s", i, test.name, marshalled,
				test.marshalled)
			t.Errorf("\n%s\n%s", marshalled, test.marshalled)
			continue
		}

		// Ensure the command is created without error via the generic
		// new command creation function.
		cmd, err := test.newCmd()
		if err != nil {
			t.Errorf("Test #%d (%s) unexpected NewCmd error: %v ",
				i, test.name, err)
		}

		// Marshal the command as created by the generic new command
		// creation function.
		marshalled, err = btcjson.MarshalCmd(btcjson.RpcVersion1, testID, cmd)
		if err != nil {
			t.Errorf("MarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !bytes.Equal(marshalled, []byte(test.marshalled)) {
			t.Errorf("Test #%d (%s) unexpected marshalled data - "+
				"got %s, want %s", i, test.name, marshalled,
				test.marshalled)
			continue
		}

		var request btcjson.Request
		if err := json.Unmarshal(marshalled, &request); err != nil {
			t.Errorf("Test #%d (%s) unexpected error while "+
				"unmarshalling JSON-RPC request: %v", i,
				test.name, err)
			continue
		}

		cmd, err = btcjson.UnmarshalCmd(&request)
		if err != nil {
			t.Errorf("UnmarshalCmd #%d (%s) unexpected error: %v", i,
				test.name, err)
			continue
		}

		if !reflect.DeepEqual(cmd, test.unmarshalled) {
			t.Errorf("Test #%d (%s) unexpected unmarshalled command "+
				"- got %s, want %s", i, test.name,
				fmt.Sprintf("(%T) %+[1]v", cmd),
				fmt.Sprintf("(%T) %+[1]v\n", test.unmarshalled))
			continue
		}
	}
}

func TestLLMQTypeValidate(t *testing.T) {
	testCases := []struct {
		llmqType  btcjson.LLMQType
		expectErr bool
	}{
		{-1, true},
		{0, true},
		{1, false},
		{2, false},
		{5, false},
		{6, true},
		{99, true},
		{100, false},
		{105, false},
		{106, false},
		{107, false},
		{108, true},
	}

	for _, tc := range testCases {
		t.Run(strconv.Itoa(int(tc.llmqType)), func(t *testing.T) {
			err := tc.llmqType.Validate()
			if (err != nil) != tc.expectErr {
				t.Errorf("LLMQ Type %d, expected error %v, got %s", tc.llmqType, tc.expectErr, err)
			}
		})
	}

}

func TestLLMQTypeString(t *testing.T) {
	testCases := []struct {
		llmqType btcjson.LLMQType
		name     string
	}{
		{0, ""},
		{btcjson.LLMQType_400_60, "llmq_400_60"},
		{btcjson.LLMQType_TEST, "llmq_test"},
		{btcjson.LLMQType_5_60, "llmq_test"}, // exception
		{999999, ""},
	}
	for _, tc := range testCases {
		t.Run(strconv.Itoa(int(tc.llmqType)), func(t *testing.T) {
			gotName := tc.llmqType.Name()
			if gotName != tc.name {
				t.Errorf("invalid llmq type name, got: %s, expected: %s", tc.llmqType.Name(), tc.name)
			}

			gotType := btcjson.GetLLMQType(tc.name)
			if (gotName != "" && tc.llmqType != gotType) || (gotName == "" && gotType != 0) {
				t.Errorf("invalid llmq type, got: %d, expected: %d", gotType, tc.llmqType)
			}
		})
	}
}
