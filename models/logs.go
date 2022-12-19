package models

type LogsResponse struct {
	Status  int     `json:"status,string"`
	Message string  `json:"message"`
	Events  []Event `json:"result"`
}

type Event struct {
	Address     string   `json:"address"`
	Topics      []string `json:"topics"`
	BlockNumber string   `json:"blockNumber"`
}
