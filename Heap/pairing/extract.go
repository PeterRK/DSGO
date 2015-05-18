package pairingheap

import (
	"unsafe"
)

func fakeHead(this **Node) *Node {
	var base = uintptr(unsafe.Pointer(this))
	var off = unsafe.Offsetof((*this).brother)
	return (*Node)(unsafe.Pointer(base - off))
}
func (heap *Heap) PopNode() (unit *Node) {
	if heap.root == nil {
		return nil
	}
	unit, heap.root = heap.root, heap.root.child
	if heap.root == nil {
		return
	}
	//一次整理最坏情况下代价为O(N)，摊还代价则为O(log N)
	//这里采用线性聚拢是不合适的，复杂之余不能持久降低宽度
	for heap.root.brother != nil {
		var list, knot = heap.root, fakeHead(&heap.root)
		for list != nil && list.brother != nil { //两两配对
			var one, another = list, list.brother
			list = another.brother
			knot.brother = merge(one, another)
			knot = knot.brother
		}
		knot.brother = list
	}
	heap.root.prev = nil
	return
}
func (heap *Heap) Pop() int {
	if heap.IsEmpty() {
		return 0
	}
	return heap.PopNode().key
}
