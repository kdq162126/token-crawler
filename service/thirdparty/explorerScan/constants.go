package explorerScan

type Chain string

type ApiKeyConfig struct {
	Key      []string
	ReadyKey chan string
}

type KeyManager struct {
	Keys       []string
	ReadyKey   chan string
	OnKeyReady func(string, func(string))
}

type ExplorerConfig struct {
	BaseURL    string
	KeyManager KeyManager
	LastBlock  int64
}

const (
	Arbitrum  Chain = "arbitrum"
	Aurora          = "aurora"
	Avalanche       = "avalanche"
	Bsc             = "bsc"
	Ethereum        = "ethereum"
	Fantom          = "fantom"
	Moonbeam        = "moonbeam"
	Moonriver       = "moonriver"
	Celo            = "celo"
)

var DefaultExplorerConfigs = map[Chain]*ExplorerConfig{
	Bsc: {
		BaseURL: "https://api.bscscan.com/api",
		KeyManager: KeyManager{
			Keys: []string{
				"611YG1AGRRS9R3X6CQJ982YY612M4Q4P2U",
				"9AEVUX9T2423AZ3M7TYE88CFRTWMPB1C1F",
				"BPCP9JXTBP5QIAAMZWRRK8GI8KSAA9TUMG",
				"4EETMQAT41NQRG2WVQ36AWG3PI9W2FK53D",
				"QD4H3C7HE2IBQVVVUR4AUWV1UJ4BGYGD9B",
				"CY7W1GA9PXJRGTTHBQ3HBH262GCYSQG7G5",
				"ZP9IZZKYQBY4KHV2HGW9NZKZ63JRUD7JK3",
				"3NUUQGW4SKUQFJAZKFT3413ACVFW5B5EIH",
				"F98NREDI815GNZFWTRRIZ3JIHU11ZFGE7Y",
			},
		},
	},
	Ethereum: {
		BaseURL: "https://api.etherscan.io/api",
		KeyManager: KeyManager{
			Keys: []string{
				"Z8JD5SFHQQ3BENRPXVZ5MTQIN3UWUPI652",
				"9WTGNISHG9FHNX91D3NNF99W3T9HB3HD22",
				"8C4X3NMBC7ZK7RWZMW8692PJZEPGQ8P4TU",
				"9W89PQ7UXWTFWCDGIFK1ER4FN4WUWF8ZIU",
				"ABYKY8BXM2NEQN7B1JVKKCGM79FZUPYYPD",
				"9YFWSQCQ5KKQNVD5VHU4VXHTPHADVFNM2A",
				"QSIRTI74W1XZJR9EMRYCXIDPN4CP9KW6I8",
				"533J2XBVHIBAW675JD2FI7EW3P6HWA5RFB",
				"K6C2P1212WDD97H41H54PYP7GDRT28DJUI",
			},
		},
	},
}
