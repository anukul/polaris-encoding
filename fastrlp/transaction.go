package fastrlp

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/umbracle/fastrlp"
)

func encodeRLPTransaction(tx *types.Transaction) []byte {
	buf, err := rlp.EncodeToBytes(tx)
	if err != nil {
		panic(err)
	}
	return buf
}

func encodeFastRLPTransaction(tx *types.Transaction) []byte {
	var dst []byte

	//if int(tx.Type()) != types.LegacyTxType {
	//	dst = append(dst, tx.Type())
	//}

	a := fastrlp.DefaultArenaPool.Get()
	defer fastrlp.DefaultArenaPool.Put(a)

	vv := a.NewArray()

	vv.Set(a.NewUint(tx.Nonce()))
	vv.Set(a.NewBigInt(tx.GasPrice()))
	vv.Set(a.NewUint(tx.Gas()))

	// Address may be empty
	if tx.To() != nil {
		vv.Set(a.NewBytes((*tx.To()).Bytes()))
	} else {
		vv.Set(a.NewNull())
	}

	vv.Set(a.NewBigInt(tx.Value()))
	vv.Set(a.NewCopyBytes(tx.Data()))

	v, r, s := tx.RawSignatureValues()

	// signature values
	vv.Set(a.NewBigInt(v))
	vv.Set(a.NewBigInt(r))
	vv.Set(a.NewBigInt(s))

	//if tx.Type() == StateTx {
	//	vv.Set(a.NewBytes((tx.From).Bytes()))
	//}

	return vv.MarshalTo(dst)
}
