package btcutil

import (

	"github.com/hawkit/btcd-demo/wire"
	"github.com/hawkit/btcd-demo/chaincfg/chainhash"
	"bytes"
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

// Bytes returns the serialized bytes for the Block. This is equivalent to
// calling Serialize on the underlying wire.MsgBlock, however it caches the
// result so subsequent calls are more efficient.
func (b *Block) Bytes() ([]byte, error)  {

	// Return the cached serialized bytes if it has already been generated.
	if len(b.serializedBlock) != 0 {
		return b.serializedBlock, nil
	}

	// Serialize the MsgBlock.
	w := bytes.NewBuffer(make([]byte, 0, b.MsgBlock.SerializeSize()))

	err := b.MsgBlock.Serialize(w)
	if err != nil {
		return nil, err
	}

	serailzedBlock := w.Bytes()
	b.serializedBlock = serailzedBlock
	return serailzedBlock, nil

}

func (b *Block) Hash() *chainhash.Hash {
	if b.blockHash != nil {
		return b.blockHash
	}
	// Cache the block hash and return it.
	hash := b.MsgBlock.BlockHash()
	b.blockHash = &hash
	return  &hash
}