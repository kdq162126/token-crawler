package explorerScan

import (
	"context"
	"time"

	"github.com/depocket/multicall-go/call"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

const (
	CALLS_PER_PERIOD = 5
	CALL_DELAY_mS    = time.Duration(1500) * time.Millisecond
)

type ExplorerScanService struct {
	Param *ExplorerConfig
	Log   *zap.Logger
}

func NewExplorerScanService(chain Chain, log *zap.Logger) *ExplorerScanService {
	return &ExplorerScanService{
		Param: newChainExplorerConfig(chain),
		Log:   log,
	}
}

func newChainExplorerConfig(chain Chain) *ExplorerConfig {
	explorerParam := DefaultExplorerConfigs[chain]
	explorerParam.LastBlock = getLastBlock(chain)

	explorerParam.KeyManager.ReadyKey = make(chan string, len(explorerParam.KeyManager.Keys)*CALLS_PER_PERIOD)
	explorerParam.KeyManager.OnKeyReady = func(key string, handle func(string)) {
		handle(key)
		go func() {
			time.Sleep(CALL_DELAY_mS)
			explorerParam.KeyManager.ReadyKey <- key
		}()
	}
	for _, apiKey := range explorerParam.KeyManager.Keys {
		for i := 0; i < CALLS_PER_PERIOD; i++ {
			explorerParam.KeyManager.ReadyKey <- apiKey
		}
	}
	return explorerParam
}

func getLastBlock(chain Chain) int64 {
	client, _ := ethclient.Dial(call.DefaultChainConfigs[call.Chain(chain)].Url)
	header, _ := client.HeaderByNumber(context.Background(), nil)
	return header.Number.Int64()
}
