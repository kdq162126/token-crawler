package main

import (
	"encoding/json"
	"go-module/service/thirdparty/explorerScan"
	"go-module/utils"
	"math/big"
	"os"

	"github.com/depocket/multicall-go/call"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

const (
	TokenAddress = "0x7d99eda556388Ad7743A1B658b9C4FC67D7A9d74"
)

func main() {
	config := zap.NewDevelopmentConfig()
	log, _ := config.Build()
	s := explorerScan.NewExplorerScanService(explorerScan.Bsc, log)

	s.Log.Sugar().Info("Start fetch transfer logs from block ", 0)
	userList := s.GetAllUsersFromTransferEvents(TokenAddress, 0)
	s.Log.Sugar().Info("Finish fetch transfer logs, results: ", len(userList), " address")

	userListChunked := utils.ChunkSlice(userList, 2000)

	users := make(map[string]*big.Int)
	holders := make(map[string]*big.Int)

	caller := call.NewContractBuilder().WithChainConfig(call.DefaultChainConfigs[call.Bsc]).AddMethod("balanceOf(address)(uint256)")
	for _, addrs := range userListChunked {
		for _, addr := range addrs {
			caller.AddCall(addr, TokenAddress, "balanceOf", common.HexToAddress(addr))
		}
		_, results, err := caller.Call(nil)
		if err != nil {
			s.Log.Sugar().Error(err)
		}

		for address, balance := range results {
			users[address] = balance[0].(*big.Int)
		}
	}

	for address, balance := range users {
		if utils.HasBalance(balance) {
			holders[address] = balance
		}
	}

	// Save holder list to json file
	data, err := json.MarshalIndent(holders, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}

	f, err := os.Create("holders.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	_, err2 := f.Write(data)

	if err2 != nil {
		log.Fatal(err2.Error())
	}
}
