package ngp2p

import (
	"context"
	"sync"

	"github.com/libp2p/go-libp2p"
	core "github.com/libp2p/go-libp2p-core"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/multiformats/go-multiaddr"
	"go.uber.org/atomic"
)

const DHTProtocolExtension = "/ngp2p/dht/" + protocolVersion

var p2pDHT *dht.IpfsDHT

func getPublicRouter() libp2p.Option {
	return libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
		var err error
		p2pDHT, err = dht.New(context.Background(), h, dht.Mode(dht.ModeAutoServer), dht.ProtocolExtension(DHTProtocolExtension))
		return p2pDHT, err
	})
}

// active DHT
func activeDHT(ctx context.Context, kademliaDHT *dht.IpfsDHT, host core.Host) {
	err := kademliaDHT.Bootstrap(ctx)
	if err != nil {
		panic(err)
	}

	connectToDHTBootstrapNodes(ctx, host, BootstrapNodes)
}

func connectToDHTBootstrapNodes(ctx context.Context, h host.Host, mas []multiaddr.Multiaddr) int32 {
	var wg sync.WaitGroup
	var numConnected atomic.Int32
	for _, ma := range mas {
		wg.Add(1)
		go func(ma multiaddr.Multiaddr) {
			pi, err := peerstore.InfoFromP2pAddr(ma)
			if err != nil {
				panic(err)
			}
			defer wg.Done()
			err = h.Connect(ctx, *pi)
			if err != nil {
				log.Errorf("error connecting to bootstrap node %q: %v", ma, err)
			} else {
				numConnected.Inc()
			}
		}(ma)
	}
	wg.Wait()
	return numConnected.Load()
}
