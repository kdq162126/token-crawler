package models

type TxResponse struct {
	Status  int    `json:"status,string"`
	Message string `json:"message"`
	Txs     []Tx   `json:"result"`
}

type Tx struct {
	BlockNumber int    `json:"blockNumber,string"`
	Hash        string `json:"hash"`
	From        string `json:"from"`
}
