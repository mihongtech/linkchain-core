package event

import "github.com/mihongtech/linkchain-core/core/meta"

type NewMinedBlockEvent struct {
	Block *meta.Block
}

type TxEvent struct {
	Tx *meta.Transaction
}

type AccountEvent struct {
	IsUpdate bool
}
