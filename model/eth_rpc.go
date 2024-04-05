package model

type EthRequest struct {
	ID      int      `json:"id"`
	Version string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

type EthResult struct {
	ID      int    `json:"id"`
	Version string `json:"jsonrpc"`
	Result  string `json:"result"`
}
