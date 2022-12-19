package explorerScan

import (
	"encoding/json"
	"fmt"
	"go-module/models"
	"go-module/utils"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/samber/lo"
)

func (s *ExplorerScanService) GetAllUsersFromTransferEvents(address string, fromBlock int) []string {
	nextBlock := fromBlock
	isFinish := false
	userList := []string{}

	for !isFinish {
		results := []models.Event{}
		results, nextBlock, isFinish = s.GetLogResults(address, nextBlock)
		addresses := lo.Uniq(lo.FlatMap(results, func(item models.Event, index int) []string { return item.Topics[1:] }))
		s.Log.Sugar().Infof("Scanned: %d addresses at %s, nextBlock: %d", len(addresses), time.Now().String(), nextBlock)
		userList = append(userList, addresses...)
	}
	return lo.Uniq(userList)
}

func (s *ExplorerScanService) GetLogResults(address string, fromBlock int) ([]models.Event, int, bool) {
	var requestUrl string
	s.Param.KeyManager.OnKeyReady(<-s.Param.KeyManager.ReadyKey, func(key string) {
		requestUrl = s.buildLogTransferEventsUrl(address, key, fromBlock)
	})

	resp, err := utils.ProxyHttpClient().Get(requestUrl)
	// Retry current block if error detected
	if err != nil {
		return nil, fromBlock, false
	}

	body, err := ioutil.ReadAll(resp.Body)
	// Retry current block if error detected
	if err != nil {
		return nil, fromBlock, false
	}

	var logsResponse models.LogsResponse
	json.Unmarshal(body, &logsResponse)

	// Return start block number for the next call
	if logsResponse.Message == "OK" && logsResponse.Status == 1 {
		nextBlock, _ := strconv.ParseInt(logsResponse.Events[len(logsResponse.Events)-1].BlockNumber[2:], 16, 32)
		return logsResponse.Events, int(nextBlock + 1), false
	}
	// Results length = 0 mean scan to the last block, return block number = -1 to stop this routine
	if logsResponse.Message == "No records found" && logsResponse.Status == 0 {
		return logsResponse.Events, (-1), true
	}
	// Rate limit -> retry
	return logsResponse.Events, fromBlock, false
}

func (s *ExplorerScanService) buildLogTransferEventsUrl(address, apikey string, fromBlock int) string {
	return fmt.Sprintf("%s?module=logs&action=getLogs", s.Param.BaseURL) +
		fmt.Sprintf("&fromBlock=%d", fromBlock) +
		fmt.Sprintf("&toBlock=%s", "lastest") +
		fmt.Sprintf("&address=%s", address) +
		fmt.Sprintf("&apikey=%s", apikey) +
		fmt.Sprintf("&topic0=%s", "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}
