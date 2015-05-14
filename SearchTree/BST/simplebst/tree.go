package simplebst

type node struct {
	key   int
	left  *node
	right *node
}
type Tree struct {
	root *node
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

func newNode(key int) (unit *node) {
	unit = new(node)
	unit.key = key
	unit.left, unit.right = nil, nil
	return
}

//成功返回true，冲突返回false
func (tree *Tree) Insert(key int) bool {
	if tree.root == nil {
		tree.root = newNode(key)
		return true
	}

	var parrent = tree.root
	for {
		if key == parrent.key {
			return false
		}
		if key < parrent.key {
			if parrent.left == nil {
				parrent.left = newNode(key)
				return true
			}
			parrent = parrent.left
		} else {
			if parrent.right == nil {
				parrent.right = newNode(key)
				return true
			}
			parrent = parrent.right
		}
	}
	return true
}

//成功返回true，没有返回false
func (tree *Tree) Remove(key int) bool {
	var target, parrent, lf = tree.root, (*node)(nil), false
	for target != nil && key != target.key {
		if key < target.key {
			target, parrent, lf = target.left, target, true
		} else {
			target, parrent, lf = target.right, target, false
		}
	}
	if target == nil {
		return false
	}

	var victim, orphan *node = nil, nil
	if target.left == nil {
		victim, orphan = target, target.right
	} else if target.right == nil {
		victim, orphan = target, target.left
	} else { //取中右，取中左也是可以的
		victim, parrent, lf = target.right, target, false
		for victim.left != nil {
			victim, parrent, lf = victim.left, victim, true
		}
		orphan = victim.right
	}

	if parrent == nil { //此时victim==target
		tree.root = orphan
	} else {
		if lf {
			parrent.left = orphan
		} else {
			parrent.right = orphan
		}
		target.key = victim.key //李代桃僵
	}

	return true
}
