package binomialheap

import (
	"unsafe"
)

func fakeHead(this **node) *node {
	var base = uintptr(unsafe.Pointer(this))
	var off = unsafe.Offsetof((*this).brother)
	return (*node)(unsafe.Pointer(base - off))
}

func (heap *Heap) merge(list *node) {
	var knot = fakeHead(&heap.list)
	for knot.brother != nil && list != nil {
		if knot.brother.height == list.height {
			var root = knot.brother
			knot.brother = root.brother
			var another = list
			list = another.brother

			if root.key > another.key {
				root, another = another, root
			}
			another.brother, root.child = root.child, another
			root.height++

			root.brother = list
			list = root
		} else {
			if knot.brother.height > list.height {
				var target = list
				list = list.brother
				target.brother = knot.brother
				knot.brother = target
			}
			knot = knot.brother
		}
	}
	if list != nil {
		knot.brother = list
	}
}
func (heap *Heap) Merge(victim *Heap) {
	if heap != victim {
		heap.merge(victim.list)
		victim.list = nil
	}
}
