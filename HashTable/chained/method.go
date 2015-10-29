package chained

import (
	"bytes"
	"unsafe"
)

func (tb *hashTable) Search(key []byte) bool {
	var index = tb.hash(key) % uint(len(tb.bucket))
	return search(tb.bucket[index], key)
}
func search(head *node, key []byte) bool {
	for ; head != nil; head = head.next {
		if bytes.Compare(key, head.key) == 0 {
			return true
		}
	}
	return false
}

func (tb *hashTable) Remove(key []byte) bool {
	var index = tb.hash(key) % uint(len(tb.bucket))
	var list = tb.bucket[index] //直接取对GC不友好，绕一下道
	for knot := fakeHead(&list); knot.next != nil; knot = knot.next {
		if bytes.Compare(key, knot.next.key) == 0 {
			knot.next = knot.next.next
			tb.bucket[index] = list
			tb.cnt--
			if tb.isWasteful() {
				tb.shrink()
			}
			return true
		}
	}
	return false
}
func fakeHead(spt **node) *node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).next)
	return (*node)(unsafe.Pointer(base - off))
}

func (tb *hashTable) Insert(key []byte) bool {
	var index = tb.hash(key) % uint(len(tb.bucket))
	if search(tb.bucket[index], key) {
		return false
	}
	var unit = new(node)
	unit.key = key
	unit.next, tb.bucket[index] = tb.bucket[index], unit

	tb.cnt++
	if tb.isCrowded() {
		tb.expand()
	}
	return true
}

func (tb *hashTable) resize(size uint) {
	var old_bucket = tb.bucket
	tb.bucket = make([]*node, size)
	for _, unit := range old_bucket {
		for unit != nil {
			var current, index = unit, tb.hash(unit.key) % size
			unit = unit.next
			current.next, tb.bucket[index] = tb.bucket[index], current
		}
	}
}
