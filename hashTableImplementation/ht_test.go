package main

import "testing"

func BenchmarkHashTableChain(b *testing.B) {
	ht := NewHashTableChain(3571)
	for i := 0; i < b.N; i++ {
		ht.Put(i, "value")
		_, _ = ht.Get(i)
		_ = ht.Update(i, "updated")
		_ = ht.Delete(i)
	}
}

func BenchmarkHashTableRehash(b *testing.B) {
	ht := NewHashTableRehash(3571)
	for i := 0; i < b.N; i++ {
		ht.Put(i, "value")
		_, _ = ht.Get(i)
		_ = ht.Update(i, "updated")
		_ = ht.Delete(i)
	}
}
