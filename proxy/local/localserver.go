package proxy

import (
	"github.com/mihongtech/linkchain-core/core/meta"
	"github.com/mihongtech/linkchain-core/node/bcsi"
)

type LocalServer struct {
	api bcsi.BCSI
}

func NewLocalServer(api bcsi.BCSI) *LocalServer {
	return &LocalServer{api: api}
}

func (s *LocalServer) GetBlockState(id meta.BlockID) (meta.TreeID, error) {
	return s.api.GetBlockState(id)
}

func (s *LocalServer) UpdateChain(head meta.Block) error {
	return s.api.UpdateChain(head)
}

func (s *LocalServer) ProcessBlock(block meta.Block) error {
	return s.api.ProcessBlock(block)
}

func (s *LocalServer) Commit(id meta.BlockID) error {
	return s.api.Commit(id)
}

func (s *LocalServer) CheckBlock(block meta.Block) error {
	return s.api.CheckBlock(block)
}

func (s *LocalServer) CheckTx(transaction meta.Transaction) error {
	return s.api.CheckTx(transaction)
}

func (s *LocalServer) FilterTx(txs []meta.Transaction) []meta.Transaction {
	return s.api.FilterTx(txs)
}
