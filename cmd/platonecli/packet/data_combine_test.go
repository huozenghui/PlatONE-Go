package packet

import (
	"testing"

	"github.com/PlatONEnetwork/PlatONE-Go/core/types"

	"github.com/PlatONEnetwork/PlatONE-Go/cmd/platonecli/test"
	utl "github.com/PlatONEnetwork/PlatONE-Go/cmd/platonecli/utils"
	"github.com/PlatONEnetwork/PlatONE-Go/common"
)

const (
	CODE_PATH = "../test/test_case/wasm/contracta.wasm"
	ABI_PATH  = "../test/test_case/wasm/contracta.cpp.abi.json"
	//TEST_ACCOUNT = "0x60ceca9c1290ee56b98d4e160ef0453f7c40d219"

)

func TestDeploy(t *testing.T) {

	codeBytes, _ := utl.ParseFileToBytes(CODE_PATH)
	abiBytes, _ := utl.ParseFileToBytes(ABI_PATH)

	call := NewDeployCall(codeBytes, abiBytes, "wasm", types.CreateTxType)

	data, _, isWrite, _ := call.CombineData()
	from := common.HexToAddress(test.TEST_ACCOUNT)
	tx := NewTxParams(from, nil, "", "", "", data)
	params, action := tx.SendMode(isWrite, "")

	server := test.MockServer("rpc")
	utl.SetHttpUrl(server.URL)

	//fmt.Printf("action is %s, params are %+v\n", action, params)

	response, err := utl.RpcCalls(action, params)

	if err != nil {
		t.Logf("FAILED, error is %s\n", err.Error())
	} else {
		t.Logf("SUCCESS, transaction hash is %s\n", response.(string))
	}
}
