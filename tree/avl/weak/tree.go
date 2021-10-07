package weak

import (
	"constraints"
)

//弱AVL树满足如下三条约束：
//nil的高度为-1
//当节点为叶节点（子节点皆为nil）时，高度必须为0
//节点和子节点的高度差为1或2
type node[T constraints.Ordered] struct {
	height int8 //足以支持2^64的节点数
	key    T
	parent *node[T]
	left   *node[T]
	right  *node[T]
}

func (unit *node[T]) Height() int8 {
	if unit == nil {
		return -1
	}
	return unit.height
}

type Tree[T constraints.Ordered] struct {
	root *node[T]
	size int
}

func (tr *Tree[T]) Size() int {
	return tr.size
}

func (tr *Tree[T]) IsEmpty() bool {
	return tr.root == nil
}

func (tr *Tree[T]) Clear() {
	tr.root = nil
	tr.size = 0
}

func (tr *Tree[T]) Search(key T) bool {
	target := tr.root
	for target != nil {
		if key == target.key {
			return true
		}
		if key < target.key {
			target = target.left
		} else {
			target = target.right
		}
	}
	return false
}

func (parent *node[T]) Hook(child *node[T]) *node[T] {
	if child != nil {
		child.parent = parent
	}
	return child
}
func (parent *node[T]) hook(child *node[T]) *node[T] {
	child.parent = parent
	return child
}

func (tr *Tree[T]) hookSubTree(super, root *node[T]) {
	if super == nil {
		tr.root = super.hook(root)
	} else {
		if root.key < super.key {
			super.left = super.hook(root)
		} else {
			super.right = super.hook(root)
		}
	}
}

func newNode[T constraints.Ordered](parent *node[T], key T) *node[T] {
	unit := new(node[T])
	//unit.height = 0
	//unit.left, unit.right = nil, nil
	unit.parent, unit.key = parent, key
	return unit
}
