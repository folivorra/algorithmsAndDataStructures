package main

import (
	"encoding/binary"
	"errors"
	"hash"
	"hash/fnv"
)

// способ разрешения коллизий - метод цепочек

type item struct {
	key   int
	value string
	next  *item
}

type bucket *item

type HashTable struct {
	buckets  []bucket
	length   int
	capacity int
	hash     hash.Hash64
}

// NewHashTable : для лучшего распределения необходимо использовать в кач-ве size простое число
func NewHashTable(cap int) *HashTable {
	return &HashTable{
		buckets:  make([]bucket, cap),
		capacity: cap,
		hash:     fnv.New64(),
	}
}

func (h *HashTable) hashIndex(key int) uint64 {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(key))
	h.hash.Reset()
	h.hash.Write(buf)
	return h.hash.Sum64() % uint64(h.capacity)
}

func (h *HashTable) Put(key int, value string) {
	index := h.hashIndex(key)
	head := h.buckets[index]
	for current := head; current != nil; current = current.next {
		if current.key == key {
			current.value = value
			return
		}
	}

	newItem := &item{key, value, head}
	h.buckets[index] = newItem
	h.length++
}

func (h *HashTable) Get(key int) (string, error) {
	index := h.hashIndex(key)
	head := h.buckets[index]

	for current := head; current != nil; current = current.next {
		if current.key == key {
			return current.value, nil
		}
	}

	return "", errors.New("not found")
}

func (h *HashTable) Delete(key int) error {
	index := h.hashIndex(key)
	head := h.buckets[index]

	if head == nil {
		return errors.New("not found")
	}

	if head.key == key {
		h.buckets[index] = head.next
		h.length--
		return nil
	}

	for prev, current := head, head.next; current != nil; prev, current = current, current.next {
		if current.key == key {
			prev.next = current.next
			h.length--
			return nil
		}
	}

	return errors.New("not found")
}

func (h *HashTable) Update(key int, value string) error {
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
