package pool

import (
	"github.com/ngchain/ngcore/ngtypes"
)

// DelBlockTxs will popping txs from txpool.
func (p *TxPool) DelBlockTxs(txs ...*ngtypes.Tx) {
	// p.Lock()
	// defer p.Unlock()

	// for i := range txs {
	// 	if p.Queuing[txs[i].Header.GetConvener()] != nil {
	// 		delete(p.Queuing[txs[i].Header.GetConvener()], txs[i].Header.GetN())
	// 	}
	// }
}