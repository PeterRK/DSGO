package cuckoo

import (
	"HashTable/hash"
)

type node struct {
	code [2]uint
	key  string
}
type table struct {
	fn     func(str string) uint
	mask   uint
	id     uint8 //0或1
	bucket []*node
}
type hashTable struct {
	a, b table
	cnt  int
}

func (tb *hashTable) Size() int {
	return tb.cnt
}
func (tb *hashTable) IsEmpty() bool {
	return tb.cnt == 0
}

func (tb *hashTable) initialize(fn1 func(str string) uint, fn2 func(str string) uint) {
	tb.cnt = 0
	tb.a.fn, tb.b.fn = fn1, fn2
	tb.a.id, tb.b.id = 0, 1
	tb.a.bucket, tb.b.bucket = make([]*node, 32), make([]*node, 16) //保持a为b的两倍
	tb.a.mask, tb.b.mask = 0x1f, 0xf
}
func NewHashTable(fn1 func(str string) uint, fn2 func(str string) uint) hash.HashTable {
	var tb = new(hashTable)
	tb.initialize(fn1, fn2)
	return tb
}

func (tb *table) hash(key string) uint {
	var code = tb.fn(key)
	for m := code >> 4; m != 0; m >>= 4 {
		code ^= m
	} //这里的表大小非质数，code后面数位的价值更大
	return code
}
