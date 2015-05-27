package avlt

//成功返回true，没有返回false。
//AVL树删除过程包括：O(log N)的搜索，O(log N)的旋转，O(log N)的平衡因子调整。
func (tree *Tree) Remove(key int32) bool {
	tree.path.clear()
	var target = tree.root
	for target != nil && key != target.key {
		if key < target.key {
			tree.path.push(target, true)
			target = target.left
		} else {
			tree.path.push(target, false)
			target = target.right
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
	} else {
		if target.state == 1 {
			tree.path.push(target, true)
			victim = target.left
			for victim.right != nil {
				tree.path.push(victim, false)
				victim = victim.right
			}
			orphan = victim.left
		} else {
			tree.path.push(target, false)
			victim = target.right
			for victim.left != nil {
				tree.path.push(victim, true)
				victim = victim.left
			}
			orphan = victim.right
		}
	}

	if tree.path.isEmpty() { //此时victim==target
		tree.root = orphan
		return true
	}
	target.key = victim.key //李代桃僵

	var root, lf = tree.hookSubTree(orphan)
	var state, stop = root.adjust(lf), false
	for state != 0 { //如果原平衡因子为0则子树高度不变
		if root.state != 0 { //2 || -2
			root, stop = root.rotate()
			if tree.path.isEmpty() {
				tree.root = root
			} else {
				root, lf = tree.hookSubTree(root)
				if !stop {
					state = root.adjust(lf)
					continue
				}
			}
		} else if !tree.path.isEmpty() {
			root, lf = tree.path.pop()
			state = root.adjust(lf)
			continue
		}
		break
	}
	return true
}
