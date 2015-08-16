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

func nextSize(size uint) (newsz uint, ok bool) {
	var start, end = 0, len(primes)
	for start < end {
		var mid = (start + end) / 2
		if size < primes[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if start == len(primes) {
		return size, false
	}
	return primes[start], true
}
