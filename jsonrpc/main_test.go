package jsonrpc_test

import (
	"testing"
	"time"

	"github.com/ngchain/ngbiz/blockchain"
	"github.com/ngchain/ngbiz/consensus"
	"github.com/ngchain/ngbiz/jsonrpc"
	"github.com/ngchain/ngbiz/ngblocks"
	"github.com/ngchain/ngbiz/ngp2p"
	"github.com/ngchain/ngbiz/ngpool"
	"github.com/ngchain/ngbiz/ngstate"
	"github.com/ngchain/ngbiz/ngtypes"
	"github.com/ngchain/ngbiz/storage"
)

// TODO: add tests for each method rather than testing the server.
func TestNewRPCServer(t *testing.T) {
	network := ngtypes.ZERONET

	db := storage.InitMemStorage()
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	store := ngblocks.Init(db, network)
	state := ngstate.InitStateFromGenesis(db, network)

	chain := blockchain.Init(db, network, store, state)
	chain.CheckHealth(network)

	localNode := ngp2p.InitLocalNode(chain, ngp2p.P2PConfig{
		P2PKeyFile:       "p2p.key",
		Network:          network,
		Port:             52520,
		DisableDiscovery: true,
	})
	localNode.GoServe()

	pool := ngpool.Init(db, chain, localNode)

	pow := consensus.InitPoWConsensus(
		db,
		chain,
		pool,
		state,
		localNode,
		consensus.PoWorkConfig{
			Network:                     network,
			DisableConnectingBootstraps: true,
		},
	)
	pow.GoLoop()

	rpc := jsonrpc.NewServer(pow, jsonrpc.ServerConfig{
		Host:                 "",
		Port:                 52521,
		DisableP2PMethods:    false,
		DisableMiningMethods: false,
	})
	go rpc.Serve()

	go func() {
		finished := time.After(2 * time.Minute)

		for {
			<-finished

			return
		}
	}()
}
