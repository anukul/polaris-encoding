package fastrlp

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ugorji/go/codec"
	"github.com/umbracle/fastrlp"
)

type Header types.Header

func (h *Header) encodeRLP() []byte {
	buf, err := rlp.EncodeToBytes(h)
	if err != nil {
		panic(err)
	}
	return buf
}

func (h *Header) encodeFastRLP() []byte {
	a := fastrlp.DefaultArenaPool.Get()
	defer fastrlp.DefaultArenaPool.Put(a)

	v := a.NewArray()
	v.Set(a.NewBytes(h.ParentHash[:]))
	v.Set(a.NewBytes(h.UncleHash[:]))
	v.Set(a.NewBytes(h.Coinbase[:]))
	v.Set(a.NewBytes(h.Root[:]))
	v.Set(a.NewBytes(h.TxHash[:]))
	v.Set(a.NewBytes(h.ReceiptHash[:]))
	v.Set(a.NewBytes(h.Bloom[:]))

	v.Set(a.NewUint(h.Difficulty.Uint64()))
	v.Set(a.NewUint(h.Number.Uint64()))
	v.Set(a.NewUint(h.GasLimit))
	v.Set(a.NewUint(h.GasUsed))
	v.Set(a.NewUint(h.Time))

	v.Set(a.NewBytes(h.Extra))
	v.Set(a.NewBytes(h.MixDigest[:]))
	v.Set(a.NewBytes(h.Nonce[:]))

	return v.MarshalTo(nil)
}

func (h *Header) encodeJSON() []byte {
	res, err := json.Marshal(h)
	if err != nil {
		panic(err)
	}
	return res
}

var mh codec.MsgpackHandle
var enc = codec.NewEncoderBytes(nil, &mh)

func (h *Header) encodeMsgpack() (b []byte) {
	enc.ResetBytes(&b)
	if err := enc.Encode(h); err != nil {
		panic(err)
	}
	return b
}
