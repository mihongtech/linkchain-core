package proxy

import (
	"github.com/mihongtech/linkchain-core/core/meta"
)

type LocalClient struct {
	server *LocalServer
}

func NewLocalClient(server *LocalServer) *LocalClient {
	return &LocalClient{server: server}
}

func (s *LocalClient) GetBlockState(id meta.BlockID) (meta.TreeID, error) {
	return s.server.GetBlockState(id)
}

func (s *LocalClient) UpdateChain(head meta.Block) error {
	return s.server.UpdateChain(head)
}

func (s *LocalClient) ProcessBlock(block meta.Block) error {
	return s.server.ProcessBlock(block)
}

func (s *LocalClient) Commit(id meta.BlockID) error {
	return s.server.Commit(id)
}

func (s *LocalClient) CheckBlock(block meta.Block) error {
	return s.server.CheckBlock(block)
}

func (s *LocalClient) CheckTx(transaction meta.Transaction) error {
	return s.server.CheckTx(transaction)
}

func (s *LocalClient) FilterTx(txs []meta.Transaction) []meta.Transaction {
	return s.server.FilterTx(txs)
}
