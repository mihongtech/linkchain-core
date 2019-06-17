package server

import (
	"reflect"
)

/**rpc handler,server,cmd,closeChan,context**/
type commandHandler func(*Server, interface{}, <-chan struct{}) (interface{}, error)

func (s *Server) SetHandleFunc(method string, handler commandHandler) {
	s.handlerPool[method] = handler
}

func (s *Server) SetCmd(method string, cmdType reflect.Type) {
	s.cmdPool[method] = cmdType
}
