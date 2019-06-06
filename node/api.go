package node

import (
	"github.com/mihongtech/linkchain-core/core/meta"
	"math/big"
)

type CoreAPI struct {
	node *Node
}

func NewPublicCoreAPI(node *Node) *CoreAPI {
	return &CoreAPI{node: node}
}

func (c *CoreAPI) GetBestBlock() *meta.Block {
	return c.node.blockchain.GetBestBlock()
}

func (c *CoreAPI) GetBlockNumber(id meta.BlockID) uint64 {
	return c.node.blockchain.GetBlockNumber(id)
}

func (c *CoreAPI) GetBlockByID(id meta.BlockID) (*meta.Block, error) {
	return c.node.blockchain.GetBlockByID(id)
}

func (c *CoreAPI) GetBlockByHeight(height uint32) (*meta.Block, error) {
	return c.node.blockchain.GetBlockByHeight(height)
}

func (c *CoreAPI) GetChainID() *big.Int {
	return c.node.blockchain.GetChainID()
}
