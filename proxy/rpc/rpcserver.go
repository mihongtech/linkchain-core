package rpc

import (
	"github.com/mihongtech/linkchain-core/common/http/server"
	"github.com/mihongtech/linkchain-core/common/util/log"
	"github.com/mihongtech/linkchain-core/core/meta"
	"github.com/mihongtech/linkchain-core/node/bcsi"
	"reflect"
)

type RPCApp struct {
	api       bcsi.BCSI
	rpcServer *server.Server
}

func NewRPCApp(cfg *server.Config, api bcsi.BCSI) (*RPCApp, error) {
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
	rpcServer.SetCmd("GetBlockState", reflect.TypeOf((*meta.BlockID)(nil)))
	rpcServer.SetCmd("UpdateChain", reflect.TypeOf((*meta.Block)(nil)))
	rpcServer.SetCmd("ProcessBlock", reflect.TypeOf((*meta.Block)(nil)))
	rpcServer.SetCmd("Commit", reflect.TypeOf((*meta.BlockID)(nil)))
	rpcServer.SetCmd("CheckBlock", reflect.TypeOf((*meta.Block)(nil)))
	rpcServer.SetCmd("CheckTx", reflect.TypeOf((*meta.Transaction)(nil)))
	rpcServer.SetCmd("FilterTx", reflect.TypeOf(([]meta.Transaction)(nil)))
	return &RPCApp{api: api, rpcServer: rpcServer}, nil
}
