package dijkstra

import (
	"unsafe"
)

type vertex struct {
	index int
	dist  uint
}
type Node struct {
	vertex
	child   *Node
	brother *Node
	prev    *Node
}

const MaxDistance = ^uint(0)

func InitHeap(size int, start int) (root *Node, list []Node) {
	list = make([]Node, size)
	for i := 0; i < size; i++ {
		list[i].index, list[i].dist, list[i].child = i, MaxDistance, nil
	}
	for i := 1; i < size; i++ {
		list[i].prev, list[i-1].brother = &list[i-1], &list[i]
	}
	list[0].prev, list[size-1].brother = nil, nil

	list[start].dist = 0
	list[start].prev, list[start].brother = nil, nil
	if start == 0 {
		list[start].child = &list[1]
	} else {
		list[start].child = &list[0]
		list[0].prev = &list[start]
		if start == size-1 {
			list[start-1].brother = nil
		} else {
			list[start+1].prev, list[start-1].brother = &list[start-1], &list[start+1]
		}
	}
	root = &list[start]
	return
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
