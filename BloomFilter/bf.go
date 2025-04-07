package BloomFilter

import (
	"hash"
	"hash/fnv"
)

type BloomFilter struct {
	size   int
	bits   []bool
	hashes []hash.Hash64
}

func NewBloomFilter(size int, hashCount int) *BloomFilter {
	bits := make([]bool, size)
	hashes := make([]hash.Hash64, hashCount)

	for i := 0; i < size; i++ {
		hashes[i] = fnv.New64()
	}

	return &BloomFilter{
		size:   size,
		bits:   bits,
		hashes: hashes,
	}
}

func (bf *BloomFilter) Add(value string) {
	for _, hashFunc := range bf.hashes {
		hashFunc.Reset()
		hashFunc.Write([]byte(value))
		index := hashFunc.Sum64() % uint64(bf.size)
		bf.bits[index] = true
	}
}

func (bf *BloomFilter) Contains(value string) bool {
	for _, hashFunc := range bf.hashes {
		hashFunc.Reset()
		hashFunc.Write([]byte(value))
		index := hashFunc.Sum64() % uint64(bf.size)
		if !bf.bits[index] {
			return false
		}
	}
	return true
}
