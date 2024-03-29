package wired

import (
	"github.com/c0mm4nd/rlp"
	logging "github.com/ipfs/go-log/v2"
	core "github.com/libp2p/go-libp2p-core"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-msgio"

	"github.com/ngchain/ngbiz/blockchain"
	"github.com/ngchain/ngbiz/ngp2p/defaults"
	"github.com/ngchain/ngbiz/ngtypes"
)

var log = logging.Logger("wired")

// Wired type.
type Wired struct {
	network ngtypes.Network
	host    core.Host // local host

	protocolID protocol.ID

	chain *blockchain.Chain
}

func NewWiredProtocol(host core.Host, network ngtypes.Network, chain *blockchain.Chain) *Wired {
	w := &Wired{
		network: network,
		host:    host,

		protocolID: protocol.ID(defaults.GetWiredProtocol(network)),

		chain: chain,
	}

	return w
}

func (w *Wired) GetWiredProtocol() protocol.ID {
	return w.protocolID
}

func (w *Wired) GoServe() {
	// register handler
	w.host.SetStreamHandler(w.protocolID, func(stream network.Stream) {
		log.Debugf("handling new stream from %s", stream.Conn().RemotePeer())
		go w.handleStream(stream)
	})
}

func (w *Wired) handleStream(stream network.Stream) {
	r := msgio.NewReader(stream)
	raw, err := r.ReadMsg()
	if err != nil {
		log.Error(err)
		return
	}

	// unmarshal it
	var msg Message

	err = rlp.DecodeBytes(raw, &msg)
	if err != nil {
		log.Error(err)
		return
	}

	if !Verify(stream.Conn().RemotePeer(), &msg) {
		w.sendReject(msg.Header.ID, stream, ErrMsgSignInvalid)
		return
	}

	switch msg.Header.Type {
	case PingMsg:
		w.onPing(stream, &msg)
	case GetChainMsg:
		w.onGetChain(stream, &msg)
	case GetSheetMsg:
		w.onGetChain(stream, &msg)
	default:
		w.sendReject(msg.Header.ID, stream, ErrMsgTypeInvalid)
	}

	err = stream.Close()
	if err != nil {
		log.Error(err)
	}
}
