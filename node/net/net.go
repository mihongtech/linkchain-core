package net

import (
	"github.com/mihongtech/linkchain-core/core"
	"github.com/mihongtech/linkchain-core/node/net/p2p/discover"
	"github.com/mihongtech/linkchain-core/node/net/p2p/peer"
)

type P2PNet interface {
	Self() *discover.Node
	AddPeer(node *discover.Node)
	Peers() []*peer.Peer
	RemovePeer(node *discover.Node)
}
type Net interface {
	core.Service
	P2PNet
}
