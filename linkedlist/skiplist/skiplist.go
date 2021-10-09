package skiplist

import (
	"DSGO/utils"
	"constraints"
	"math"
	"time"
	"unsafe"
)

type SkipList[T constraints.Ordered] interface {
	Size() int
	IsEmpty() bool
	Clear()
	Insert(T) bool
	Search(T) bool
	Remove(T) bool
	Travel(func(T))
}

const LevelFactor = 3

type bNode[T constraints.Ordered] struct {
	next []*bNode[T]
}

type lNode[T constraints.Ordered] struct {
	bNode[T]
	key T
}

func (n *bNode[T]) asLeaf() *lNode[T] {
	return (*lNode[T])(unsafe.Pointer(n))
}

type skipList[T constraints.Ordered] struct {
	bNode[T]    //head
	size        int
	floor, ceil int    //非零
	magic       uint32 //随机状态
	knots       []*bNode[T]
}

func (l *skipList[T]) rand() uint32 {
	l.magic = l.magic*1103515245 + 12345
	return l.magic
}

func New[T constraints.Ordered]() SkipList[T] {
	l := new(skipList[T])
	l.init()
	return l
}

func (l *skipList[T]) init() {
	l.magic = uint32(time.Now().UnixNano())
	l.next = make([]*bNode[T], 1)
	l.size = 0
	l.floor, l.ceil = 1, LevelFactor
	l.knots = make([]*bNode[T], 1)
}

func (l *skipList[T]) Clear() {
	l.next = l.next[:1]
	l.next[0] = nil
	l.size = 0
	l.floor, l.ceil = 1, LevelFactor
	l.knots = l.knots[:1]
	l.knots = nil
}

func (l *skipList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *skipList[T]) Size() int {
	return l.size
}

func (l *skipList[T]) Travel(doit func(T)) {
	for node := l.next[0]; node != nil; node = node.next[0] {
		doit(node.asLeaf().key)
	}
}

func (l *skipList[T]) Search(key T) bool {
	knot, node := &l.bNode, (*lNode[T])(nil)
	for i := len(l.next) - 1; i >= 0; i-- {
		for node = knot.next[i].asLeaf(); node != nil &&
			node.key < key; node = knot.next[i].asLeaf() {
			knot = &node.bNode
		}
	}
	return node != nil && node.key == key
}

func (l *skipList[T]) search(key T) *lNode[T] {
	knot, node := &l.bNode, (*lNode[T])(nil)
	for i := len(l.next) - 1; i >= 0; i-- {
		for node = knot.next[i].asLeaf(); node != nil &&
			node.key < key; node = knot.next[i].asLeaf() {
			knot = &node.bNode
		}
		l.knots[i] = knot
	}
	return node
}

//成功返回true，冲突返回false
func (l *skipList[T]) Insert(key T) bool {
	node := l.search(key)
	if node != nil && node.key == key {
		return false
	}

	level := 1
	for level < len(l.next) &&
		l.rand() <= (math.MaxUint32/uint32(LevelFactor)) {
		level++
	}
	node = newLeaf[T](level)
	node.key = key
	for i := 0; i < level; i++ {
		node.next[i] = l.knots[i].next[i]
		l.knots[i].next[i] = &node.bNode
	}

	l.size++
	if l.size == l.ceil {
		l.floor = l.ceil
		l.ceil *= LevelFactor
		l.next = append(l.next, nil)
		l.knots = append(l.knots, nil)
	}
	return true
}

//成功返回true，没有返回false
func (l *skipList[T]) Remove(key T) bool {
	node := l.search(key)
	if node == nil || node.key != key {
		return false
	}

	level := utils.Min(len(l.knots), len(node.next))
	for i := 0; i < level; i++ {
		l.knots[i].next[i] = node.next[i]
	}

	l.size--
	if l.size < l.floor && len(l.next) > 1 {
		l.ceil = l.floor
		l.floor /= LevelFactor
		last := len(l.next) - 1
		for knot := &l.bNode; knot.next[last] != nil; {
			knot, knot.next[last] = knot.next[last], nil
		} //此处不清理恐有内存泄漏
		l.next = l.next[:last]
		l.knots = l.knots[:last]
	}
	return true
}
