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

const factor = 3

type node struct {
	next []*node
	key  int
}
type skipList struct {
	heads []*node
	knots []*node
	level int //非零
	mark  int //非零
	ceil  int //非零
	floor int //非零
	rand  mt19937
}

func NewSkipList() SkipList {
	var list = new(skipList)
	list.initialize()
	return list
}

func (list *skipList) initialize() {
	list.rand.initialize(uint32(time.Now().Unix()))
	list.heads, list.knots = make([]*node, 1), make([]*node, 1)
	list.level, list.mark, list.ceil, list.floor = 1, 1, factor, 1
}

func (list *skipList) IsEmpty() bool {
	return list.Size() == 0
}
func (list *skipList) Size() int {
	return list.mark - 1
}

func (list *skipList) Travel(doit func(int)) {
	for unit := list.heads[0]; unit != nil; unit = unit.next[0] {
		doit(unit.key)
	}
}
func (list *skipList) Search(key int) bool {
	var knot = (*node)(unsafe.Pointer(list))
	for i := list.level - 1; i >= 0; i-- {
		for knot.next[i] != nil && knot.next[i].key < key {
			knot = knot.next[i]
		}
	}
	var target = knot.next[0]
	return target != nil && target.key == key
}

//成功返回true，冲突返回false
func (list *skipList) Insert(key int) bool {
	var knot = (*node)(unsafe.Pointer(list))
	for i := list.level - 1; i >= 0; i-- {
		for knot.next[i] != nil && knot.next[i].key < key {
			knot = knot.next[i]
		}
		list.knots[i] = knot
	}
	var target = knot.next[0]
	if target != nil && target.key == key {
		return false
	}

	list.mark++
	if list.mark == list.ceil {
		list.floor = list.ceil
		list.ceil *= factor
		list.level++
		list.heads = append(list.heads, nil)
		list.knots = append(list.knots, (*node)(unsafe.Pointer(list)))
	}

	var lv = 1
	for list.rand.next() <= (^uint32(0)/uint32(factor)) &&
		lv < list.level {
		lv++
	}
	target = new(node)
	target.key = key
	target.next = make([]*node, lv)
	for i := 0; i < lv; i++ {
		target.next[i] = list.knots[i].next[i]
		list.knots[i].next[i] = target
	}
	return true
}

//成功返回true，没有返回false
func (list *skipList) Remove(key int) bool {
	var knot = (*node)(unsafe.Pointer(list))
	for i := list.level - 1; i >= 0; i-- {
		for knot.next[i] != nil && knot.next[i].key < key {
			knot = knot.next[i]
		}
		list.knots[i] = knot
	}
	var target = knot.next[0]
	if target == nil || target.key != key {
		return false
	}

	var lv = min(len(target.next), list.level)
	for i := 0; i < lv; i++ {
		list.knots[i].next[i] = target.next[i]
	}

	list.mark--
	if list.mark < list.floor { //注意不能==
		list.ceil = list.floor
		list.floor /= factor
		list.level--
		list.heads = list.heads[:list.level]
		list.knots = list.knots[:list.level]
	}
	return true
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
