package rpc

import (
	"reflect"

	"github.com/mihongtech/linkchain-core/common/http/server"
	"github.com/mihongtech/linkchain-core/common/util/log"
	"github.com/mihongtech/linkchain-core/node/bcsi"
)

type BCSIRPCServer struct {
	api       bcsi.BCSI
	rpcServer *server.Server
}

func NewBCSIRPCServer(cfg *server.Config, api bcsi.BCSI) (*BCSIRPCServer, error) {
	//create rpc server
	rpcServer, err := server.NewRPCServer(cfg, api)
	if err != nil {
		log.Error("NewRPCApp", "start rpc server failed", err)
		return nil, err
	}

	//set handler
	rpcServer.SetHandleFunc("GetBlockState", onGetBlockState)
	rpcServer.SetHandleFunc("UpdateChain", onUpdateChain)
	rpcServer.SetHandleFunc("ProcessBlock", onProcessBlock)
	rpcServer.SetHandleFunc("Commit", onCommit)
	rpcServer.SetHandleFunc("CheckBlock", onCheckBlock)
	rpcServer.SetHandleFunc("CheckTx", onCheckTx)
	rpcServer.SetHandleFunc("FilterTx", onFilterTx)
	//set cmd
	rpcServer.SetCmd("GetBlockState", reflect.TypeOf((*BlockIDCmd)(nil)))
	rpcServer.SetCmd("UpdateChain", reflect.TypeOf((*BlockCmd)(nil)))
	rpcServer.SetCmd("ProcessBlock", reflect.TypeOf((*BlockCmd)(nil)))
	rpcServer.SetCmd("Commit", reflect.TypeOf((*BlockIDCmd)(nil)))
	rpcServer.SetCmd("CheckBlock", reflect.TypeOf((*BlockCmd)(nil)))
	rpcServer.SetCmd("CheckTx", reflect.TypeOf((*TransactionCmd)(nil)))
	rpcServer.SetCmd("FilterTx", reflect.TypeOf((*TransactionsCmd)(nil)))
	return &BCSIRPCServer{api: api, rpcServer: rpcServer}, nil
}

func (s *BCSIRPCServer) SetUp(i interface{}) bool {
	return true
}

func (s *BCSIRPCServer) Start() bool {
	s.rpcServer.Start()
	return true
}

func (s *BCSIRPCServer) Stop() bool {
	s.rpcServer.Stop()
	return true
}
