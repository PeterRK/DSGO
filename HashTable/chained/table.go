package chained

import (
	"HashTable/hash"
)

type node struct {
	key  []byte
	next *node
}
type hashTable struct {
	hash   func(str []byte) uint
	bucket []*node
	cnt    int
}

func (tb *hashTable) Size() int {
	return tb.cnt
}
func (tb *hashTable) IsEmpty() bool {
	return tb.cnt == 0
}
func (tb *hashTable) isCrowded() bool {
	return tb.cnt*2 > len(tb.bucket)*3
}
func (tb *hashTable) isWasteful() bool {
	return tb.cnt*10 < len(tb.bucket)
}

func (tb *hashTable) initialize(fn func(str []byte) uint) {
	tb.cnt, tb.hash = 0, fn
	tb.bucket = make([]*node, primes[0])
}
func NewHashTable(fn func(str []byte) uint) hash.HashTable {
	var tb = new(hashTable)
	tb.initialize(fn)
	return tb
}

//此数列借鉴自SGI STL
var primes = []uint{
	17, 29, 53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593, 49157, 98317, 196613,
	393241, 786433, 1572869, 3145739, 6291469, 12582917, 25165843, 50331653, 1610612741}

func (tb *hashTable) expand() {
	var size = uint(len(tb.bucket))
	if size != primes[len(primes)-1] {
		var a, b = 0, len(primes)
		for a < b {
			var m = (a + b) / 2
			if size < primes[m] {
				b = m
			} else {
				a = m + 1
			}
		}
		tb.resize(primes[a])
	}
}
func (tb *hashTable) shrink() {
	var size = uint(len(tb.bucket))
	if size != primes[0] {
		var a, b = len(primes) - 1, -1
		for a > b {
			var m = (a + b + 1) / 2
			if size > primes[m] {
				b = m
			} else {
				a = m - 1
			}
		}
		tb.resize(primes[a])
	}
}
