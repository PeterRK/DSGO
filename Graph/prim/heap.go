package prim

import (
	"unsafe"
)

type vertex struct {
	id   int  //本顶点编号
	lnk  int  //关联顶点编号
	dist uint //与关联顶点间的距离
}
type Node struct {
	vertex
	child *Node
	prev  *Node //父兄节点
	next  *Node //弟节点
}

func fakeHead(spt **Node) *Node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).next)
	return (*Node)(unsafe.Pointer(base - off))
}
func merge(one *Node, another *Node) *Node {
	if one.dist > another.dist {
		one, another = another, one
	}
	another.next = one.child
	if one.child != nil {
		one.child.prev = another
	}
	one.child, another.prev = another, one
	return one
}
func Insert(root *Node, unit *Node) *Node {
	if unit == nil {
		return root
	}
	unit.child, unit.next, unit.prev = nil, nil, nil
	if root == nil {
		root = unit
	} else {
		root = merge(root, unit)
	}
	return root
}
func Extract(root *Node) *Node {
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

func FloatUp(root *Node, target *Node, distance uint) *Node {
	if target == nil || distance >= target.dist {
		return root
	}
	target.dist = distance
	if target == root {
		return root
	}

	for {
		var brother = target
		for brother.prev.child != brother {
			brother = brother.prev
		}
		var parent = brother.prev
		if parent.dist <= target.dist {
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
