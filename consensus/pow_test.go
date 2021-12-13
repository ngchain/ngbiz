package consensus_test

import (
	"testing"

	"github.com/ngchain/ngbiz/blockchain"
	"github.com/ngchain/ngbiz/consensus"
	"github.com/ngchain/ngbiz/ngblocks"
	"github.com/ngchain/ngbiz/ngp2p"
	"github.com/ngchain/ngbiz/ngpool"
	"github.com/ngchain/ngbiz/ngstate"
	"github.com/ngchain/ngbiz/ngtypes"
	"github.com/ngchain/ngbiz/storage"
)

func TestNewConsensusManager(t *testing.T) {
	db := storage.InitMemStorage()

	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	net := ngtypes.ZERONET
	store := ngblocks.Init(db, net)
	state := ngstate.InitStateFromGenesis(db, net)
	chain := blockchain.Init(db, net, store, nil)

	localNode := ngp2p.InitLocalNode(chain, ngp2p.P2PConfig{
		Network:          net,
		Port:             52520,
		DisableDiscovery: true,
	})
	pool := ngpool.Init(db, chain, localNode)

	consensus.InitPoWConsensus(db, chain, pool, state, localNode, consensus.PoWorkConfig{
		Network:                     net,
		DisableConnectingBootstraps: true,
	})
}
