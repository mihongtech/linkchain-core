package rpc

import (
	"reflect"

	"github.com/mihongtech/linkchain-core/common/http/server"
	"github.com/mihongtech/linkchain-core/common/util/log"
	"github.com/mihongtech/linkchain-core/core/meta"
	"github.com/mihongtech/linkchain-core/node/bcsi"
)

func onGetBlockState(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockIDCmd)
	if !ok {
		log.Error("onGetBlockState Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	blockId := meta.BlockID{}
	if err := blockId.DecodeFromBytes([]byte(c.BlockId)); err != nil {
		log.Error("onGetBlockState cmd decode", err)
		return nil, err
	}

	treeId, err := s.Context.(bcsi.BCSI).GetBlockState(blockId)
	if err != nil {
		log.Error("onGetBlockState GetBlockState return", err)
		return nil, err
	}

	buff, err := treeId.EncodeToBytes()
	if err != nil {
		log.Error("onGetBlockState result encode treeID", err)
		return nil, err
	}
	return string(buff), nil
}

func onUpdateChain(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockCmd)
	if !ok {
		log.Error("onCommit Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	block := meta.Block{}
	if err := block.DecodeFromBytes([]byte(c.Block)); err != nil {
		log.Error("onUpdateChain cmd decode", err)
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).UpdateChain(block)
}

func onProcessBlock(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockCmd)
	if !ok {
		log.Error("onCommit Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	block := meta.Block{}
	if err := block.DecodeFromBytes([]byte(c.Block)); err != nil {
		log.Error("onProcessBlock cmd decode", err)
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).ProcessBlock(block)
}

func onCommit(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockIDCmd)
	if !ok {
		log.Error("onCommit Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	blockId := meta.BlockID{}
	if err := blockId.DecodeFromBytes([]byte(c.BlockId)); err != nil {
		log.Error("onCommit cmd decode", err)
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).Commit(blockId)
}

func onCheckBlock(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockCmd)
	if !ok {
		log.Error("onCheckBlock Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	block := meta.Block{}
	if err := block.DecodeFromBytes([]byte(c.Block)); err != nil {
		log.Error("onCheckBlock cmd decode", err)
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).CheckBlock(block)
}

func onCheckTx(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*TransactionCmd)
	if !ok {
		log.Error("onCheckTx Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	transaction := meta.Transaction{}
	if err := transaction.DecodeFromBytes([]byte(c.Transaction)); err != nil {
		log.Error("onCheckBlock cmd decode", err)
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).CheckTx(transaction)
}

func onFilterTx(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(TransactionsCmd)
	if !ok {
		log.Error("onFilterTx Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	transactions := meta.Transactions{}
	if err := transactions.DecodeFromBytes([]byte(c.Transactions)); err != nil {
		log.Error("onFilterTx cmd decode", err)
		return nil, err
	}

	result := s.Context.(bcsi.BCSI).FilterTx(transactions.Txs)
	resultTxs := meta.NewTransactions(result...)
	buff, err := resultTxs.EncodeToBytes()
	if err != nil {
		log.Error("onFilterTx result encode transactions", err)
		return nil, err
	}

	return string(buff), nil
}
