package tree

import (
	"golang.org/x/exp/constraints"
)

type node[T constraints.Ordered] struct {
	key   T
	left  *node[T]
	right *node[T]
}

type NaiveBST[T constraints.Ordered] struct {
	root *node[T]
}

func (tr *NaiveBST[T]) IsEmpty() bool {
	return tr.root == nil
}

func (tr *NaiveBST[T]) Search(key T) bool {
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

func newNode[T constraints.Ordered](key T) (unit *node[T]) {
	unit = new(node[T])
	unit.key = key
	//unit.left, unit.right = nil, nil
	return unit
}

//成功返回true，冲突返回false
func (tr *NaiveBST[T]) Insert(key T) bool {
	if tr.root == nil {
		tr.root = newNode(key)
		return true
	}

	parent := tr.root
	for {
		switch {
		case key < parent.key:
			if parent.left == nil {
				parent.left = newNode(key)
				return true
			}
			parent = parent.left
		case key > parent.key:
			if parent.right == nil {
				parent.right = newNode(key)
				return true
			}
			parent = parent.right
		default:
			return false
		}
	}
}

//成功返回true，没有返回false
func (tr *NaiveBST[T]) Remove(key T) bool {
	target, parent := tr.root, (*node[T])(nil)
	for target != nil && key != target.key {
		if key < target.key {
			target, parent = target.left, target
		} else {
			target, parent = target.right, target
		}
	}
	if target == nil {
		return false
	}

	var victim, orphan *node[T]
	switch {
	case target.left == nil:
		victim, orphan = target, target.right
	case target.right == nil:
		victim, orphan = target, target.left
	default: //取中右，取中左也是可以的
		victim, parent = target.right, target
		for victim.left != nil {
			victim, parent = victim.left, victim
		}
		orphan = victim.right
	}

	if parent == nil { //此时victim==target
		tr.root = orphan
	} else {
		if victim.key < parent.key {
			parent.left = orphan
		} else {
			parent.right = orphan
		}
		target.key = victim.key //李代桃僵
	}
	return true
}
