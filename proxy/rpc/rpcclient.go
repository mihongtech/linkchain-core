package rpc

import (
	"encoding/hex"
	"encoding/json"
	"github.com/mihongtech/linkchain-core/common/http/client"
	"github.com/mihongtech/linkchain-core/common/math"
	"github.com/mihongtech/linkchain-core/common/util/log"
	"github.com/mihongtech/linkchain-core/core/meta"
)

type BCSIRPCClient struct {
	cfg *client.Config
}

func NewBCSIRPCClient(cfg *client.Config) *BCSIRPCClient {
	return &BCSIRPCClient{cfg: cfg}
}

func (c *BCSIRPCClient) GetBlockState(id meta.BlockID) (meta.TreeID, error) {
	buff, err := id.EncodeToBytes()
	if err != nil {
		log.Error("BCSIRPCClient", "GetBlockState cmd encode", err)
		return math.Hash{}, err
	}
	cmd := BlockIDCmd{BlockId: hex.EncodeToString(buff)}

	response, err := client.RPC("GetBlockState", cmd, c.cfg)
	if err != nil {
		log.Error("BCSIRPCClient", "GetBlockState rpc connect", err)
		return math.Hash{}, err
	}

	rsp := CommonRSP{}
	if err = json.Unmarshal([]byte(response), &rsp); err != nil {
		log.Error("BCSIRPCClient", "GetBlockState response json Unmarshal", err)
		return math.Hash{}, err
	}

	responseBuff, err := hex.DecodeString(rsp.Data)
	if err != nil {
		log.Error("BCSIRPCClient", "GetBlockState response hex decode", err)
		return math.Hash{}, err
	}

	treeId := meta.TreeID{}
	err = treeId.DecodeFromBytes(responseBuff)
	return treeId, err
}

func (c *BCSIRPCClient) UpdateChain(head meta.Block) error {
	buff, err := head.EncodeToBytes()
	if err != nil {
		log.Error("BCSIRPCClient", "UpdateChain cmd encode", err)
		return err
	}
	cmd := BlockCmd{Block: hex.EncodeToString(buff)}
	_, err = client.RPC("UpdateChain", cmd, c.cfg)
	if err != nil {
		log.Error("BCSIRPCClient", "UpdateChain rpc connect", err)
		return err
	}
	return nil
}

func (c *BCSIRPCClient) ProcessBlock(block meta.Block) error {
	buff, err := block.EncodeToBytes()
	if err != nil {
		log.Error("BCSIRPCClient", "ProcessBlock cmd encode", err)
		return err
	}
	cmd := BlockCmd{Block: hex.EncodeToString(buff)}
	_, err = client.RPC("ProcessBlock", cmd, c.cfg)
	if err != nil {
		log.Error("BCSIRPCClient", "ProcessBlock rpc connect", err)
		return err
	}
	return nil
}

func (c *BCSIRPCClient) Commit(id meta.BlockID) error {
	buff, err := id.EncodeToBytes()
	if err != nil {
		log.Error("BCSIRPCClient", "Commit cmd encode", err)
		return err
	}
	cmd := BlockIDCmd{BlockId: hex.EncodeToString(buff)}
	_, err = client.RPC("Commit", cmd, c.cfg)
	if err != nil {
		log.Error("BCSIRPCClient", "Commit rpc connect", err)
		return err
	}
	return nil
}

func (c *BCSIRPCClient) CheckBlock(block meta.Block) error {
	buff, err := block.EncodeToBytes()
	if err != nil {
		log.Error("BCSIRPCClient", "CheckBlock cmd encode", err)
		return err
	}
	cmd := BlockCmd{Block: hex.EncodeToString(buff)}
	_, err = client.RPC("CheckBlock", cmd, c.cfg)
	if err != nil {
		log.Error("BCSIRPCClient", "CheckBlock rpc connect", err)
		return err
	}
	return nil
}

func (c *BCSIRPCClient) CheckTx(transaction meta.Transaction) error {
	buff, err := transaction.EncodeToBytes()
	if err != nil {
		log.Error("BCSIRPCClient", "CheckTx cmd encode", err)
		return err
	}
	cmd := TransactionCmd{Transaction: hex.EncodeToString(buff)}
	_, err = client.RPC("CheckTx", cmd, c.cfg)
	if err != nil {
		log.Error("BCSIRPCClient", "CheckTx rpc connect", err)
		return err
	}
	return nil
}

func (c *BCSIRPCClient) FilterTx(txs []meta.Transaction) []meta.Transaction {
	filterTxs := make([]meta.Transaction, 0)
	transactions := meta.NewTransactions(txs...)
	buff, err := transactions.EncodeToBytes()
	if err != nil {
		log.Error("BCSIRPCClient", "FilterTx cmd encode", err)
		return filterTxs
	}
	cmd := TransactionsCmd{Transactions: hex.EncodeToString(buff)}

	response, err := client.RPC("FilterTx", cmd, c.cfg)
	if err != nil {
		log.Error("BCSIRPCClient", "FilterTx rpc connect", err)
		return filterTxs
	}

	rsp := CommonRSP{}
	if err = json.Unmarshal([]byte(response), &rsp); err != nil {
		log.Error("BCSIRPCClient", "FilterTx response json Unmarshal", err)
		return filterTxs
	}

	responseBuff, err := hex.DecodeString(rsp.Data)
	if err != nil {
		log.Error("BCSIRPCClient", "FilterTx response hex decode", err)
		return filterTxs
	}

	resultTxs := meta.Transactions{}
	if err = resultTxs.DecodeFromBytes(responseBuff); err != nil {
		log.Error("BCSIRPCClient", "FilterTx cmd decode", err)
		return filterTxs
	}

	return resultTxs.Txs
}
