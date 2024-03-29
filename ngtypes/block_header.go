package ngtypes

import (
	"bytes"
	"errors"

	"github.com/c0mm4nd/rlp"
	"github.com/cbergoon/merkletree"
	"golang.org/x/crypto/sha3"
)

// BlockHeader is the fix-sized header of the block, which is able to
// describe itself.
type BlockHeader struct {
	Network Network // 1

	Height    uint64 // 4
	Timestamp uint64 // 4

	PrevBlockHash []byte // 32
	TxTrieHash    []byte // 32
	SubTrieHash   []byte // 32

	Difficulty []byte // 32
	Nonce      []byte `rlp:"tail"` // 8
}

// CalculateHash calcs the hash of the Block header with sha3_256, aiming to
// get the merkletree hash when summarizing subs into the header.
func (x *BlockHeader) CalculateHash() ([]byte, error) {
	raw, err := rlp.EncodeToBytes(x)
	if err != nil {
		return nil, err
	}

	hash := sha3.Sum256(raw)

	return hash[:], nil
}

// ErrNotBlockHeader means the var is not a block header
var ErrNotBlockHeader = errors.New("not a block header")

// Equals checks whether the block headers equal
func (x *BlockHeader) Equals(other merkletree.Content) (bool, error) {
	header, ok := other.(*BlockHeader)
	if !ok {
		return false, ErrNotBlockHeader
	}

	if x.Network != header.Network {
		return false, nil
	}
	if x.Height != header.Height {
		return false, nil
	}
	if x.Timestamp != header.Timestamp {
		return false, nil
	}
	if !bytes.Equal(x.PrevBlockHash, header.PrevBlockHash) {
		return false, nil
	}
	if !bytes.Equal(x.TxTrieHash, header.TxTrieHash) {
		return false, nil
	}
	if !bytes.Equal(x.SubTrieHash, header.SubTrieHash) {
		return false, nil
	}
	if !bytes.Equal(x.Difficulty, header.Difficulty) {
		return false, nil
	}
	if !bytes.Equal(x.Nonce, header.Nonce) {
		return false, nil
	}

	return true, nil
}
