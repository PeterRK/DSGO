package pairing

import (
	dshp "DSGO/heap"
	"unsafe"
)

type NodeG[T any] struct {
	Val   T
	child *NodeG[T]
	prev  *NodeG[T] //父兄节点
	next  *NodeG[T] //弟节点
}

func (node *NodeG[T]) hook(peer *NodeG[T]) *NodeG[T] {
	if peer != nil {
		peer.prev = node
	}
	return peer
}

type heap[T any] struct {
	less func(a, b *T) bool
	root *NodeG[T]
	size int
}

func New[T any](less func(a, b *T) bool) dshp.NodeHeap[NodeG[T]] {
	hp := new(heap[T])
	hp.less = less
	return hp
}

func (hp *heap[T]) Size() int {
	return hp.size
}

func (hp *heap[T]) IsEmpty() bool {
	return hp.root == nil
}

func (hp *heap[T]) Clear() {
	hp.root = nil
	hp.size = 0
}

func (hp *heap[T]) merge(master, slave *NodeG[T]) *NodeG[T] {
	if hp.less(&slave.Val, &master.Val) {
		master, slave = slave, master
	}
	slave.next = slave.hook(master.child)
	master.child, slave.prev = slave, master
	return master
}

func (hp *heap[T]) Push(node *NodeG[T]) {
	if node != nil {
		node.prev, node.next, node.child = nil, nil, nil
		if hp.root == nil {
			hp.root = node
		} else {
			hp.root = hp.merge(hp.root, node)
		}
		hp.size++
	}
}

func (hp *heap[T]) Top() *NodeG[T] {
	if hp.IsEmpty() {
		return nil
	}
	return hp.root
}

func (hp *heap[T]) Pop() *NodeG[T] {
	node := hp.root
	if node != nil {
		hp.root = hp.collect(node.child)
		hp.size--
		node.child = nil
	}
	return node
}

func (hp *heap[T]) FloatUp(node *NodeG[T]) {
	if node == nil {
		return
	}
	if super := node.prev; super != nil && hp.less(&node.Val, &super.Val) {
		node.prev = nil
		if super.next == node { //super为兄
			super.next, node.next = super.hook(node.next), nil
		} else { //super为父
			super.child, node.next = super.hook(node.next), nil
		}
		hp.root = hp.merge(hp.root, node)
	}
}

func fakeHeadG[T any](spt **NodeG[T]) *NodeG[T] {
	base := uintptr(unsafe.Pointer(spt))
	off := unsafe.Offsetof((*spt).next)
	return (*NodeG[T])(unsafe.Pointer(base - off))
}

func (hp *heap[T]) collect(head *NodeG[T]) *NodeG[T] {
	if head != nil && head.next != nil {
		list, last := head, fakeHeadG(&head)
		for list != nil && list.next != nil {
			master, slave := list, list.next
			list = slave.next
			last.next = last.hook(hp.merge(master, slave))
			last = last.next
		}
		head.prev = nil
		if list == nil {
			head, list = last, last.prev
		} else {
			head, list = list, last
		}
		for list != nil {
			last, list = list, list.prev
			head = hp.merge(head, last)
		}
		head.prev, head.next = nil, nil
	}
	return head
}
