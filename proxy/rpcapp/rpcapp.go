package rpcapp

import (
	"github.com/mihongtech/linkchain-core/common/math"
	"github.com/mihongtech/linkchain-core/core/meta"
	"github.com/mihongtech/linkchain-core/proxy"
)

type RPCApp struct {
	proxy.BaseApp
}

func (s *RPCApp) GetBlockState(id meta.BlockID) (meta.TreeID, error) {
	return math.Hash{}, nil
}

func (s *RPCApp) UpdateChain(head *meta.Block) error {
	return nil
}

func (s *RPCApp) ProcessBlock(block *meta.Block) error {
	return nil
}

func (s *RPCApp) Commit(id meta.BlockID) error {
	return nil
}

func (s *RPCApp) CheckBlock(block *meta.Block) error {
	return nil
}

func (s *RPCApp) CheckTx(transaction meta.Transaction) error {
	return nil
}

func (s *RPCApp) FilterTx(txs []meta.Transaction) []meta.Transaction {
	return nil
}
