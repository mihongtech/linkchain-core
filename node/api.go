package node

import (
	"github.com/mihongtech/linkchain-core/core/meta"
	"github.com/mihongtech/linkchain-core/node/net/p2p/discover"
	"github.com/mihongtech/linkchain-core/node/net/p2p/peer"
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

func (c *CoreAPI) Self() *discover.Node {
	return c.node.p2pSvc.Self()
}

func (c *CoreAPI) AddPeer(node *discover.Node) {
	c.node.p2pSvc.AddPeer(node)
}
func (c *CoreAPI) Peers() []*peer.Peer {
	return c.node.p2pSvc.Peers()
}
func (c *CoreAPI) RemovePeer(node *discover.Node) {
	c.node.p2pSvc.RemovePeer(node)
}
