package avl

import (
	"fmt"
	"constraints"
)

type node[T constraints.Ordered] struct {
	key    T
	state  int8 //(2), 1, 0, -1, (-2)
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

func newNode[T constraints.Ordered](parent *node[T], key T) *node[T] {
	unit := new(node[T])
	//unit.state = 0
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
	fmt.Println(root.state, root.key)
	root.right.debug(indent+1)
}
