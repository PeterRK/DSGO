package rank

//成功返回序号（从1开始），冲突返回序号的负值。
//AVL树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tr *Tree[T]) Insert(key T) int {
	root, rank := (*node[T])(nil), int32(1)
	if tr.root == nil {
		tr.root = newNode[T](nil, key)
	} else {
		root, rank = tr.insert(key)
		if root == nil {
			return int(-rank)
		}
		tr.rebalanceAfterInsert(root, key)
	}
	tr.size++
	return int(rank)
}

//插入节点，root != nil
func (tr *Tree[T]) insert(key T) (*node[T], int32) {
	root, base := tr.root, int32(0)
	for {
		root.weight++
		switch {
		case key < root.key:
			if root.left == nil {
				root.left = newNode[T](root, key)
				return root, base + 1
			}
			root = root.left
		case key > root.key:
			base += root.subRank()
			if root.right == nil {
				root.right = newNode[T](root, key)
				return root, base + 1
			}
			root = root.right
		default: //key == root.key
			return nil, base + root.subRank()
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
