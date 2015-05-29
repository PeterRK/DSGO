package rbt

//AVL树的平衡因子有5态，需要3bit存储空间。
//而红黑树的平衡因子只需1bit，有时候可以巧妙地隐藏掉。
type node struct {
	key   int32
	black bool
	left  *node
	right *node
}
type Tree struct {
	root *node
	path stack
}

func (tree *Tree) IsEmpty() bool {
	return tree.root == nil
}

func (tree *Tree) Search(key int32) bool {
	var target = tree.root
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

func (tree *Tree) hookSubTree(root *node) {
	if tree.path.isEmpty() {
		tree.root = root
	} else {
		if super, lf := tree.path.top(); lf {
			super.left = root
		} else {
			super.right = root
		}
	}
}
