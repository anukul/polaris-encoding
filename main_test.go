package main

import (
	"crypto/rand"
	"gotest.tools/assert"
	"testing"
)

func generateSlice(len uint) []byte {
	res := make([]byte, len)
	_, err := rand.Read(res)
	if err != nil {
		panic(err)
	}
	return res
}

var input = generateSlice(1000000)

func TestFastRLP(t *testing.T) {
	assert.DeepEqual(t, encodeRLP(input), encodeFastRLP(input))
}

func BenchmarkRLP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeRLP(input)
	}
}

func BenchmarkFastRLP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeFastRLP(input)
	}
}
