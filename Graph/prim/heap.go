package prim

import (
	"unsafe"
)

type vertex struct {
	id   int
	next int
	dist uint
}
type Node struct {
	vertex
	child   *Node
	brother *Node
	prev    *Node
}

func fakeHead(this **Node) *Node {
	var base = uintptr(unsafe.Pointer(this))
	var off = unsafe.Offsetof((*this).brother)
	return (*Node)(unsafe.Pointer(base - off))
}
func merge(one *Node, another *Node) *Node {
	if one.dist > another.dist {
		one, another = another, one
	}
	another.brother = one.child
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
	unit.child, unit.brother, unit.prev = nil, nil, nil
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
	for root.brother != nil {
		var list, knot = root, fakeHead(&root)
		for list != nil && list.brother != nil {
			var one, another = list, list.brother
			list = another.brother
			knot.brother = merge(one, another)
			knot = knot.brother
		}
		knot.brother = list
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
		var big_bro = target
		for big_bro.prev.child != big_bro {
			big_bro = big_bro.prev
		}
		var parent = big_bro.prev
		if parent.dist <= target.dist {
			return root
		}

		parent.brother, target.brother = target.brother, parent.brother
		if parent.brother != nil {
			parent.brother.prev = parent
		}
		if target.brother != nil {
			target.brother.prev = target
		}

		parent.child = target.child
		if parent.child != nil {
			parent.child.prev = parent
		}

		if big_bro != target {
			parent.prev, target.prev = target.prev, parent.prev
			parent.prev.brother = parent
			target.child, big_bro.prev = big_bro, target
		} else {
			target.prev = parent.prev
			target.child, parent.prev = parent, target
		}

		if target.prev == nil {
			root = target
			break
		} else {
			var super = target.prev
			if super.brother == parent {
				super.brother = target
			} else {
				super.child = target
			}
		}
	}
	return root
}
