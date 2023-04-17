package fastrlp

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"gotest.tools/assert"
)

var receipt = &types.Receipt{
	PostState:         common.Hash{2}.Bytes(),
	CumulativeGasUsed: 3,
	Logs: []*types.Log{
		{
			Address: common.BytesToAddress([]byte{0x22}),
			// derived fields:
			BlockNumber: big.NewInt(1).Uint64(),
			TxHash:      common.BytesToHash([]byte{0x03, 0x14}),
			TxIndex:     1,
			BlockHash:   common.BytesToHash([]byte{0x03, 0x14}),
			Index:       2,
		},
		{
			Address: common.BytesToAddress([]byte{0x02, 0x22}),
			// derived fields:
			BlockNumber: big.NewInt(1).Uint64(),
			TxHash:      common.BytesToHash([]byte{0x03, 0x14}),
			TxIndex:     1,
			BlockHash:   common.BytesToHash([]byte{0x03, 0x14}),
			Index:       3,
		},
	},
	// derived fields:
	TxHash:            common.BytesToHash([]byte{0x03, 0x14}),
	GasUsed:           2,
	EffectiveGasPrice: big.NewInt(22),
	BlockHash:         common.BytesToHash([]byte{0x03, 0x14}),
	BlockNumber:       big.NewInt(1),
	TransactionIndex:  1,
}

func TestFastRLPReceipt(t *testing.T) {
	assert.DeepEqual(t, encodeRLPReceipt(receipt), encodeFastRLPReceipt(receipt))
}

func BenchmarkRLPReceipt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeRLPReceipt(receipt)
	}
}

func BenchmarkFastRLPReceipt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeFastRLPReceipt(receipt)
	}
}
