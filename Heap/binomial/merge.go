package binomial

import (
	"unsafe"
)

func fakeHead(spt **node) *node {
	base := uintptr(unsafe.Pointer(spt))
	off := unsafe.Offsetof((*spt).peer)
	return (*node)(unsafe.Pointer(base - off))
}

//list是从少到多的
func (hp *Heap) merge(list *node) {
	knot := fakeHead(&hp.list)
	for list != nil {
		one, another := list, knot.peer
		if another == nil || one.level < another.level {
			list, one.peer = one.peer, another
			knot.peer = one
		} else if one.level > another.level {
			knot = knot.peer
		} else { //同级合并
			list, knot.peer = one.peer, another.peer

			if one.key > another.key {
				one, another = another, one
			}
			another.peer, one.child = one.child, another
			one.level++

			one.peer, list = list, one //可能会有一项逆序，不影响大局
		}
	}
}
func (hp *Heap) Merge(victim *Heap) {
	if hp != victim && !victim.IsEmpty() {
		if hp.IsEmpty() {
			*hp = *victim
		} else {
			if hp.top.key > victim.top.key {
				hp.top = victim.top
			}
			hp.merge(victim.list)
		}
		victim.Clear()
	}
}
