package binomial

import (
	"errors"
)

//二项堆的Push和Top操作的复杂度为O(1)，其余核心操作复杂度为O(logN)。
type Heap struct {
	list *node
	top  *node
}
type node struct {
	key   int
	level uint
	peer  *node
	child *node
}

func (hp *Heap) IsEmpty() bool {
	return hp.list == nil
}
func (hp *Heap) Clear() {
	hp.list, hp.top = nil, nil
}

func (hp *Heap) Top() (int, error) {
	if hp.IsEmpty() {
		return 0, errors.New("empty")
	}
	return hp.top.key, nil
}
func (hp *Heap) Push(key int) {
	var unit = new(node)
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
func reverse(list *node) *node {
	var head = (*node)(nil)
	for list != nil {
		var current = list
		list = list.peer
		current.peer, head = head, current
	}
	return head
}
func (hp *Heap) Pop() (int, error) {
	if hp.IsEmpty() {
		return 0, errors.New("empty")
	}
	var key = hp.top.key

	var knot = fakeHead(&hp.list)
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
	return key, nil
}
