package skiplist

import (
	"DSGO/utils"
	"constraints"
	"time"
	"unsafe"
)

type SkipList[T constraints.Ordered] interface {
	Size() int
	IsEmpty() bool
	Insert(T) bool
	Search(T) bool
	Remove(T) bool
	Travel(func(T))
}

const LevelFactor = 3

type node[T constraints.Ordered] struct {
	next []*node[T]
	key  T
}

type skipList[T constraints.Ordered] struct {
	heads, knots []*node[T]
	size, level  int //非零
	ceil, floor  int //非零
	rand         utils.Random
}

func NewSkipList[T constraints.Ordered]() SkipList[T] {
	l := new(skipList[T])
	l.init()
	return l
}

func (l *skipList[T]) init() {
	l.rand = utils.NewXorshift(uint32(time.Now().Unix()))
	l.heads, l.knots = make([]*node[T], 1), make([]*node[T], 1)
	l.level, l.size, l.ceil, l.floor = 1, 1, LevelFactor, 1
}

func (l *skipList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *skipList[T]) Size() int {
	return l.size - 1
}

func (l *skipList[T]) Travel(doit func(T)) {
	for unit := l.heads[0]; unit != nil; unit = unit.next[0] {
		doit(unit.key)
	}
}

func (l *skipList[T]) Search(key T) bool {
	knot := (*node[T])(unsafe.Pointer(l))
	for i := l.level - 1; i >= 0; i-- {
		for knot.next[i] != nil && knot.next[i].key < key {
			knot = knot.next[i]
		}
	}
	target := knot.next[0]
	return target != nil && target.key == key
}

//成功返回true，冲突返回false
func (l *skipList[T]) Insert(key T) bool {
	knot := (*node[T])(unsafe.Pointer(l))
	for i := l.level - 1; i >= 0; i-- {
		for knot.next[i] != nil && knot.next[i].key < key {
			knot = knot.next[i]
		}
		l.knots[i] = knot
	}
	target := knot.next[0]
	if target != nil && target.key == key {
		return false
	}

	l.size++
	if l.size == l.ceil {
		l.floor = l.ceil
		l.ceil *= LevelFactor
		l.level++
		l.heads = append(l.heads, nil)
		l.knots = append(l.knots, (*node[T])(unsafe.Pointer(l)))
	}

	lv := 1
	for lv < l.level &&
		l.rand.Next() <= (^uint32(0)/uint32(LevelFactor)) {
		lv++
	}
	target = new(node[T])
	target.key = key
	target.next = make([]*node[T], lv)
	for i := 0; i < lv; i++ {
		target.next[i] = l.knots[i].next[i]
		l.knots[i].next[i] = target
	}
	return true
}

//成功返回true，没有返回false
func (l *skipList[T]) Remove(key T) bool {
	knot := (*node[T])(unsafe.Pointer(l))
	for i := l.level - 1; i >= 0; i-- {
		for knot.next[i] != nil && knot.next[i].key < key {
			knot = knot.next[i]
		}
		l.knots[i] = knot
	}
	target := knot.next[0]
	if target == nil || target.key != key {
		return false
	}

	lv := utils.Min(len(target.next), l.level)
	for i := 0; i < lv; i++ {
		l.knots[i].next[i] = target.next[i]
	}

	l.size--
	if l.size < l.floor { //注意不能==
		l.ceil = l.floor
		l.floor /= LevelFactor
		l.level--
		l.heads = l.heads[:l.level]
		l.knots = l.knots[:l.level]
	}
	return true
}
