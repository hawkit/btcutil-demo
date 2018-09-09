package btcutil

import (

	"github.com/hawkit/btcd-demo/wire"
	"github.com/hawkit/btcd-demo/chaincfg/chainhash"
)

// Block defines a bitcoin block that provides easier and more efficent
// manipulation of raw blocks. It also memoizes hashes for the block
// and its transactions on their first access so subsequent accesses
// don't have to repeat the relatively expensive hashing operations.
type Block struct {
	MsgBlock *wire.MsgBlock // Underlying MsgBlock
	serializedBlock []byte // Serialized bytes for the block
	serializedBlockNoWithness []byte // Serialized bytes for block w/o witeness data
	blockHash *chainhash.Hash // Cached block hash
	blockHeight int32 // Height in the main block chain
	transactions []*Tx // Transactions
	txnsGenerated bool // All wrapped transactions generated
}