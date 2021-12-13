package defaults

import (
	"encoding/hex"

	"github.com/ngchain/ngbiz/ngtypes"
)

// MaxBlocks limits the max number of blocks which are transfered on p2p network.
const MaxBlocks = 1000

// pattern: /ngp2p/protocol-name/version
const (
	protocolVersion = "/0.0.1"
)

func getGenesisBlockHash(network ngtypes.Network) string {
	return hex.EncodeToString(ngtypes.GetGenesisBlock(network).GetHash())
}

func GetWiredProtocol(network ngtypes.Network) string {
	return "/ngp2p/wired/" + getGenesisBlockHash(network) + protocolVersion
}

func GetDHTProtocolExtension(network ngtypes.Network) string {
	return "/ngp2p/dht/" + getGenesisBlockHash(network) + protocolVersion
}

func GetBroadcastBlockTopic(network ngtypes.Network) string {
	return "/ngp2p/broadcast/block/" + getGenesisBlockHash(network) + protocolVersion
}

func GetBroadcastTxTopic(network ngtypes.Network) string {
	return "/ngp2p/broadcast/tx/" + getGenesisBlockHash(network) + protocolVersion
}
