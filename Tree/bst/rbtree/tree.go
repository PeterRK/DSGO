package rbtree

type node struct {
	key   int
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

func (tree *Tree) Search(key int) bool {
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

func (tree *Tree) hookSubTree(subroot *node) {
	if tree.path.isEmpty() {
		tree.root = subroot
	} else {
		if super, lf := tree.path.top(); lf {
			super.left = subroot
		} else {
			super.right = subroot
		}
	}
}
