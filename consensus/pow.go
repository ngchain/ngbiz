package consensus

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger/v3"
	logging "github.com/ipfs/go-log/v2"
	"github.com/ngchain/ngbiz/ngp2p"
	"github.com/ngchain/secp256k1"
	"github.com/pkg/errors"

	"github.com/ngchain/ngbiz/blockchain"
	"github.com/ngchain/ngbiz/ngblocks"
	"github.com/ngchain/ngbiz/ngpool"
	"github.com/ngchain/ngbiz/ngstate"
	"github.com/ngchain/ngbiz/ngtypes"
)

var log = logging.Logger("pow")

// PoWork is a proof on work consensus manager.
type PoWork struct {
	PoWorkConfig

	SyncMod *syncModule

	Chain     *blockchain.Chain
	Pool      *ngpool.TxPool
	State     *ngstate.State
	LocalNode *ngp2p.LocalNode

	db *badger.DB
}

type PoWorkConfig struct {
	Network                     ngtypes.Network
	StrictMode                  bool
	SnapshotMode                bool
	DisableConnectingBootstraps bool
}

// InitPoWConsensus creates and initializes the PoW consensus.
func InitPoWConsensus(db *badger.DB, chain *blockchain.Chain, pool *ngpool.TxPool, state *ngstate.State, localNode *ngp2p.LocalNode, config PoWorkConfig) *PoWork {
	pow := &PoWork{
		PoWorkConfig: config,
		SyncMod:      nil,
		Chain:        chain,
		Pool:         pool,
		State:        state,
		LocalNode:    localNode,

		db: db,
	}

	// init sync before miner to prevent bootstrap sync from mining job update
	pow.SyncMod = newSyncModule(pow, localNode)
	if !pow.DisableConnectingBootstraps {
		pow.SyncMod.bootstrap()
	}

	// run reporter
	go pow.reportLoop()

	return pow
}

// GetBlockTemplate is a generator of new block. But the generated block has no nonce.
func (pow *PoWork) GetBlockTemplate(privateKey *secp256k1.PrivateKey) *ngtypes.Block {
	currentBlock := pow.Chain.GetLatestBlock()

	currentBlockHash := currentBlock.GetHash()

	blockTime := uint64(time.Now().Unix())

	blockHeight := currentBlock.Header.Height + 1
	newDiff := ngtypes.GetNextDiff(blockHeight, blockTime, currentBlock)

	newBlock := ngtypes.NewBareBlock(
		pow.Network,
		blockHeight,
		blockTime,
		currentBlockHash,
		newDiff,
	)

	var extraData []byte // FIXME

	genTx := pow.createGenerateTx(privateKey, blockHeight, extraData)
	txs := pow.Pool.GetPack(blockHeight)
	txsWithGen := append([]*ngtypes.Tx{genTx}, txs...)

	err := newBlock.ToUnsealing(txsWithGen)
	if err != nil {
		log.Error(err)
	}

	return newBlock
}

// GoLoop ignites all loops.
func (pow *PoWork) GoLoop() {
	go pow.eventLoop()
	go pow.SyncMod.loop()
}

// channel receiver for broadcasts events.
func (pow *PoWork) eventLoop() {
	go func() {
		for {
			block := <-pow.LocalNode.OnBlock
			err := pow.ImportBlock(block)
			if err != nil {
				log.Warnf("failed to put new block from p2p: %s", err)
			}
		}
	}()

	go func() {
		tx := <-pow.LocalNode.OnTx
		err := pow.Pool.PutTx(tx)
		if err != nil {
			log.Warnf("failed to put new tx from p2p network: %s", err)
		}
	}()
}

var ErrChainOnSyncing = errors.New("chain is syncing")

// MinedNewBlock means the local (from rpc) mined new block and need to add it into the chain.
// called by submitBlock and submitWork
func (pow *PoWork) MinedNewBlock(block *ngtypes.Block) error {
	if pow.SyncMod.Locker.IsLocked() {
		return fmt.Errorf("cannot import mined block: %w", ErrChainOnSyncing)
	}

	// check block first
	err := pow.db.Update(func(txn *badger.Txn) error {
		// check block first
		if err := pow.Chain.CheckBlock(block); err != nil {
			return err
		}

		// block is valid
		err := ngblocks.PutNewBlock(txn, block)
		if err != nil {
			return err
		}

		err = pow.State.Upgrade(txn, block) // handle Block Txs inside
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	hash := block.GetHash()
	log.Warnf("mined a new block: %x@%d", hash, block.Header.Height)

	pow.Pool.Reset()

	err = pow.LocalNode.BroadcastBlock(block)
	if err != nil {
		return fmt.Errorf("%w: failed to broadcast the new mined block", err)
	}

	return nil
}

func (pow *PoWork) ImportBlock(block *ngtypes.Block) error {
	if pow.SyncMod.Locker.IsLocked() {
		return errors.Wrap(ErrChainOnSyncing, "cannot import external block")
	}

	err := pow.Chain.ApplyBlock(block)
	if err != nil {
		return err
	}

	return nil
}
