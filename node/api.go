package node

import (
	"github.com/mihongtech/linkchain-core/common/math"
	"github.com/mihongtech/linkchain-core/core/meta"
	"github.com/mihongtech/linkchain-core/node/chain/storage"
	"github.com/mihongtech/linkchain-core/node/config"
	"github.com/mihongtech/linkchain-core/node/event"
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

/**chainReader inteface**/

func (c *CoreAPI) HasBlock(hash meta.BlockID) bool {
	return c.node.blockchain.HasBlock(hash)
}

func (c *CoreAPI) GetHeader(hash math.Hash, height uint64) *meta.BlockHeader {
	return c.node.blockchain.GetHeader(hash, height)
}

func (c *CoreAPI) GetChainConfig() *config.ChainConfig {
	return c.node.blockchain.GetChainConfig()
}

func (c *CoreAPI) GetBestBlock() *meta.Block {
	return c.node.blockchain.GetBestBlock()
}

func (c *CoreAPI) GetBlockNumber(id meta.BlockID) uint64 {
	return c.node.blockchain.GetBlockNumber(id)
}

func (c *CoreAPI) GetBlockByID(hash meta.BlockID) (*meta.Block, error) {
	return c.node.blockchain.GetBlockByID(hash)
}

func (c *CoreAPI) GetBlockByHeight(height uint32) (*meta.Block, error) {
	return c.node.blockchain.GetBlockByHeight(height)
}

func (c *CoreAPI) GetChainID() *big.Int {
	return c.node.blockchain.GetChainID()
}

/**P2PNet inteface**/
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

/**Tx inteface**/
func (c *CoreAPI) ProcessTx(tx *meta.Transaction) error {
	if err := c.node.txPool.ProcessTx(tx); err != nil {
		return err
	}
	c.node.newTxEvent.Send(event.TxEvent{tx})
	return nil
}

func (c *CoreAPI) GetTXByID(id meta.TxID) (*meta.Transaction, meta.BlockID, uint64, uint64) {
	tx, blockId, number, index := storage.GetTransaction(c.node.db, id)
	if tx == nil {
		return nil, math.Hash{}, 0, 0
	} else {
		return tx, blockId, number, index
	}
}
