package rpc

import (
	"encoding/hex"
	"reflect"

	"github.com/mihongtech/linkchain-core/common/http/server"
	"github.com/mihongtech/linkchain-core/common/util/log"
	"github.com/mihongtech/linkchain-core/core/meta"
	"github.com/mihongtech/linkchain-core/node/bcsi"
)

func onGetBlockState(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockIDCmd)
	if !ok {
		log.Error("BCSIRPCServer", "onGetBlockState Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	buff, err := hex.DecodeString(c.BlockId)
	if err != nil {
		log.Error("BCSIRPCServer", "onGetBlockState hex cmd decode", err)
		return nil, err
	}

	blockId := meta.BlockID{}
	if err := blockId.DecodeFromBytes(buff); err != nil {
		log.Error("BCSIRPCServer", "onGetBlockState cmd decode", err)
		return nil, err
	}

	treeId, err := s.Context.(bcsi.BCSI).GetBlockState(blockId)
	if err != nil {
		log.Error("BCSIRPCServer", "onGetBlockState GetBlockState return", err)
		return nil, err
	}

	treeBuff, err := treeId.EncodeToBytes()
	if err != nil {
		log.Error("BCSIRPCServer", "onGetBlockState result encode treeID", err)
		return nil, err
	}
	return &CommonRSP{
		Data: hex.EncodeToString(treeBuff),
	}, nil
}

func onUpdateChain(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockCmd)
	if !ok {
		log.Error("BCSIRPCServer", "onUpdateChain Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}

	buff, err := hex.DecodeString(c.Block)
	if err != nil {
		log.Error("BCSIRPCServer", "onUpdateChain hex cmd decode", err)
		return nil, err
	}
	block := meta.Block{}
	if err := block.DecodeFromBytes(buff); err != nil {
		log.Error("BCSIRPCServer", "onUpdateChain cmd decode", err)
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).UpdateChain(block)
}

func onProcessBlock(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockCmd)
	if !ok {
		log.Error("BCSIRPCServer", "onProcessBlock Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	buff, err := hex.DecodeString(c.Block)
	if err != nil {
		log.Error("BCSIRPCServer", "onProcessBlock hex cmd decode", err)
		return nil, err
	}
	block := meta.Block{}
	if err := block.DecodeFromBytes(buff); err != nil {
		log.Error("BCSIRPCServer", "onProcessBlock cmd decode", err)
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).ProcessBlock(block)
}

func onCommit(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockIDCmd)
	if !ok {
		log.Error("BCSIRPCServer", "onCommit Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	buff, err := hex.DecodeString(c.BlockId)
	if err != nil {
		log.Error("BCSIRPCServer", "onCommit hex cmd decode", err)
		return nil, err
	}
	blockId := meta.BlockID{}
	if err := blockId.DecodeFromBytes(buff); err != nil {
		log.Error("BCSIRPCServer", "onCommit cmd decode", err)
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).Commit(blockId)
}

func onCheckBlock(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*BlockCmd)
	if !ok {
		log.Error("BCSIRPCServer", "onCheckBlock Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	buff, err := hex.DecodeString(c.Block)
	if err != nil {
		log.Error("BCSIRPCServer", "onCheckBlock hex cmd decode", err)
		return nil, err
	}
	block := meta.Block{}
	if err := block.DecodeFromBytes(buff); err != nil {
		log.Error("BCSIRPCServer", "onCheckBlock cmd decode", err)
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).CheckBlock(block)
}

func onCheckTx(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*TransactionCmd)
	if !ok {
		log.Error("BCSIRPCServer", "onCheckTx Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	buff, err := hex.DecodeString(c.Transaction)
	if err != nil {
		log.Error("BCSIRPCServer", "onCheckTx hex cmd decode", err)
		return nil, err
	}
	transaction := meta.Transaction{}
	if err := transaction.DecodeFromBytes(buff); err != nil {
		log.Error("BCSIRPCServer", "onCheckTx cmd decode", err)
		return nil, err
	}
	return nil, s.Context.(bcsi.BCSI).CheckTx(transaction)
}

func onFilterTx(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*TransactionsCmd)
	if !ok {
		log.Error("BCSIRPCServer", "onFilterTx Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	buff, err := hex.DecodeString(c.Transactions)
	if err != nil {
		log.Error("BCSIRPCServer", "onFilterTx hex cmd decode", err)
		return nil, err
	}
	transactions := meta.Transactions{}
	if err := transactions.DecodeFromBytes(buff); err != nil {
		log.Error("BCSIRPCServer", "onFilterTx cmd decode", err)
		return nil, err
	}

	result := s.Context.(bcsi.BCSI).FilterTx(transactions.Txs)
	resultTxs := meta.NewTransactions(result...)
	resultBuff, err := resultTxs.EncodeToBytes()
	if err != nil {
		log.Error("BCSIRPCServer", "onFilterTx result encode transactions", err)
		return nil, err
	}
	return &CommonRSP{
		Data: hex.EncodeToString(resultBuff),
	}, nil
}
