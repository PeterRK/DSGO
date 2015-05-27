package avlt

type node struct {
	key   int32
	state int8 //(2), 1, 0, -1, (-2)
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

func (unit *node) adjust(ltor bool) (oldst int8) {
	oldst = unit.state
	if ltor {
		unit.state--
	} else {
		unit.state++
	}
	return oldst
}

func (tree *Tree) hookSubTree(root *node) (super *node, lf bool) {
	super, lf = tree.path.pop()
	if lf {
		super.left = root
	} else {
		super.right = root
	}
	return super, lf
}
