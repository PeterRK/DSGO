package binomial

import (
	"constraints"
	"unsafe"
)

//二项堆的Push和Top操作的复杂度为O(1)，其余核心操作复杂度为O(logN)。
type Heap[T constraints.Ordered] struct {
	list *node[T]
	top  *node[T]
	size int
}

type node[T constraints.Ordered] struct {
	key   T
	level uint8
	peer  *node[T]
	child *node[T]
}

func (hp *Heap[T]) Size() int {
	return hp.size
}

func (hp *Heap[T]) IsEmpty() bool {
	return hp.list == nil
}

func (hp *Heap[T]) Clear() {
	hp.list, hp.top = nil, nil
	hp.size = 0
}

func (hp *Heap[T]) Top() T {
	if hp.IsEmpty() {
		panic("empty heap")
	}
	return hp.top.key
}

func (hp *Heap[T]) Push(key T) {
	hp.size++
	unit := new(node[T])
	unit.key, unit.level = key, 0
	unit.peer, unit.child = nil, nil
	if hp.IsEmpty() {
		hp.list, hp.top = unit, unit
	} else {
		if key < hp.top.key {
			hp.top = unit
		}
		hp.merge(unit)
	}
}

//list是从少到多的，而child相反
func reverse[T constraints.Ordered](list *node[T]) *node[T] {
	head := (*node[T])(nil)
	for list != nil {
		unit := list
		list = list.peer
		unit.peer, head = head, unit
	}
	return head
}

func (hp *Heap[T]) Pop() T {
	if hp.IsEmpty() {
		panic("empty heap")
	}
	hp.size--
	key := hp.top.key

	knot := fakeHead(&hp.list)
	for knot.peer != hp.top {
		knot = knot.peer
	}
	knot.peer = knot.peer.peer

	hp.merge(reverse(hp.top.child))
	hp.top = hp.list
	if hp.list != nil {
		for pt := hp.list.peer; pt != nil; pt = pt.peer {
			if pt.key < hp.top.key {
				hp.top = pt
			}
		}
	}
	return key
}

func fakeHead[T constraints.Ordered](spt **node[T]) *node[T] {
	base := uintptr(unsafe.Pointer(spt))
	off := unsafe.Offsetof((*spt).peer)
	return (*node[T])(unsafe.Pointer(base - off))
}

//list是从少到多的
func (hp *Heap[T]) merge(list *node[T]) {
	knot := fakeHead(&hp.list)
	for list != nil {
		a, b := list, knot.peer
		if b == nil || a.level < b.level {
			list, a.peer = a.peer, b
			knot.peer = a
		} else if a.level > b.level {
			knot = knot.peer
		} else { //同级合并
			list, knot.peer = a.peer, b.peer

			if a.key > b.key {
				a, b = b, a
			}
			b.peer, a.child = a.child, b
			a.level++

			a.peer, list = list, a //可能会有一项逆序，不影响大局
		}
	}
}

func (hp *Heap[T]) Merge(other *Heap[T]) {
	if hp != other && !other.IsEmpty() {
		if hp.IsEmpty() {
			*hp = *other
		} else {
			if hp.top.key > other.top.key {
				hp.top = other.top
			}
			hp.merge(other.list)
			hp.size += other.size
		}
		other.Clear()
	}
}
