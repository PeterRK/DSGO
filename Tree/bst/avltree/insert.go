package avltree

//成功返回true，冲突返回false。
//AVL树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Insert(key int32) bool {
	if tr.root == nil {
		tr.root = newNode(nil, key)
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
	var root = tr.root
	for {
		switch {
		case key < root.key:
			if root.left == nil {
				root.left = newNode(root, key)
				return root
			}
			root = root.left
		case key > root.key:
			if root.right == nil {
				root.right = newNode(root, key)
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

func newNode(parent *node, key int32) (unit *node) {
	unit = new(node)
	//unit.state = 0
	//unit.left, unit.right = nil, nil
	unit.parent, unit.key = parent, key
	return unit
}
