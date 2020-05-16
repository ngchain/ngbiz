package storage

import (
	"bytes"
	"fmt"

	"github.com/dgraph-io/badger/v2"

	"github.com/ngchain/ngcore/ngtypes"
	"github.com/ngchain/ngcore/utils"
)

// PutNewBlock puts a new block into db.
func (c *Chain) PutNewBlock(block *ngtypes.Block) error {
	if block == nil {
		return fmt.Errorf("block is nil")
	}

	hash, _ := block.CalculateHash()
	if !bytes.Equal(hash, ngtypes.GetGenesisBlockHash()) {
		// when block is not genesis block, checking error
		if block.GetHeight() != 0 {
			if b, _ := c.GetBlockByHeight(block.GetHeight()); b != nil {
				if hashInDB, _ := b.CalculateHash(); bytes.Equal(hash, hashInDB) {
					return nil
				}

				return fmt.Errorf("has block in same height: %s", b)
			}
		}

		if _, err := c.GetBlockByHash(block.GetPrevHash()); err != nil {
			return fmt.Errorf("no prev block in storage: %x, %v", block.GetPrevHash(), err)
		}
	}

	err := c.db.Update(func(txn *badger.Txn) error {
		raw, _ := utils.Proto.Marshal(block)
		log.Infof("putting block@%d: %x", block.Header.Height, hash)

		// put block hash & height
		err := txn.Set(append(blockPrefix, hash...), raw)
		if err != nil {
			return err
		}
		err = txn.Set(append(blockPrefix, utils.PackUint64LE(block.Header.Height)...), hash)
		if err != nil {
			return err
		}

		// put txs
		err = putTxs(txn, block.Txs...)
		if err != nil {
			return err
		}

		// update helper
		err = txn.Set(append(blockPrefix, latestHeightTag...), utils.PackUint64LE(block.Header.Height))
		if err != nil {
			return err
		}
		err = txn.Set(append(blockPrefix, latestHashTag...), hash)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

func putTxs(txn *badger.Txn, txs ...*ngtypes.Tx) error {
	for i := range txs {
		hash, _ := txs[i].CalculateHash()

		raw, err := utils.Proto.Marshal(txs[i])
		if err != nil {
			return err
		}

		err = txn.Set(append(txPrefix, hash...), raw)
		if err != nil {
			return err
		}
	}

	return nil
}