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

func (tree *Tree) hookSubTree(subtree *node) {
	if tree.path.isEmpty() {
		tree.root = subtree
	} else {
		if super, lf := tree.path.top(); lf {
			super.left = subtree
		} else {
			super.right = subtree
		}
	}
}
