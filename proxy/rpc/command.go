package rpc

type BlockIDCmd struct {
	BlockId string `json:"blockId"`
}

type BlockCmd struct {
	Block string `json:"block"`
}

type TransactionCmd struct {
	Transaction string `json:"Transaction"`
}

type TransactionsCmd struct {
	Transactions string `json:"Transactions"`
}
