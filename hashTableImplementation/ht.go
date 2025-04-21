package main

import (
	"encoding/binary"
	"errors"
	"hash"
	"hash/fnv"
)

type HashTable interface {
	Put(key int, value string)
	Get(key int) (string, error)
	Delete(key int) error
	Update(key int, value string) error
}

// способ разерешения коллизий - метод повторного хэширования (двойное хэширование)
type itemRehash struct {
	key     int
	value   string
	deleted bool
}

// для ускорения выполнения(если нет угрозы race condition)
var hashFnv = fnv.New64()

type HashFunc func(key int, size int) int

type HashTableRehash struct {
	items      []*itemRehash
	count      int
	size       int
	hashFirst  HashFunc
	hashSecond HashFunc
}

func h1(key int, size int) int {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(key))
	hashFnv.Reset()
	_, _ = hashFnv.Write(buf)
	return int(hashFnv.Sum64() % uint64(size))
}

func h2(key int, size int) int {
	return (key % (size - 1)) + 1
}

func (ht *HashTableRehash) resizeTable() {
	newSize := ht.size * 2
	newTable := make([]*itemRehash, newSize)
	for i := 0; i < ht.size; i++ {
		if ht.items[i] != nil && !ht.items[i].deleted {
			indexFirst := ht.hashFirst(ht.items[i].key, newSize)
			indexSecond := ht.hashSecond(ht.items[i].key, newSize)
			for j := 0; j < newSize; j++ {
				if newTable[indexFirst] == nil {
					newTable[indexFirst] = ht.items[i]
					break
				}
				indexFirst = (indexFirst + indexSecond) % newSize
			}
		}
	}
	ht.size = newSize
	ht.items = newTable
}

// NewHashTableRehash нужно использовать простое число
func NewHashTableRehash(size int) *HashTableRehash {
	return &HashTableRehash{
		items:      make([]*itemRehash, size),
		size:       size,
		hashFirst:  h1,
		hashSecond: h2,
	}
}

func (ht *HashTableRehash) Put(key int, value string) {
	if float64(ht.count)/float64(ht.size) > 0.65 {
		ht.resizeTable()
	}
	indexFirst := ht.hashFirst(key, ht.size)
	indexSecond := ht.hashSecond(key, ht.size)
	for i := 0; i < ht.size; i++ {
		if ht.items[indexFirst] == nil || ht.items[indexFirst].deleted {
			ht.items[indexFirst] = &itemRehash{
				key:     key,
				value:   value,
				deleted: false,
			}
			return
		}
		indexFirst = (indexFirst + indexSecond) % ht.size
	}
}

func (ht *HashTableRehash) Get(key int) (string, error) {
	indexFirst := ht.hashFirst(key, ht.size)
	indexSecond := ht.hashSecond(key, ht.size)
	for i := 0; i < ht.size; i++ {
		if ht.items[indexFirst] != nil {
			if ht.items[indexFirst].key == key && !ht.items[indexFirst].deleted {
				return ht.items[indexFirst].value, nil
			}
		} else {
			return "", errors.New("not found")
		}
		indexFirst = (indexFirst + indexSecond) % ht.size
	}
	return "", errors.New("not found")
}

func (ht *HashTableRehash) Delete(key int) error {
	indexFirst := ht.hashFirst(key, ht.size)
	indexSecond := ht.hashSecond(key, ht.size)
	for i := 0; i < ht.size; i++ {
		if ht.items[indexFirst] != nil {
			if ht.items[indexFirst].key == key {
				ht.items[indexFirst].deleted = true
				return nil
			}
		} else {
			return errors.New("not found")
		}
		indexFirst = (indexFirst + indexSecond) % ht.size
	}
	return errors.New("not found")
}

func (ht *HashTableRehash) Update(key int, value string) error {
	indexFirst := ht.hashFirst(key, ht.size)
	indexSecond := ht.hashSecond(key, ht.size)
	for i := 0; i < ht.size; i++ {
		if ht.items[indexFirst] != nil {
			if ht.items[indexFirst].key == key && !ht.items[indexFirst].deleted {
				ht.items[indexFirst].value = value
				return nil
			}
		} else {
			return errors.New("not found")
		}
		indexFirst = (indexFirst + indexSecond) % ht.size
	}
	return errors.New("not found")
}

// способ разрешения коллизий - метод цепочек
type itemChain struct {
	key   int
	value string
	next  *itemChain
}

type HashTableChain struct {
	buckets []*itemChain
	size    int
	hash    hash.Hash64
}

// NewHashTableChain для лучшего распределения необходимо использовать в кач-ве size простое число
func NewHashTableChain(size int) *HashTableChain {
	return &HashTableChain{
		buckets: make([]*itemChain, size),
		size:    size,
		hash:    fnv.New64(),
	}
}

func (h *HashTableChain) hashIndex(key int) uint64 {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(key))
	h.hash.Reset()
	_, _ = h.hash.Write(buf)
	return h.hash.Sum64() % uint64(h.size)
}

func (h *HashTableChain) Put(key int, value string) {
	index := h.hashIndex(key)
	head := h.buckets[index]
	for current := head; current != nil; current = current.next {
		if current.key == key {
			current.value = value
			return
		}
	}

	newItem := &itemChain{key, value, head}
	h.buckets[index] = newItem
}

func (h *HashTableChain) Get(key int) (string, error) {
	index := h.hashIndex(key)
	head := h.buckets[index]

	for current := head; current != nil; current = current.next {
		if current.key == key {
			return current.value, nil
		}
	}

	return "", errors.New("not found")
}

func (h *HashTableChain) Delete(key int) error {
	index := h.hashIndex(key)
	head := h.buckets[index]

	if head == nil {
		return errors.New("not found")
	}

	if head.key == key {
		h.buckets[index] = head.next
		return nil
	}

	for prev, current := head, head.next; current != nil; prev, current = current, current.next {
		if current.key == key {
			prev.next = current.next
			return nil
		}
	}

	return errors.New("not found")
}

func (h *HashTableChain) Update(key int, value string) error {
	index := h.hashIndex(key)
	head := h.buckets[index]

	for current := head; current != nil; current = current.next {
		if current.key == key {
			current.value = value
			return nil
		}
	}

	return errors.New("not found")
}
