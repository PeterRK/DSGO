package avltree

//成功返回true，冲突返回false。
//AVL树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tree *Tree) Insert(key int) bool {
	if tree.root == nil {
		tree.root = newNode(key)
		return true
	}
	tree.path.clear()
	var parent = tree.root
	for {
		if key < parent.key {
			tree.path.push(parent, true)
			if parent.left == nil {
				parent.left = newNode(key)
				break
			}
			parent = parent.left
		} else if key > parent.key {
			tree.path.push(parent, false)
			if parent.right == nil {
				parent.right = newNode(key)
				break
			}
			parent = parent.right
		} else { //key == parent.key
			return false
		}
	}

	var state, subroot, lf = int8(0), parent, false
	for !tree.path.isEmpty() && state == 0 {
		subroot, lf = tree.path.pop()
		state = subroot.adjust(!lf)
	}
	if subroot.isUnbalance() {
		newsub, _ := subroot.rotate()
		if tree.path.isEmpty() {
			tree.root = newsub
		} else {
			tree.hookSubTree(newsub)
		}
	}
	return true
}

func newNode(key int) (unit *node) {
	unit = new(node)
	unit.key, unit.state = key, 0
	unit.left, unit.right = nil, nil
	return unit
}
