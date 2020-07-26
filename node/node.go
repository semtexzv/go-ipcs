package node

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/p2p/discovery"
	"github.com/semtexzv/go-ipcs/base"
	"time"

	"github.com/libp2p/go-libp2p"
)

const ID = "/ipcs/0.0.0"

var ctx = context.Background()

type Behavior struct {
	priv crypto.PrivKey
	pub  crypto.PubKey

	node host.Host
}

func (s *Behavior) HandlePeerFound(info peer.AddrInfo) {
	fmt.Println("Peer found", info.ID)
	base.Unwrap(s.node.Connect(ctx, info))
}

func Run() {

	priv, pub, err := crypto.GenerateKeyPairWithReader(crypto.Ed25519, 256, rand.Reader)
	base.Unwrap(err)

	opts := []libp2p.Option{
		libp2p.Identity(priv),
	}

	node, err := libp2p.New(ctx, opts...)
	base.Unwrap(err)
	fmt.Println("Listening on: ", node.ID(),node.Addrs())

	behavior := Behavior{
		priv: priv,
		pub:  pub,
		node: node,
	}

	mdns, err := discovery.NewMdnsService(ctx, node, time.Second, "")
	base.Unwrap(err)
	mdns.RegisterNotifee(&behavior)

	_, err = dht.New(ctx, node, dht.ProtocolPrefix(ID), dht.BucketSize(20))
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Minute)

}
