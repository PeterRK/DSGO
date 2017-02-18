package chained

import (
	"DSGO/HashTable/hash"
	"unsafe"
)

type node struct {
	key  []byte
	next *node
}

func fakeHead(spt **node) *node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).next)
	return (*node)(unsafe.Pointer(base - off))
}

type hashTable struct {
	hash       func(str []byte) uint32
	bucket     []*node
	cnt        int
	old_bucket []*node //旧表（仅在rehash过程中有内容）
	next_line  int     //标记待处理的旧表行
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
func (tb *hashTable) isMoving() bool {
	return len(tb.old_bucket) != 0
}
func (tb *hashTable) stopMoving() {
	tb.next_line = 0
	tb.old_bucket = nil //GC
}

func (tb *hashTable) initialize(fn func(str []byte) uint32) {
	tb.hash = fn
	//tb.cnt, tb.next_line = 0, 0
	tb.bucket = make([]*node, primes[0])
}
func NewHashTable(fn func(str []byte) uint32) hash.HashTable {
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
