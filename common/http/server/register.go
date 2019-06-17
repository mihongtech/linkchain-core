package server

import (
	"reflect"
)

type commandHandler func(*Server, interface{}, <-chan struct{}) (interface{}, error)

//handler pool
var handlerPool = map[string]commandHandler{
	//"getBlockChainInfo": getBlockChainInfo,
}

var cmdPool = map[string]reflect.Type{
	//"version":    reflect.TypeOf((*rpcobject.VersionCmd)(nil)),
}

func SetHandleFunc(method string, handler commandHandler) {
	handlerPool[method] = handler
}

func SetCmd(method string, cmdType reflect.Type) {
	cmdPool[method] = cmdType
}
