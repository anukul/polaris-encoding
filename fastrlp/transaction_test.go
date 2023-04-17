package fastrlp

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"gotest.tools/assert"
)

var address = common.HexToAddress("0x5")

var tx = types.NewTx(&types.LegacyTx{
	To:       &address,
	Nonce:    2,
	Value:    big.NewInt(2),
	Gas:      2,
	GasPrice: big.NewInt(22),
})

func TestFastRLPTransaction(t *testing.T) {
	assert.DeepEqual(t, encodeRLPTransaction(tx), encodeFastRLPTransaction(tx))
}

func BenchmarkRLPTransaction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeRLPTransaction(tx)
	}
}

func BenchmarkFastRLPTransaction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeFastRLPTransaction(tx)
	}
}
