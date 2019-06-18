package rpc

type BlockIDCmd struct {
	BlockId string `json:"blockId"`
}

type BlockCmd struct {
	Block string `json:"block"`
}

type TransactionCmd struct {
	Transaction string `json:"transaction"`
}

type TransactionsCmd struct {
	Transactions string `json:"transactions"`
}

type CommonRSP struct {
	Data string `json:"data"`
}
