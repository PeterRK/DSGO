package avlt

//成功返回true，冲突返回false。
//AVL树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Insert(key int32) bool {
	if tr.root == nil {
		tr.root = newNode(key)
	} else {
		var root = tr.insert(key)
		if root == nil {
			return false
		}
		tr.rebalanceAfterInsert(root, key)
	}
	return true
}

//插入节点，root != nil
func (tr *Tree) insert(key int32) *node {
	tr.path.clear()
	var root = tr.root
	for {
		switch {
		case key < root.key:
			tr.path.push(root, true)
			if root.left == nil {
				root.left = newNode(key)
				return root
			}
			root = root.left
		case key > root.key:
			tr.path.push(root, false)
			if root.right == nil {
				root.right = newNode(key)
				return root
			}
			root = root.right
		default: //key == root.key
			return nil
		}
	}
}

//回溯矫正
func (tr *Tree) rebalanceAfterInsert(root *node, key int32) {
	var state, lf = int8(0), false
	for !tr.path.isEmpty() && state == 0 {
		root, lf = tr.path.pop()
		state = root.adjust(!lf)
	}
	if state != 0 && root.state != 0 { //2 || -2
		root, _ = root.rotate()
		if tr.path.isEmpty() {
			tr.root = root
		} else {
			tr.hookSubTree(root)
		}
	}
}

func newNode(key int32) (unit *node) {
	unit = new(node)
	unit.key = key
	//unit.state = 0
	//unit.left, unit.right = nil, nil
	return unit
}
