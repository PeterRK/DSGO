package wavl

//成功返回序号（从1开始），冲突返回序号的负值。
//AVL树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Insert(key int) int {
	var root, rank = (*node)(nil), int32(1)
	if tr.root == nil {
		tr.root = newNode(nil, key)
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
	var root, base = tr.root, int32(0)
	for {
		root.weight++
		switch {
		case key < root.key:
			if root.left == nil {
				root.left = newNode(root, key)
				return root, base + 1
			}
			root = root.left
		case key > root.key:
			base += root.subRank()
			if root.right == nil {
				root.right = newNode(root, key)
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
	for {
		var state = root.state
		if key < root.key {
			root.state++
		} else {
			root.state--
		}
		if state == 0 && root.parent != nil {
			root = root.parent
			continue
		}
		if state != 0 && root.state != 0 { //2 || -2
			var super = root.parent
			root, _ = root.rotate()
			if super == nil {
				tr.root = super.hook(root)
			} else {
				if key < super.key {
					super.left = super.hook(root)
				} else {
					super.right = super.hook(root)
				}
			}
		}
		break
	}
}

func newNode(parent *node, key int) (unit *node) {
	unit = new(node)
	//unit.state = 0
	unit.weight = 1
	//unit.left, unit.right = nil, nil
	unit.parent, unit.key = parent, key
	return unit
}
