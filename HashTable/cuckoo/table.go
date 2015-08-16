package cuckoo

import (
	"HashTable/hash"
)

const WAYS = 3

type node struct {
	code [WAYS]uint
	key  []byte
}
type table struct {
	hash   func([]byte) uint
	bucket []*node
}
type hashTable struct {
	core [WAYS]table
	idx  int
	cnt  int
}

func (tb *hashTable) Size() int {
	return tb.cnt
}
func (tb *hashTable) IsEmpty() bool {
	return tb.cnt == 0
}

func (tb *hashTable) initialize(fn [WAYS]func(str []byte) uint) {
	tb.idx, tb.cnt = 0, 0
	var sz = 8 //2^n
	for i := WAYS - 1; i >= 0; i-- {
		tb.core[i].hash = fn[i]
		tb.core[i].bucket = make([]*node, sz)
		sz *= 2
	}
}
func NewHashTable(fn [WAYS]func(str []byte) uint) hash.HashTable {
	var tb = new(hashTable)
	tb.initialize(fn)
	return tb
}
