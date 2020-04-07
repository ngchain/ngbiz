package ngtypes

import (
	"encoding/hex"
	"math/big"
	"time"
)

var GenesisPKHex = "04ebe755cbcbc97771e285bc8d7846b80f0270327ba442cc93d47fb904cc4796d52a5e43dfe527fbd879770283641d1e092d7f59826d90e459ccd9507457970da2"
var GenesisPK, _ = hex.DecodeString(GenesisPKHex)

const (
	Version   = -1
	NetworkID = -1
)

var (
	// MinimumDifficulty is the minimum of pow difficulty because my laptop has 50 h/s, I believe you can either
	MinimumDifficulty = big.NewInt(50 * 10)
	MaxTarget         = new(big.Int).SetBytes([]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}) // new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0)) // Target = MaxTarget / diff
	GenesisTarget     = new(big.Int).Div(MaxTarget, MinimumDifficulty)
	GenesisNonce      = new(big.Int).SetUint64(0)

	genesisTimestamp = time.Date(2020, time.February, 2, 2, 2, 2, 2, time.UTC).Unix()
)

func GetBig0() *big.Int {
	return big.NewInt(0)
}

func GetBig0Bytes() []byte {
	return big.NewInt(0).Bytes()
}

func GetBig1() *big.Int {
	return big.NewInt(1)
}

var (
	BlockMaxTxsSize = 1 << 25 // 32M
)

// PoW
const (
	TargetTime      = 12 * time.Second
	BlockCheckRound = 10
	VaultCheckRound = 3
)

// Units
var (
	FloatNG        = 1000000.0
	MegaNG         = new(big.Int).Mul(NG, big.NewInt(1000000))
	MegaNGSymbol   = "MNG"
	NG             = new(big.Int).SetUint64(1000000)
	NGSymbol       = "NG"
	MicroNG        = GetBig1()
	MicroNGSymbol  = "μNG"
	OneBlockReward = new(big.Int).Mul(NG, big.NewInt(10)) // 10NG
)
