package fastrlp

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"gotest.tools/assert"
)

var header = &Header{
	ParentHash:  common.HexToHash("0000H45H"),
	UncleHash:   common.HexToHash("0000H45H"),
	Coinbase:    common.HexToAddress("0000H45H"),
	Root:        common.HexToHash("0000H00H"),
	TxHash:      common.HexToHash("0000H45H"),
	ReceiptHash: common.HexToHash("0000H45H"),
	Difficulty:  big.NewInt(1337),
	Number:      big.NewInt(1337),
	GasLimit:    1338,
	GasUsed:     1338,
	Time:        1338,
	Extra:       []byte("Extra data Extra data Extra data  Extra data  Extra data  Extra data  Extra data Extra data"),
	MixDigest:   common.HexToHash("0x0000H45H"),
}

func TestFastRLPHeader(t *testing.T) {
	assert.DeepEqual(t, header.encodeRLP(), header.encodeFastRLP())
}

func BenchmarkRLPHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		header.encodeRLP()
	}
}

func BenchmarkFastRLPHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		header.encodeFastRLP()
	}
}

func BenchmarkJSONHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		header.encodeJSON()
	}
}

func BenchmarkMsgpackHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		header.encodeMsgpack()
	}
}
