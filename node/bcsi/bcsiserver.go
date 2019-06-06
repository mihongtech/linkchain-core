package bcsi

import "github.com/mihongtech/linkchain-core/core/meta"

type BCSIServer struct {
}

func GetBlockState(id meta.BlockID) meta.TreeID {
	return id
}

func UpdateChain(head *meta.Block) error {
	return nil
}

func ProcessBlock(block *meta.Block) error {
	return nil
}

func Commit(id meta.BlockID) error {
	return nil
}

func CheckBlock(block *meta.Block) error {
	return nil
}

func CheckTx(transaction meta.Transaction) error {
	return nil
}

func FilterTx(txs []meta.Transaction) []meta.Transaction {
	return nil
}
