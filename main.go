package main

import (
	"github.com/mihongtech/linkchain-core/node"
	"github.com/mihongtech/linkchain-core/node/config"
)

func main() {
	n := node.NewNode(config.BaseConfig{})
	n.Setup(nil)
	n.Start()
	n.Stop()
}
