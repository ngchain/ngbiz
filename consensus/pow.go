package consensus

import (
	"fmt"
	"time"

	"github.com/ngchain/ngcore/ngtypes/ngproto"

	"github.com/dgraph-io/badger/v3"
	logging "github.com/ipfs/go-log/v2"

	"github.com/ngchain/secp256k1"

	"github.com/ngchain/ngcore/blockchain"
	"github.com/ngchain/ngcore/ngblocks"
	"github.com/ngchain/ngcore/ngp2p"
	"github.com/ngchain/ngcore/ngpool"
	"github.com/ngchain/ngcore/ngstate"
	"github.com/ngchain/ngcore/ngtypes"
)

var log = logging.Logger("pow")

// PoWork is a proof on work consensus manager
type PoWork struct {
	PoWorkConfig

	SyncMod *syncModule

	Chain     *blockchain.Chain
	Pool      *ngpool.TxPool
	State     *ngstate.State
	LocalNode *ngp2p.LocalNode

	db *badger.DB

	// for miner
	foundBlockCh chan *ngtypes.Block
}

type PoWorkConfig struct {
	Network                     ngproto.NetworkType
	StrictMode                  bool
	SnapshotMode                bool
	DisableConnectingBootstraps bool

	MiningThread int
	PrivateKey   *secp256k1.PrivateKey
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

		foundBlockCh: make(chan *ngtypes.Block),
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
func (pow *PoWork) GetBlockTemplate() *ngtypes.Block {
	currentBlock := pow.Chain.GetLatestBlock()

	currentBlockHash := currentBlock.GetHash()

	blockTime := time.Now().Unix()

	blockHeight := currentBlock.Height + 1
	newDiff := ngtypes.GetNextDiff(blockHeight, blockTime, currentBlock)

	newBareBlock := ngtypes.NewBareBlock(
		pow.Network,
		blockHeight,
		blockTime,
		currentBlockHash,
		newDiff,
	)

	var extraData []byte // FIXME

	genTx := pow.createGenerateTx(blockHeight, extraData)
	txs := pow.Pool.GetPack(currentBlockHash).Txs
	txsWithGen := append([]*ngtypes.Tx{genTx}, txs...)

	newUnsealingBlock, err := newBareBlock.ToUnsealing(txsWithGen)
	if err != nil {
		log.Error(err)
	}

	return newUnsealingBlock
}

// GoLoop ignites all loops
func (pow *PoWork) GoLoop() {
	go pow.eventLoop()
	go pow.SyncMod.loop()
}

// channel receiver for broadcasts events
func (pow *PoWork) eventLoop() {
	for {
		select {
		case block := <-pow.LocalNode.OnBlock:
			err := pow.ImportBlock(block)
			if err != nil {
				log.Warnf("failed to put new block from p2p: %s", err)
				continue
			}
		case tx := <-pow.LocalNode.OnTx:
			err := pow.Pool.PutTx(tx)
			if err != nil {
				log.Warnf("failed to put new tx from p2p network: %s", err)
			}

		case newBlock := <-pow.foundBlockCh:
			err := pow.MinedNewBlock(newBlock)
			if err != nil {
				log.Warnf("error on handling the mined block: %s", err)
			}
		default:
			// miner

		}
	}
}

// MinedNewBlock means the consensus mined new block and need to add it into the chain.
func (pow *PoWork) MinedNewBlock(block *ngtypes.Block) error {
	if pow.SyncMod.Locker.IsLocked() {
		return fmt.Errorf("cannot import mined block: chain is syncing")
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
	log.Warnf("mined a new block: %x@%d", hash, block.GetHeight())

	pow.Pool.Reset()

	err = pow.LocalNode.BroadcastBlock(block)
	if err != nil {
		return fmt.Errorf("failed to broadcast the new mined block")
	}

	return nil
}

func (pow *PoWork) ImportBlock(block *ngtypes.Block) error {
	if pow.SyncMod.Locker.IsLocked() {
		return fmt.Errorf("cannot import external block: chain is syncing")
	}

	err := pow.Chain.ApplyBlock(block)
	if err != nil {
		return err
	}

	return nil
}
