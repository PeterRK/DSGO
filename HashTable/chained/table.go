package chained

import (
	"unsafe"
)

//此数列借鉴自SGI STL
var size_primes = []uint{
	53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593, 49157, 98317, 196613,
	393241, 786433, 1572869, 3145739, 6291469, 12582917, 25165843, 50331653, 1610612741}

func nextSize(size uint) (newsz uint, ok bool) {
	var start, end = 0, len(size_primes)
	for start < end {
		var mid = (start + end) / 2
		if size < size_primes[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if start == len(size_primes) {
		return size, false
	}
	return size_primes[start], true
}

type node struct {
	key  string
	next *node
}

func fakeHead(this **node) *node {
	var base = uintptr(unsafe.Pointer(this))
	var off = unsafe.Offsetof((*this).next)
	return (*node)(unsafe.Pointer(base - off))
}

type HashTable struct {
	hash   func(str string) uint
	bucket []*node
	cnt    int
}

func (table *HashTable) Size() int {
	return table.cnt
}
func (table *HashTable) IsEmpty() bool {
	return table.cnt == 0
}
func (table *HashTable) isCrowded() bool {
	return table.cnt*2 > len(table.bucket)*3
}

func (table *HashTable) Initialize(fn func(str string) uint) {
	table.cnt = 0
	table.hash = fn
	table.bucket = make([]*node, size_primes[0])
}
