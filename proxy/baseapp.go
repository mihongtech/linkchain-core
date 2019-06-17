package proxy

import (
	"github.com/mihongtech/linkchain-core/common/math"
	"github.com/mihongtech/linkchain-core/core/meta"
)

type BaseApp struct {
}

func (s *BaseApp) GetBlockState(id meta.BlockID) (meta.TreeID, error) {
	return math.Hash{}, nil
}

func (s *BaseApp) UpdateChain(head meta.Block) error {
	return nil
}

func (s *BaseApp) ProcessBlock(block meta.Block) error {
	return nil
}

func (s *BaseApp) Commit(id meta.BlockID) error {
	return nil
}

func (s *BaseApp) CheckBlock(block meta.Block) error {
	return nil
}

func (s *BaseApp) CheckTx(transaction meta.Transaction) error {
	return nil
}

func (s *BaseApp) FilterTx(txs []meta.Transaction) []meta.Transaction {
	return nil
}
