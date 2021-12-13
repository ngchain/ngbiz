package blockchain

import (
	"github.com/dgraph-io/badger/v3"
	logging "github.com/ipfs/go-log/v2"

	"github.com/ngchain/ngbiz/ngblocks"
	"github.com/ngchain/ngbiz/ngstate"
	"github.com/ngchain/ngbiz/ngtypes"
)

var log = logging.Logger("chain")

type Chain struct {
	*badger.DB

	*ngblocks.BlockStore
	*ngstate.State

	Network ngtypes.Network
}

func Init(db *badger.DB, network ngtypes.Network, store *ngblocks.BlockStore, state *ngstate.State) *Chain {
	chain := &Chain{
		DB: db,

		BlockStore: store,
		State:      state,

		Network: network,
	}

	return chain
}
