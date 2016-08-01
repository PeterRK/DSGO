package heap

import (
	"unsafe"
)

type node struct {
	Vertex
	child *node
	prev  *node
	next  *node
}

func NewVectorPH(size int) []node {
	return make([]node, size)
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
func FloatUp(root *node, target *node, dist uint) *node {
	if target != nil {
		target.Dist = dist
		if super := target.prev; super != nil && super.Dist > target.Dist {
			target.prev = nil
			if super.next == target { //super为兄
				super.next, target.next = super.hook(target.next), nil
			} else { //super为父
				super.child, target.next = super.hook(target.next), nil
			}
			root = merge(root, target)
		}
	}
	return root
}
