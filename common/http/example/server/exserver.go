package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/mihongtech/linkchain-core/common/http/example"
	"github.com/mihongtech/linkchain-core/common/http/server"
)

func main() {
	httpConfig := server.NewConfig(time.Now().Unix(), "localhost:8081", "mihongtech", "mihongtech")
	rpcServe, err := server.NewRPCServer(httpConfig)
	if err != nil {
		fmt.Printf("start rpc server:%s", err)
		return
	}
	server.SetHandleFunc("getinfo", getinfo)
	server.SetCmd("getinfo", reflect.TypeOf((*example.InfoCmd)(nil)))
	rpcServe.Start()

	select {
	case <-rpcServe.RequestedProcessShutdown():
		fmt.Printf("stop:%s", err)
	}

	//go func() {
	//
	//	server.shutdownRequestChannel <- struct{}{}
	//}()
}

func getinfo(s *server.Server, cmd interface{}, closeChan <-chan struct{}) (interface{}, error) {
	c, ok := cmd.(*example.InfoCmd)
	if !ok {
		fmt.Println("Type error:", reflect.TypeOf(cmd))
		return nil, nil
	}
	fmt.Printf("input id :%s", c.Id)
	c.Id += "222"
	return c.Id, nil
}
