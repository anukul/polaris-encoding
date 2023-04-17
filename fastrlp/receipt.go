package fastrlp

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/umbracle/fastrlp"
)

func encodeLog(l *types.Log, a *fastrlp.Arena) *fastrlp.Value {
	v := a.NewArray()
	v.Set(a.NewBytes(l.Address.Bytes()))

	topics := a.NewArray()
	for _, t := range l.Topics {
		topics.Set(a.NewBytes(t.Bytes()))
	}

	v.Set(topics)
	v.Set(a.NewBytes(l.Data))

	return v
}

func encodeReceiptLogs(r *types.Receipt, a *fastrlp.Arena) *fastrlp.Value {
	if len(r.Logs) == 0 {
		// There are no receipts, write the RLP null array entry
		return a.NewNullArray()
	}

	logs := a.NewArray()

	for _, l := range r.Logs {
		logs.Set(encodeLog(l, a))
	}

	return logs
}

func encodeFastRLPReceipt(r *types.Receipt) []byte {
	var dst []byte

	if int(r.Type) != types.LegacyTxType {
		dst = append(dst, r.Type)
	}

	a := fastrlp.DefaultArenaPool.Get()
	defer fastrlp.DefaultArenaPool.Put(a)

	v := a.NewArray()

	if r.Status != 0 {
		v.Set(a.NewUint(r.Status))
	} else {
		v.Set(a.NewBytes(r.PostState))
	}

	v.Set(a.NewUint(r.CumulativeGasUsed))
	v.Set(a.NewBytes(r.Bloom[:]))
	v.Set(encodeReceiptLogs(r, a))

	return v.MarshalTo(dst)
}

func encodeRLPReceipt(r *types.Receipt) []byte {
	buf, err := rlp.EncodeToBytes(r)
	if err != nil {
		panic(err)
	}
	return buf
}
