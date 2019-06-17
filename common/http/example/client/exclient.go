package main

import (
	"fmt"
	"github.com/mihongtech/linkchain-core/common/http/client"
	"github.com/mihongtech/linkchain-core/common/http/example"
)

func main() {
	httpConfig := client.Config{RPCServer: "localhost:8081", RPCPassword: "mihongtech", RPCUser: "mihongtech"}
	getInfo := &example.InfoCmd{"111"}
	result, err := client.RPC("getinfo", getInfo, &httpConfig)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("success", result)
}
