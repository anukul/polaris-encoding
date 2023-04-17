package fastssz

import (
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/umbracle/fastrlp"
)

type Header struct {
	ParentHash      []byte `ssz-size:"32"`
	UncleHash       []byte `ssz-size:"32"`
	Coinbase        []byte `ssz-size:"20"`
	Root            []byte `ssz-size:"32"`
	TxHash          []byte `ssz-size:"32"`
	ReceiptHash     []byte `ssz-size:"32"`
	Bloom           []byte `ssz-size:"256"`
	Difficulty      []byte `ssz-size:"32"`
	Number          uint64
	GasLimit        uint64
	GasUsed         uint64
	Time            uint64
	Extra           []byte `ssz-max:"32"`
	MixDigest       []byte `ssz-size:"32"`
	Nonce           []byte `ssz-size:"8"`
	BaseFee         []byte `ssz-size:"32"`
	WithdrawalsHash []byte `ssz-size:"32"`
}

func (h *Header) encodeRLP() []byte {
	res, err := rlp.EncodeToBytes(h)
	if err != nil {
		panic(err)
	}
	return res
}

func (h *Header) encodeFastRLP() []byte {
	a := fastrlp.DefaultArenaPool.Get()
	defer fastrlp.DefaultArenaPool.Put(a)

	v := a.NewArray()
	v.Set(a.NewBytes(h.ParentHash))
	v.Set(a.NewBytes(h.UncleHash))
	v.Set(a.NewBytes(h.Coinbase))
	v.Set(a.NewBytes(h.Root))
	v.Set(a.NewBytes(h.TxHash))
	v.Set(a.NewBytes(h.ReceiptHash))
	v.Set(a.NewBytes(h.Bloom))

	v.Set(a.NewBytes(h.Difficulty))
	v.Set(a.NewUint(h.Number))
	v.Set(a.NewUint(h.GasLimit))
	v.Set(a.NewUint(h.GasUsed))
	v.Set(a.NewUint(h.Time))

	v.Set(a.NewBytes(h.Extra))
	v.Set(a.NewBytes(h.MixDigest))
	v.Set(a.NewBytes(h.Nonce))
	v.Set(a.NewBytes(h.BaseFee))
	v.Set(a.NewBytes(h.WithdrawalsHash))

	return v.MarshalTo(nil)
}

func (h *Header) encodeSSZ() []byte {
	res, err := h.MarshalSSZ()
	if err != nil {
		panic(err)
	}
	return res
}
