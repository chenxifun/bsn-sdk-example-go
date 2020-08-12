package models

type ChainData struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type TxData struct {
	TxId        string `json:"txId,omitempty"`
	BlockNumber string `json:"blockNumber,omitempty"`
	BlockHash   string `json:"blockHash,omitempty"`
}
