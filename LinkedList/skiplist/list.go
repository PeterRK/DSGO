package skiplist

import (
	"time"
	"unsafe"
)

type SkipList interface {
	Size() int
	IsEmpty() bool
	Insert(key int) bool
	Search(key int) bool
	Remove(key int) bool
	Travel(doit func(int))
}

const LEVEL_FACTOR = 3

type node struct {
	next []*node
	key  int
}
type skipList struct {
	heads, knots []*node
	cnt, level   int //非零
	ceil, floor  int //非零
	rand         Random
}

func NewSkipList() SkipList {
	var l = new(skipList)
	l.initialize()
	return l
}

func (l *skipList) initialize() {
	l.rand = NewEasyRand(uint(time.Now().Unix()))
	l.heads, l.knots = make([]*node, 1), make([]*node, 1)
	l.level, l.cnt, l.ceil, l.floor = 1, 1, LEVEL_FACTOR, 1
}

func (l *skipList) IsEmpty() bool {
	return l.Size() == 0
}
func (l *skipList) Size() int {
	return l.cnt - 1
}

func (l *skipList) Travel(doit func(int)) {
	for unit := l.heads[0]; unit != nil; unit = unit.next[0] {
		doit(unit.key)
	}
}
func (l *skipList) Search(key int) bool {
	var knot = (*node)(unsafe.Pointer(l))
	for i := l.level - 1; i >= 0; i-- {
		for knot.next[i] != nil && knot.next[i].key < key {
			knot = knot.next[i]
		}
	}
	var target = knot.next[0]
	return target != nil && target.key == key
}

//成功返回true，冲突返回false
func (l *skipList) Insert(key int) bool {
	var knot = (*node)(unsafe.Pointer(l))
	for i := l.level - 1; i >= 0; i-- {
		for knot.next[i] != nil && knot.next[i].key < key {
			knot = knot.next[i]
		}
		l.knots[i] = knot
	}
	var target = knot.next[0]
	if target != nil && target.key == key {
		return false
	}

	l.cnt++
	if l.cnt == l.ceil {
		l.floor = l.ceil
		l.ceil *= LEVEL_FACTOR
		l.level++
		l.heads = append(l.heads, nil)
		l.knots = append(l.knots, (*node)(unsafe.Pointer(l)))
	}

	var lv = 1
	for lv < l.level &&
		l.rand.Next() <= (^uint32(0)/uint32(LEVEL_FACTOR)) {
		lv++
	}
	target = new(node)
	target.key = key
	target.next = make([]*node, lv)
	for i := 0; i < lv; i++ {
		target.next[i] = l.knots[i].next[i]
		l.knots[i].next[i] = target
	}
	return true
}

//成功返回true，没有返回false
func (l *skipList) Remove(key int) bool {
	var knot = (*node)(unsafe.Pointer(l))
	for i := l.level - 1; i >= 0; i-- {
		for knot.next[i] != nil && knot.next[i].key < key {
			knot = knot.next[i]
		}
		l.knots[i] = knot
	}
	var target = knot.next[0]
	if target == nil || target.key != key {
		return false
	}

	var lv = min(len(target.next), l.level)
	for i := 0; i < lv; i++ {
		l.knots[i].next[i] = target.next[i]
	}

	l.cnt--
	if l.cnt < l.floor { //注意不能==
		l.ceil = l.floor
		l.floor /= LEVEL_FACTOR
		l.level--
		l.heads = l.heads[:l.level]
		l.knots = l.knots[:l.level]
	}
	return true
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
