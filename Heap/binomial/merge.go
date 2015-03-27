package binomial

import (
	"unsafe"
)

func fakeHead(spt **node) *node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).peer)
	return (*node)(unsafe.Pointer(base - off))
}

//list是从少到多的
func (heap *Heap) merge(list *node) {
	var knot = fakeHead(&heap.list)
	for knot.peer != nil && list != nil {
		if knot.peer.level == list.level {
			var root = knot.peer
			knot.peer = root.peer
			var another = list
			list = another.peer

			if root.key > another.key {
				root, another = another, root
			}
			another.peer, root.child = root.child, another
			root.level++

			root.peer, list = list, root
		} else {
			if knot.peer.level > list.level {
				var target = list
				list = list.peer
				target.peer, knot.peer = knot.peer, target
			}
			knot = knot.peer
		}
	}
	if list != nil {
		knot.peer = list
	}
}
func (heap *Heap) Merge(victim *Heap) {
	if heap != victim {
		if heap.top.key > victim.top.key {
			heap.top = victim.top
		}
		heap.merge(victim.list)
		victim.list, victim.top = nil, nil
	}
}
