package avltree

type node struct {
	key   int
	state int8 //1, 0, -1
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

func (unit *node) isUnbalance() bool {
	//	return unit.state == 2 || unit.state == -2
	return (unit.state & -unit.state) == 2
}
func (unit *node) adjust(ltor bool) (oldst int8) {
	oldst = unit.state
	if ltor {
		unit.state--
	} else {
		unit.state++
	}
	return oldst
}

func (tree *Tree) hookSubTree(subroot *node) (super *node, lf bool) {
	super, lf = tree.path.pop()
	if lf {
		super.left = subroot
	} else {
		super.right = subroot
	}
	return super, lf
}
