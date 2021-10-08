package redblack

import (
	"constraints"
	"fmt"
)

//AVL树的平衡因子有5态，需要3bit存储空间。
//而红黑树的平衡因子只需1bit，有时候可以巧妙地隐藏掉。
type node[T constraints.Ordered] struct {
	key    T
	black  bool
	parent *node[T]
	left   *node[T]
	right  *node[T]
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

func newNode[T constraints.Ordered](parent *node[T], key T) (unit *node[T]) {
	unit = new(node[T])
	//unit.black = false
	//unit.left, unit.right = nil, nil
	unit.parent, unit.key = parent, key
	return unit
}

func (root *node[T]) debug(indent int) {
	if root == nil {
		return
	}
	root.left.debug(indent+1)
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
	fmt.Println(root.black, root.key)
	root.right.debug(indent+1)
}