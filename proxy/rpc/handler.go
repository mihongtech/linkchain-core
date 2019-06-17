package rpc

import (
	"github.com/mihongtech/linkchain-core/common/http/server"
	"github.com/mihongtech/linkchain-core/common/util/log"
	"github.com/mihongtech/linkchain-core/core/meta"
	"github.com/mihongtech/linkchain-core/node/bcsi"
	"reflect"
)

func onGetBlockState(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockIDCmd)
	if !ok {
		log.Error("Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	blockId := meta.BlockID{}
	if err := blockId.DecodeFromBytes([]byte(c.BlockId)); err != nil {
		return nil, err
	}
	return s.Context.(bcsi.BCSI).GetBlockState(blockId)
}

func onUpdateChain(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockCmd)
	if !ok {
		log.Error("Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	block := meta.Block{}
	if err := block.DecodeFromBytes([]byte(c.Block)); err != nil {
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).UpdateChain(block)
}

func onProcessBlock(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockCmd)
	if !ok {
		log.Error("Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	block := meta.Block{}
	if err := block.DecodeFromBytes([]byte(c.Block)); err != nil {
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).ProcessBlock(block)
}

func onCommit(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockIDCmd)
	if !ok {
		log.Error("Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	blockId := meta.BlockID{}
	if err := blockId.DecodeFromBytes([]byte(c.BlockId)); err != nil {
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).Commit(blockId)
}

func onCheckBlock(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockCmd)
	if !ok {
		log.Error("Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	block := meta.Block{}
	if err := block.DecodeFromBytes([]byte(c.Block)); err != nil {
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).CheckBlock(block)
}

func onCheckTx(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*TransactionCmd)
	if !ok {
		log.Error("Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	transaction := meta.Transaction{}
	if err := transaction.DecodeFromBytes([]byte(c.Transaction)); err != nil {
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).CheckTx(transaction)
}

func onFilterTx(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(TransactionsCmd)
	if !ok {
		log.Error("Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	txs := make([]meta.Transaction, 0)
	for i := range c.Transactions {
		transaction := meta.Transaction{}
		if err := transaction.DecodeFromBytes([]byte(c.Transactions[i])); err != nil {
			return nil, err
		}
		txs = append(txs, transaction)
	}

	return s.Context.(bcsi.BCSI).FilterTx(txs), nil
}
