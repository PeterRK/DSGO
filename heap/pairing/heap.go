package pairing

import (
	"constraints"
	"unsafe"
)

//虽然Fibonacci理论复杂度更好，但配对堆实际上更为实用。
type Heap[T constraints.Ordered] struct {
	root *Node[T]
	size int
}

type Node[T constraints.Ordered] struct {
	key   T
	child *Node[T]
	prev  *Node[T] //父兄节点
	next  *Node[T] //弟节点
}

func (node *Node[T]) hook(peer *Node[T]) *Node[T] {
	if peer != nil {
		peer.prev = node
	}
	return peer
}

func (hp *Heap[T]) Size() int {
	return hp.size
}

func (hp *Heap[T]) IsEmpty() bool {
	return hp.root == nil
}

func (hp *Heap[T]) Clear() {
	hp.root = nil
	hp.size = 0
}

func (hp *Heap[T]) Top() T {
	if hp.IsEmpty() {
		panic("empty heap")
	}
	return hp.root.key
}

//master != nil && slave != nil
func merge[T constraints.Ordered](master, slave *Node[T]) *Node[T] {
	if master.key > slave.key {
		master, slave = slave, master
	}
	slave.next = slave.hook(master.child)
	master.child, slave.prev = slave, master
	return master
}

func (hp *Heap[T]) Merge(other *Heap[T]) {
	if hp != other && !other.IsEmpty() {
		if hp.IsEmpty() {
			*hp = *other
		} else {
			hp.root = merge(hp.root, other.root)
			hp.size += other.size
		}
		other.Clear()
	}
}

//这货Push时不怎么管整理，到Pop时再做
func (hp *Heap[T]) PushNode(node *Node[T]) {
	if node != nil {
		node.prev, node.next, node.child = nil, nil, nil
		if hp.root == nil {
			hp.root = node
		} else {
			hp.root = merge(hp.root, node)
		}
		hp.size++
	}
}

func (hp *Heap[T]) Push(key T) *Node[T] {
	node := new(Node[T])
	node.key = key
	hp.PushNode(node)
	return node
}

func fakeHead[T constraints.Ordered](spt **Node[T]) *Node[T] {
	base := uintptr(unsafe.Pointer(spt))
	off := unsafe.Offsetof((*spt).next)
	return (*Node[T])(unsafe.Pointer(base - off))
}

func collect[T constraints.Ordered](head *Node[T]) *Node[T] {
	if head != nil && head.next != nil {
		list, last := head, fakeHead(&head)
		for list != nil && list.next != nil { //两两配对
			master, slave := list, list.next
			list = slave.next
			last.next = last.hook(merge(master, slave))
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
			head = merge(head, last)
		}
		head.prev, head.next = nil, nil
	}
	return head
}

func collect_v2[T constraints.Ordered](head *Node[T]) *Node[T] {
	if head != nil {
		for head.next != nil {
			list, knot := head, fakeHead(&head)
			for list != nil && list.next != nil { //两两配对
				master, slave := list, list.next
				list = slave.next
				knot.next = merge(master, slave)
				knot = knot.next
			}
			knot.next = list
		}
		head.prev = nil
	}
	return head
}

func (hp *Heap[T]) PopNode() *Node[T] {
	node := hp.root
	if node != nil {
		hp.root = collect(node.child)
		hp.size--
		node.child = nil
	}
	return node
}

func (hp *Heap[T]) Pop() T {
	node := hp.PopNode()
	if node == nil {
		panic("empty heap")
	}
	return node.key
}

func (hp *Heap[T]) Remove(node *Node[T]) {
	if node != nil {
		hp.size--
		if super := node.prev; super == nil { //根
			hp.root = collect(node.child)
		} else {
			if super.child == node { //super为父
				super.child = super.hook(node.next)
			} else { //super为兄
				super.next = super.hook(node.next)
			}
			if child := collect(node.child); child != nil {
				hp.root = merge(hp.root, child)
			}
		}
		node.child, node.prev, node.next = nil, nil, nil
	}
}

func (hp *Heap[T]) FloatUp(node *Node[T], value T) {
	if node != nil && value < node.key {
		node.key = value
		if super := node.prev; super != nil && super.key > value {
			node.prev = nil
			if super.next == node { //super为兄
				super.next, node.next = super.hook(node.next), nil
			} else { //super为父
				super.child, node.next = super.hook(node.next), nil
			}
			hp.root = merge(hp.root, node)
		}
	}
}
