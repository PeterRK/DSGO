package avl

//成功返回true，冲突返回false。
//AVL树插入过程包括：O(logN)的搜索，O(1)的旋转，O(logN)的平衡因子调整。
func (tr *Tree[T]) Insert(key T) bool {
	if tr.root == nil {
		tr.root = newNode[T](nil, key)
	} else {
		root := tr.insert(key)
		if root == nil {
			return false
		}
		tr.rebalanceAfterInsert(root, key)
	}
	tr.size++
	return true
}

//插入节点，root != nil
func (tr *Tree[T]) insert(key T) *node[T] {
	root := tr.root
	for {
		switch {
		case key < root.key:
			if root.left == nil {
				root.left = newNode[T](root, key)
				return root
			}
			root = root.left
		case key > root.key:
			if root.right == nil {
				root.right = newNode[T](root, key)
				return root
			}
			root = root.right
		default: //key == root.key
			return nil
		}
	}
}

//回溯矫正
func (tr *Tree[T]) rebalanceAfterInsert(root *node[T], key T) {
	for {
		state := root.state
		if key < root.key {
			root.state++
		} else {
			root.state--
		}
		if state == 0 {
			if root.parent != nil {
				root = root.parent
				continue
			}
		} else if root.state != 0 {
			super := root.parent
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
