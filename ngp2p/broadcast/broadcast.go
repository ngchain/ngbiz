package broadcast

import (
	"context"

	logging "github.com/ipfs/go-log/v2"
	core "github.com/libp2p/go-libp2p-core"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	"github.com/ngchain/ngbiz/ngp2p/defaults"
	"github.com/ngchain/ngbiz/ngtypes"
)

type Broadcast struct {
	PubSub *pubsub.PubSub
	node   core.Host

	network       ngtypes.Network
	topics        map[string]*pubsub.Topic
	subscriptions map[string]*pubsub.Subscription

	blockTopic string
	txTopic    string

	OnBlock chan *ngtypes.Block
	OnTx    chan *ngtypes.Tx
}

var log = logging.Logger("bcast")

func NewBroadcastProtocol(node core.Host, network ngtypes.Network, blockCh chan *ngtypes.Block, txCh chan *ngtypes.Tx) *Broadcast {
	var err error

	b := &Broadcast{
		PubSub:        nil,
		node:          node,
		network:       network,
		topics:        make(map[string]*pubsub.Topic),
		subscriptions: make(map[string]*pubsub.Subscription),

		blockTopic: defaults.GetBroadcastBlockTopic(network),
		txTopic:    defaults.GetBroadcastTxTopic(network),

		OnBlock: blockCh,
		OnTx:    txCh,
	}

	b.PubSub, err = pubsub.NewFloodSub(context.Background(), node)
	if err != nil {
		panic(err)
	}

	b.topics[b.blockTopic], err = b.PubSub.Join(b.blockTopic)
	if err != nil {
		panic(err)
	}

	b.subscriptions[b.blockTopic], err = b.topics[b.blockTopic].Subscribe()
	if err != nil {
		panic(err)
	}

	b.topics[b.txTopic], err = b.PubSub.Join(b.txTopic)
	if err != nil {
		panic(err)
	}

	b.subscriptions[b.txTopic], err = b.topics[b.txTopic].Subscribe()
	if err != nil {
		panic(err)
	}

	return b
}

func (b *Broadcast) GoServe() {
	go b.blockListener(b.subscriptions[b.blockTopic])
	go b.txListener(b.subscriptions[b.txTopic])
}

func (b *Broadcast) blockListener(sub *pubsub.Subscription) {
	for {
		msg, err := sub.Next(context.Background())
		if err != nil {
			log.Error(err)
			continue
		}

		go b.onBroadcastBlock(msg)
	}
}

func (b *Broadcast) txListener(sub *pubsub.Subscription) {
	for {
		msg, err := sub.Next(context.Background())
		if err != nil {
			log.Error(err)
			continue
		}

		go b.onBroadcastTx(msg)
	}
}
