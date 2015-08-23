package simplebst

type node struct {
	key   int32
	left  *node
	right *node
}
type Tree struct {
	root *node
}

func (tr *Tree) IsEmpty() bool {
	return tr.root == nil
}

func (tr *Tree) Search(key int32) bool {
	var target = tr.root
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

func newNode(key int32) (unit *node) {
	unit = new(node)
	unit.key = key
	unit.left, unit.right = nil, nil
	return unit
}

//成功返回true，冲突返回false
func (tr *Tree) Insert(key int32) bool {
	if tr.root == nil {
		tr.root = newNode(key)
		return true
	}

	var parrent = tr.root
	for {
		switch {
		case key < parrent.key:
			if parrent.left == nil {
				parrent.left = newNode(key)
				return true
			}
			parrent = parrent.left
		case key > parrent.key:
			if parrent.right == nil {
				parrent.right = newNode(key)
				return true
			}
			parrent = parrent.right
		default:
			return false
		}
	}
	return true
}

//成功返回true，没有返回false
func (tr *Tree) Remove(key int32) bool {
	var target, parrent = tr.root, (*node)(nil)
	for target != nil && key != target.key {
		if key < target.key {
			target, parrent = target.left, target
		} else {
			target, parrent = target.right, target
		}
	}
	if target == nil {
		return false
	}

	var victim, orphan *node
	switch {
	case target.left == nil:
		victim, orphan = target, target.right
	case target.right == nil:
		victim, orphan = target, target.left
	default: //取中右，取中左也是可以的
		victim, parrent = target.right, target
		for victim.left != nil {
			victim, parrent = victim.left, victim
		}
		orphan = victim.right
	}

	if parrent == nil { //此时victim==target
		tr.root = orphan
	} else {
		if victim.key < parrent.key {
			parrent.left = orphan
		} else {
			parrent.right = orphan
		}
		target.key = victim.key //李代桃僵
	}
	return true
}
