package wavl

//成功返回true，冲突返回false。
//AVL树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Insert(key int) int {
	var root, rank = (*node)(nil), int32(1)
	if tr.root == nil {
		tr.root = newNode(key)
	} else {
		root, rank = tr.insert(key)
		if root != nil {
			tr.rebalanceAfterInsert(root, key)
		} else {
			rank = -rank
		}
	}
	return int(rank)
}

//插入节点，root != nil
func (tr *Tree) insert(key int) (*node, int32) {
	tr.path.clear()
	var root, base = tr.root, int32(0)
	for {
		root.weight++
		switch {
		case key < root.key:
			tr.path.push(root, true)
			if root.left == nil {
				root.left = newNode(key)
				return root, base + 1
			}
			root = root.left
		case key > root.key:
			base += root.subRank()
			tr.path.push(root, false)
			if root.right == nil {
				root.right = newNode(key)
				return root, base + 1
			}
			root = root.right
		default: //key == root.key
			return nil, base + root.subRank()
		}
	}
}

//回溯矫正
func (tr *Tree) rebalanceAfterInsert(root *node, key int) {
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

func newNode(key int) (unit *node) {
	unit = new(node)
	unit.key = key
	//unit.state = 0
	unit.weight = 1
	//unit.left, unit.right = nil, nil
	return unit
}
