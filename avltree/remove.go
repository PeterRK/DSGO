package avltree

//AVL树删除过程包括：O(log N)的搜索，O(log N)的旋转，O(log N)的平衡因子调整
func (tree *Tree) Remove(key int) {
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
		return
	}

	var victim, orphan *node = nil, nil
	if target.left == nil {
		victim, orphan = target, target.right
	} else if target.right == nil {
		victim, orphan = target, target.left
	} else {
		if target.balance == 1 {
			tree.path.push(target, true)
			victim = target.left //取中左
			for victim.right != nil {
				tree.path.push(victim, false)
				victim = victim.right
			}
			orphan = victim.left
		} else {
			tree.path.push(target, false)
			victim = target.right //取中右
			for victim.left != nil {
				tree.path.push(victim, true)
				victim = victim.left
			}
			orphan = victim.right
		}
	} //如果需要删除的节点有两个儿子，那么问题可以被转化成删除另一个只有一个儿子的节点的问题

	if tree.path.isEmpty() { //此时victim==target
		tree.root = orphan
	} else {
		var parent, victim_is_left = tree.path.pop() //非根，parent!=nil
		var tips = parent.balance
		if victim_is_left {
			parent.left = orphan
			parent.balance-- //左支减短
		} else {
			parent.right = orphan
			parent.balance++ //右支减短
		}

		for tips != 0 { //如果原balance为0则子树高度不变
			if parent.balance == 0 { //无需旋转，super无需重接子树
				if tree.path.isEmpty() {
					break
				}
				parent, victim_is_left = tree.path.pop()
				tips = parent.balance
				if victim_is_left {
					parent.balance-- //左支减短
				} else {
					parent.balance++ //右支减短
				}
				continue
			}
			var subtree, keep_height = parent.rotate()
			if tree.path.isEmpty() { //到根
				tree.root = subtree
				break
			}
			var super, sub_is_left = tree.path.pop()
			if keep_height {
				if sub_is_left {
					super.left = subtree
				} else {
					super.right = subtree
				}
				break
			}
			//降高旋转
			tips = super.balance
			if sub_is_left {
				super.left = subtree
				super.balance-- //左支减短
			} else {
				super.right = subtree
				super.balance++ //右支减短
			}
			parent = super
		}
		target.key = victim.key //李代桃僵
	}
	return
}
