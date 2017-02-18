package chained

import (
	"bytes"
)

func (tb *hashTable) Search(key []byte) bool {
	var code = tb.hash(key)
	var index = code % uint32(len(tb.bucket))
	var found = search(tb.bucket[index], key)
	if tb.isMoving() {
		if !found { //尝试从旧表中查找
			index = code % uint32(len(tb.old_bucket))
			found = search(tb.old_bucket[index], key)
		}
		tb.moveLine() //推进rehash过程
	}
	return found
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
	var code, done = tb.hash(key), false
	var index = code % uint32(len(tb.bucket))
	tb.bucket[index], done = remove(tb.bucket[index], key)
	if tb.isMoving() {
		if !done { //尝试从旧表中删除
			index = code % uint32(len(tb.old_bucket))
			tb.old_bucket[index], done = remove(tb.old_bucket[index], key)
		}
		tb.moveLine()
	}
	if done {
		tb.cnt--
		if !tb.isMoving() && tb.isWasteful() {
			tb.shrink()
		}
	}
	return done
}
func remove(head *node, key []byte) (*node, bool) {
	for knot := fakeHead(&head); knot.next != nil; knot = knot.next {
		if bytes.Compare(key, knot.next.key) == 0 {
			knot.next = knot.next.next
			return head, true
		}
	}
	return head, false
}

func (tb *hashTable) Insert(key []byte) bool {
	var code = tb.hash(key)
	var index = code % uint32(len(tb.bucket))
	var conflict = search(tb.bucket[index], key)
	if tb.isMoving() {
		if !conflict {
			var index = code % uint32(len(tb.old_bucket))
			conflict = search(tb.old_bucket[index], key)
		}
		tb.moveLine()
	}
	if !conflict {
		var unit = new(node)
		unit.key = key
		unit.next, tb.bucket[index] = tb.bucket[index], unit
		tb.cnt++
		if !tb.isMoving() && tb.isCrowded() {
			tb.expand()
		}
		return true
	}
	return false
}

func (tb *hashTable) resize(size uint) { //size != 0
	tb.old_bucket, tb.bucket = tb.bucket, make([]*node, size)
}
func (tb *hashTable) moveLine() {
	var size = uint32(len(tb.bucket))
	for head := tb.old_bucket[tb.next_line]; head != nil; {
		var unit, index = head, tb.hash(head.key) % size
		head = head.next
		unit.next, tb.bucket[index] = tb.bucket[index], unit
	}
	tb.old_bucket[tb.next_line] = nil
	tb.next_line++
	if tb.next_line == len(tb.old_bucket) {
		tb.stopMoving() //rehash完成
	}
}
