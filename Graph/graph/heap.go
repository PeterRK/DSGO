package graph

import (
	"unsafe"
)

type node struct {
	Vertex
	child *node
	prev  *node //父兄节点
	next  *node //弟节点
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
	return &list[start], list
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
	another.next = one.child
	if one.child != nil {
		one.child.prev = another
	}
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
func Extract(root *node) *node {
	if root == nil {
		return nil
	}
	root = root.child
	if root == nil {
		return nil
	}
	for root.next != nil {
		var list, knot = root, fakeHead(&root)
		for list != nil && list.next != nil {
			var one, another = list, list.next
			list = another.next
			knot.next = merge(one, another)
			knot = knot.next
		}
		knot.next = list
	}
	root.prev = nil
	return root
}

func FloatUp(root *node, target *node, distance uint) *node {
	if target == nil || distance >= target.Dist {
		return root
	}
	target.Dist = distance
	if target == root {
		return root
	}

	for {
		var brother = target
		for brother.prev.child != brother {
			brother = brother.prev
		}
		var parent = brother.prev
		if parent.Dist <= target.Dist {
			return root
		}

		target.next, parent.next = parent.next, target.next
		if parent.next != nil {
			parent.next.prev = parent
		}
		if target.next != nil {
			target.next.prev = target
		}

		parent.child = target.child
		if parent.child != nil {
			parent.child.prev = parent
		}

		if brother != target {
			parent.prev, target.prev = target.prev, parent.prev
			parent.prev.next = parent
			target.child, brother.prev = brother, target
		} else {
			target.prev = parent.prev
			target.child, parent.prev = parent, target
		}

		if target.prev == nil {
			root = target
			break
		} else {
			var super = target.prev
			if super.next == parent {
				super.next = target
			} else {
				super.child = target
			}
		}
	}
	return root
}
