package ngblocks

import (
	"bytes"
	"errors"

	"github.com/c0mm4nd/rlp"
	"github.com/dgraph-io/badger/v3"

	"github.com/ngchain/ngbiz/ngtypes"
	"github.com/ngchain/ngbiz/utils"
)

// initWithGenesis will initialize the store with genesis block & vault.
func (store *BlockStore) initWithGenesis() {
	if !store.hasGenesisBlock(store.Network) {
		log.Warnf("initializing with genesis block")

		block := ngtypes.GetGenesisBlock(store.Network)

		if err := store.Update(func(txn *badger.Txn) error {
			hash := block.GetHash()
			raw, _ := rlp.EncodeToBytes(block)
			log.Infof("putting block@%d: %x", block.Header.Height, hash)
			err := txn.Set(append(blockPrefix, hash...), raw)
			if err != nil {
				return err
			}
			err = txn.Set(append(blockPrefix, utils.PackUint64LE(block.Header.Height)...), hash)
			if err != nil {
				return err
			}

			err = txn.Set(append(blockPrefix, latestHeightTag...), utils.PackUint64LE(block.Header.Height))
			if err != nil {
				return err
			}
			err = txn.Set(append(blockPrefix, latestHashTag...), hash)
			if err != nil {
				return err
			}

			err = txn.Set(append(blockPrefix, originHeightTag...), utils.PackUint64LE(block.Header.Height))
			if err != nil {
				return err
			}
			err = txn.Set(append(blockPrefix, originHashTag...), hash)
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
			panic(err)
		}
	}
}

// hasGenesisBlock checks whether the genesis block is in db.
func (store *BlockStore) hasGenesisBlock(network ngtypes.Network) bool {
	has := false

	if err := store.View(func(txn *badger.Txn) error {
		item, err := txn.Get(append(blockPrefix, utils.PackUint64LE(0)...))
		if err != nil {
			return err
		}
		hash, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		if hash != nil {
			has = true
		}
		if !bytes.Equal(hash, ngtypes.GetGenesisBlock(network).GetHash()) {
			panic("wrong genesis block in db")
		}

		return nil
	}); err != nil && !errors.Is(err, badger.ErrKeyNotFound) {
		panic(err)
	}

	return has
}

// hasOrigin checks whether the genesis vault is in db.
func (store *BlockStore) hasOrigin(network ngtypes.Network) bool {
	has := false

	if err := store.View(func(txn *badger.Txn) error {
		item, err := txn.Get(append(blockPrefix, originHeightTag...))
		if err != nil {
			return err
		}

		has = item != nil

		item, err = txn.Get(append(blockPrefix, originHashTag...))
		if err != nil {
			return err
		}

		has = has && item != nil

		hash, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		item, err = txn.Get(append(blockPrefix, hash...))
		if err != nil {
			return err
		}
		rawBlock, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		var originBlock ngtypes.Block
		err = rlp.DecodeBytes(rawBlock, &originBlock)
		if err != nil {
			return err
		}

		has = has && originBlock.Header.Network == network

		return nil
	}); err != nil && !errors.Is(err, badger.ErrKeyNotFound) {
		panic(err)
	}

	return has
}

// initWithBlockchain initialize the store by importing the external store.
// func (store *BlockStore) initWithBlockchain(blocks ...*ngtypes.Block) error {
//	/* Put start */
//	err := store.Update(func(txn *badger.Txn) error {
//		for i := 0; i < len(blocks); i++ {
//			block := blocks[i]
//			hash := block.Hash()
//			raw, _ := rlp.EncodeToBytes(block)
//			log.Infof("putting block@%d: %x", block.Height, hash)
//			err := txn.Set(append(blockPrefix, hash...), raw)
//			if err != nil {
//				return err
//			}
//			err = txn.Set(append(blockPrefix, utils.PackUint64LE(block.Height)...), hash)
//			if err != nil {
//				return err
//			}
//			err = txn.Set(append(blockPrefix, latestHeightTag...), utils.PackUint64LE(block.Height))
//			if err != nil {
//				return err
//			}
//			err = txn.Set(append(blockPrefix, latestHashTag...), hash)
//			if err != nil {
//				return err
//			}
//		}
//		return nil
//	})
//
//	return err
// }

func (store *BlockStore) InitFromCheckpoint(block *ngtypes.Block) error {
	err := store.Update(func(txn *badger.Txn) error {
		hash := block.GetHash()
		raw, _ := rlp.EncodeToBytes(block)
		log.Infof("putting block@%d: %x", block.Header.Height, hash)
		err := txn.Set(append(blockPrefix, hash...), raw)
		if err != nil {
			return err
		}
		err = txn.Set(append(blockPrefix, utils.PackUint64LE(block.Header.Height)...), hash)
		if err != nil {
			return err
		}

		err = txn.Set(append(blockPrefix, latestHeightTag...), utils.PackUint64LE(block.Header.Height))
		if err != nil {
			return err
		}
		err = txn.Set(append(blockPrefix, latestHashTag...), hash)
		if err != nil {
			return err
		}

		err = txn.Set(append(blockPrefix, originHeightTag...), utils.PackUint64LE(block.Header.Height))
		if err != nil {
			return err
		}
		err = txn.Set(append(blockPrefix, originHashTag...), hash)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
