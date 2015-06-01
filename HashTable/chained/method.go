package chained

import (
	"unsafe"
)

func (tb *hashTable) Search(key string) bool {
	var index = tb.hash(key) % uint(len(tb.bucket))
	for unit := tb.bucket[index]; unit != nil; unit = unit.next {
		if key == unit.key {
			return true
		}
	}
	return false
}

//成功返回true，没有返回false
func (tb *hashTable) Remove(key string) bool {
	var index = tb.hash(key) % uint(len(tb.bucket))
	for knot := fakeHead(&tb.bucket[index]); knot.next != nil; knot = knot.next {
		if key == knot.next.key {
			knot.next = knot.next.next
			tb.cnt--
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

//成功返回true，冲突返回false
func (tb *hashTable) Insert(key string) bool {
	var index = tb.hash(key) % uint(len(tb.bucket))
	for unit := tb.bucket[index]; unit != nil; unit = unit.next {
		if key == unit.key {
			return false
		}
	}
	var unit = new(node)
	unit.key = key
	unit.next, tb.bucket[index] = tb.bucket[index], unit

	tb.cnt++
	if tb.isCrowded() {
		if newsz, ok := nextSize(uint(len(tb.bucket))); ok {
			tb.resize(newsz)
		}
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
