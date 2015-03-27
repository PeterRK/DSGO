package avltree

//成功返回true，没有返回false。
//AVL树删除过程包括：O(log N)的搜索，O(log N)的旋转，O(log N)的平衡因子调整。
func (tree *Tree) Remove(key int) bool {
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

	var subroot, lf = tree.hookSubTree(orphan)
	var state = subroot.adjust(lf)
	for state != 0 { //如果原平衡因子为0则子树高度不变
		if subroot.isUnbalance() {
			var newsub, balanced = subroot.rotate()
			if tree.path.isEmpty() {
				tree.root = newsub
			} else {
				subroot, lf = tree.hookSubTree(newsub)
				if !balanced {
					state = subroot.adjust(lf)
					continue
				}
			}
		} else if !tree.path.isEmpty() {
			subroot, lf = tree.path.pop()
			state = subroot.adjust(lf)
			continue
		}
		break
	}
	return true
}
