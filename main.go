package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/umbracle/fastrlp"
)

func encodeRLP(x []byte) []byte {
	res, err := rlp.EncodeToBytes(x)
	if err != nil {
		panic(err)
	}
	return res
}

func encodeFastRLP(x []byte) []byte {
	a := fastrlp.DefaultArenaPool.Get()
	defer fastrlp.DefaultArenaPool.Put(a)
	v := a.NewBytes(x)
	return v.MarshalTo(nil)
}

func main() {
	fmt.Println(encodeRLP([]byte{1, 2, 3}))
}
