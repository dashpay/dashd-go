package rpcclient

import (
	"encoding/json"
	"os/exec"
	"reflect"
	"testing"

	"github.com/dashevo/dashd-go/btcjson"
)

func TestGetBlockCount(t *testing.T) {
	client, err := New(connCfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Shutdown()

	n := int64(1700)
	client.httpClient.Transport = mockRoundTripperFunc(&n)
	blockCount, err := client.GetBlockCount()
	if err != nil {
		t.Fatal(err)
	}

	if blockCount != n {
		t.Fatal("block count did not match")
	}
}

func TestMasternodeStatus(t *testing.T) {
	client, err := New(connCfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Shutdown()

	client.httpClient.Transport = mockRoundTripperFunc(&btcjson.MasternodeStatusResult{})
	result, err := client.MasternodeStatus()
	if err != nil {
		t.Fatal(err)
	}

	cli := &btcjson.MasternodeStatusResult{}
	compareWithCliCommand(t, result, cli, "masternode", "status")
}

func TestMasternodeCount(t *testing.T) {
	client, err := New(connCfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Shutdown()

	client.httpClient.Transport = mockRoundTripperFunc(&btcjson.MasternodeCountResult{})
	result, err := client.MasternodeCount()
	if err != nil {
		t.Fatal(err)
	}

	cli := &btcjson.MasternodeCountResult{}
	compareWithCliCommand(t, result, cli, "masternode", "count")
}

func TestMasternodeCurrent(t *testing.T) {
	client, err := New(connCfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Shutdown()

	client.httpClient.Transport = mockRoundTripperFunc(&btcjson.MasternodeResult{})
	result, err := client.MasternodeCurrent()
	if err != nil {
		t.Fatal(err)
	}

	cli := &btcjson.MasternodeResult{}
	compareWithCliCommand(t, result, cli, "masternode", "current")
}

func TestMasternodeOutputs(t *testing.T) {
	t.Skip("error: -32601: Method not found (wallet method is disabled because no wallet is loaded)")

	client, err := New(connCfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Shutdown()

	var resp map[string]string
	client.httpClient.Transport = mockRoundTripperFunc(&resp)
	result, err := client.MasternodeOutputs()
	if err != nil {
		t.Fatal(err)
	}

	cli := &map[string]string{}
	compareWithCliCommand(t, &result, cli, "masternode", "outputs")
}

func TestMasternodeWinner(t *testing.T) {
	client, err := New(connCfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Shutdown()

	client.httpClient.Transport = mockRoundTripperFunc(&btcjson.MasternodeResult{})
	result, err := client.MasternodeWinner()
	if err != nil {
		t.Fatal(err)
	}

	cli := &btcjson.MasternodeResult{}
	compareWithCliCommand(t, result, cli, "masternode", "winner")
}

func TestMasternodeWinners(t *testing.T) {
	client, err := New(connCfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Shutdown()

	testCases := []struct {
		name   string
		count  int
		filter string
	}{
		{
			name:   "no params",
			count:  0,
			filter: "",
		},
		{
			name:   "just count",
			count:  20,
			filter: "",
		},
		{
			name:   "count and filter",
			count:  30,
			filter: "yP8A3",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var resp map[string]string
			client.httpClient.Transport = mockRoundTripperFunc(&resp)
			result, err := client.MasternodeWinners(tc.count, tc.filter)
			if err != nil {
				t.Fatal(err)
			}
			cli := &map[string]string{}
			compareWithCliCommand(t, &result, cli, "masternode", "winners")
		})
	}
}

func TestMasternodeList(t *testing.T) {
	client, err := New(connCfg, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Shutdown()

	var resp map[string]string
	client.httpClient.Transport = mockRoundTripperFunc(&resp)
	result, err := client.MasternodeList("addr", "")
	if err != nil {
		t.Fatal(err)
	}

	cli := &map[string]string{}
	compareWithCliCommand(t, result, cli, "masternodelist", "addr")

	resultJSON, err := client.MasternodeListJSON("")
	if err != nil {
		t.Fatal(err)
	}

	cliJSON := &map[string]btcjson.MasternodelistResultJSON{}
	compareWithCliCommand(t, &resultJSON, cliJSON, "masternodelist", "json")
}

func compareWithCliCommand(t *testing.T, rpc, cli interface{}, cmds ...string) {
	modifyThenCompareWithCliCommand(t, nil, rpc, cli, cmds...)
}

func modifyThenCompareWithCliCommand(t *testing.T, modify func(interface{}), rpc, cli interface{}, cmds ...string) {
	cmd := append([]string{"-testnet"}, cmds...)
	out, err := exec.Command("dash-cli", cmd...).Output()
	if err != nil {
		t.Fatal("Could not run dash-cli command", err)
	}

	if err := json.Unmarshal(out, &cli); err != nil {
		t.Log(string(out))
		t.Fatal("Could not marshal dash-cli output", err)
	}

	if modify != nil {
		modify(rpc)
		modify(cli)
	}

	if !reflect.DeepEqual(rpc, cli) {
		t.Error("rpc result did not match cli")
		t.Logf("rpc result: %#v", rpc)
		t.Logf("cli result: %#v", cli)
		t.Logf("cli string: %s", string(out))
	}
}
