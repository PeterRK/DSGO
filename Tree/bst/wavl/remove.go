package wavl

//成功返回序号（从1开始），没有返回-1。
//AVL树删除过程包括：O(log N)的搜索，O(log N)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Remove(key int) int {
	var target, rank = tr.findRemoveTarget(key)
	if target == nil {
		return -1
	}
	for node := target.parent; node != nil; node = node.parent {
		node.weight--
	}
	var victim, orphan = findRemoveVictim(target)

	if victim.parent == nil { //此时victim==target
		tr.root = ((*node)(nil)).tryHook(orphan)
	} else {
		tr.rebalanceAfterRemove(victim, orphan, victim.key)
		target.key = victim.key //调整好了再修正值
	}
	return int(rank)
}

func (tr *Tree) findRemoveTarget(key int) (*node, int32) {
	var target, base = tr.root, int32(0)
	for target != nil {
		if key == target.key {
			return target, base + target.subRank()
		}
		if key < target.key {
			target = target.left
		} else {
			base += target.subRank()
			target = target.right
		}
	}
	return target, -1
}

func findRemoveVictim(target *node) (victim *node, orphan *node) {
	switch {
	case target.left == nil:
		victim, orphan = target, target.right
	case target.right == nil:
		victim, orphan = target, target.left
	default:
		target.weight--
		if target.state == 1 {
			victim = target.left
			for victim.right != nil {
				victim.weight--
				victim = victim.right
			}
			orphan = victim.left
		} else {
			victim = target.right
			for victim.left != nil {
				victim.weight--
				victim = victim.left
			}
			orphan = victim.right
		}
	}
	return victim, orphan
}

func (tr *Tree) rebalanceAfterRemove(victim *node, orphan *node, key int) {
	var root = victim.parent
	var state, stop = root.state, false
	if key < root.key {
		root.left = root.tryHook(orphan)
		root.state--
	} else {
		root.right = root.tryHook(orphan)
		root.state++
	}

	for state != 0 { //如果原平衡因子为0则子树高度不变
		var super = root.parent
		if super == nil {
			if root.state != 0 { //2 || -2
				root, _ = root.rotate()
				tr.root = super.hook(root)
			}
			break
		} else {
			if root.state != 0 { //2 || -2
				root, stop = root.rotate()
				if key < super.key {
					super.left = super.hook(root)
				} else {
					super.right = super.hook(root)
				}
				if stop {
					break
				}
			}
			root, state = super, super.state
			if key < root.key {
				root.state--
			} else {
				root.state++
			}
		}
	}
}
