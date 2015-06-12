package graph

import (
	"unsafe"
)

type node struct {
	Vertex
	child *node
	prev  *node
	next  *node
}

func NewVector(size int) []node {
	return make([]node, size)
}
func NewHeap(size int, start int) (root *node, list []node) {
	list = make([]node, size)
	for i := 0; i < size; i++ {
		list[i].Index, list[i].Dist, list[i].child = i, MaxDistance, nil
	}
	for i := 1; i < size; i++ {
		list[i].prev, list[i-1].next = &list[i-1], &list[i]
	}
	list[0].prev, list[size-1].next = nil, nil

	list[start].Dist = 0
	list[start].prev, list[start].next = nil, nil
	if start == 0 {
		list[start].child = &list[1]
	} else {
		list[start].child = &list[0]
		list[0].prev = &list[start]
		if start == size-1 {
			list[start-1].next = nil
		} else {
			list[start+1].prev, list[start-1].next = &list[start-1], &list[start+1]
		}
	}
	root = &list[start]
	return
}

func (unit *node) hook(peer *node) *node {
	if peer != nil {
		peer.prev = unit
	}
	return peer
}
func fakeHead(spt **node) *node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).next)
	return (*node)(unsafe.Pointer(base - off))
}
func merge(one *node, another *node) *node {
	if one.Dist > another.Dist {
		one, another = another, one
	}
	another.next = another.hook(one.child)
	one.child, another.prev = another, one
	return one
}

func Insert(root *node, unit *node) *node {
	if unit != nil {
		unit.child, unit.next, unit.prev = nil, nil, nil
		if root == nil {
			root = unit
		} else {
			root = merge(root, unit)
		}
	}
	return root
}

func collect(head *node) *node {
	if head != nil {
		for head.next != nil {
			var list, knot = head, fakeHead(&head)
			for list != nil && list.next != nil {
				var one, another = list, list.next
				list = another.next
				knot.next = merge(one, another)
				knot = knot.next
			}
			knot.next = list
		}
		head.prev = nil
	}
	return head
}

func Extract(root *node) *node {
	if root != nil {
		return collect(root.child)
	}
	return nil
}

func FloatUp(root *node, target *node, distance uint) *node {
	if target != nil && distance < target.Dist {
		target.Dist = distance
		var super = target.prev
		if super != nil {
			if super.child == target {
				if super.Dist > distance {
					super.child, target.next = super.hook(target.next), nil
					root = merge(root, target)
				}
			} else {
				super.next, target.next = super.hook(target.next), nil
				root = merge(root, target)
			}
		}
	}
	return root
}
