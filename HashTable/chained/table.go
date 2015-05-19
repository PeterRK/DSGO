package hashtable

import (
	"unsafe"
)

type node struct {
	key  string
	next *node
}

func fakeHead(this **node) *node {
	var base = uintptr(unsafe.Pointer(this))
	var off = unsafe.Offsetof((*this).next)
	return (*node)(unsafe.Pointer(base - off))
}

type hashTable struct {
	hash   func(str string) uint
	bucket []*node
	cnt    int
}
type HashTable interface {
	IsEmpty() bool
	Size() int
	Insert(key string) bool
	Search(key string) bool
	Remove(key string) bool
}

func NewHashTable(hash func(str string) uint) HashTable {
	var table = new(hashTable)
	table.cnt = 0
	table.hash = hash
	table.bucket = make([]*node, initSize())
	return table
}

func (table *hashTable) Size() int {
	return table.cnt
}
func (table *hashTable) IsEmpty() bool {
	return table.cnt == 0
}
func (table *hashTable) isCrowded() bool {
	return table.cnt*2 > len(table.bucket)*3
}

func (table *hashTable) Search(key string) bool {
	var index = table.hash(key) % uint(len(table.bucket))
	for unit := table.bucket[index]; unit != nil; unit = unit.next {
		if key == unit.key {
			return true
		}
	}
	return false
}

//成功返回true，没有返回false
func (table *hashTable) Remove(key string) bool {
	var index = table.hash(key) % uint(len(table.bucket))
	for knot := fakeHead(&table.bucket[index]); knot.next != nil; knot = knot.next {
		if key == knot.next.key {
			knot.next = knot.next.next
			table.cnt--
			return true
		}
	}
	return false
}
