package fastssz

import (
	"github.com/ethereum/go-ethereum/common"
	"gotest.tools/assert"
	"math/big"
	"testing"
)

func convert(x *big.Int, len int) []byte {
	b := make([]byte, len)
	x.FillBytes(b)
	return b
}

var header = &Header{
	ParentHash:      common.HexToHash("0000H45H").Bytes(),
	UncleHash:       common.HexToHash("0000H45H").Bytes(),
	Coinbase:        common.HexToAddress("0000H45H").Bytes(),
	Root:            common.HexToHash("0000H00H").Bytes(),
	TxHash:          common.HexToHash("0000H45H").Bytes(),
	ReceiptHash:     common.HexToHash("0000H45H").Bytes(),
	Bloom:           make([]byte, 256),
	Difficulty:      convert(big.NewInt(1337), 32),
	Number:          1337,
	GasLimit:        1338,
	GasUsed:         1338,
	Time:            1338,
	Extra:           []byte("Extra data Extra data Extra data"),
	MixDigest:       common.HexToHash("0x0000H45H").Bytes(),
	Nonce:           convert(big.NewInt(1338), 8),
	BaseFee:         convert(big.NewInt(1338), 32),
	WithdrawalsHash: common.HexToHash("0000H45H").Bytes(),
}

func TestFastRLPHeader(t *testing.T) {
	assert.DeepEqual(t, header.encodeRLP(), header.encodeFastRLP())
}

func TestFastSSZHeader(t *testing.T) {
	e := header.encodeSSZ()
	h := &Header{}
	if err := h.UnmarshalSSZ(e); err != nil {
		panic(err)
	}
	assert.DeepEqual(t, h, header)
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

func BenchmarkSSZHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		header.encodeSSZ()
	}
}
