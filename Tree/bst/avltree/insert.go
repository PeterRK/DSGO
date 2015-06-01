package avltree

//成功返回true，冲突返回false。
//AVL树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Insert(key int32) bool {
	if tr.root == nil {
		tr.root = newNode(nil, key)
		return true
	}

	var root = tr.root
	for {
		if key < root.key {
			if root.left == nil {
				root.left = newNode(root, key)
				break
			}
			root = root.left
		} else if key > root.key {
			if root.right == nil {
				root.right = newNode(root, key)
				break
			}
			root = root.right
		} else { //key == root.key
			return false
		}
	}

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
	return true
}

func newNode(parent *node, key int32) (unit *node) {
	unit = new(node)
	unit.key, unit.state = key, 0
	unit.parent = parent
	unit.left, unit.right = nil, nil
	return unit
}
