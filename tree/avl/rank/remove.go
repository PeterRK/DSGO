package rank

//成功返回序号（从1开始），没有返回0。
//AVL树删除过程包括：O(logN)的搜索，O(logN)的旋转，O(logN)的平衡因子调整。
func (tr *Tree[T]) Remove(key T) int {
	target, rank := tr.findRemoveTarget(key)
	if target == nil {
		return 0
	}
	for node := target.parent; node != nil; node = node.parent {
		node.weight--
	}
	victim, orphan := target.findVictim()

	if victim.parent == nil { //此时victim==target
		tr.root = ((*node[T])(nil)).Hook(orphan)
	} else {
		tr.removeAndRebalance(victim, orphan)
		target.key = victim.key //调整好了再修正值
	}
	tr.size--
	return int(rank)
}

func (tr *Tree[T]) findRemoveTarget(key T) (*node[T], int32) {
	target, base := tr.root, int32(0)
	for target != nil {
		if key == target.key {
			return target, base + target.rank()
		}
		if key < target.key {
			target = target.left
		} else {
			base += target.rank()
			target = target.right
		}
	}
	return target, -1
}

func (target *node[T]) findVictim() (victim, orphan *node[T]) {
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

//victim.parent != nil
func (tr *Tree[T]) removeAndRebalance(victim, orphan *node[T]) {
	root := victim.parent
	state, stop := root.state, false
	if victim.key < root.key {
		root.left = root.Hook(orphan)
		root.state--
	} else {
		root.right = root.Hook(orphan)
		root.state++
	}

	for state != 0 { //如果原平衡因子为0则子树高度不变
		super := root.parent
		if super == nil {
			if root.state != 0 { //2 || -2
				root, _ = root.rotate()
				tr.root = super.hook(root)
			}
			break
		} else {
			if root.state != 0 { //2 || -2
				root, stop = root.rotate()
				if victim.key < super.key {
					super.left = super.hook(root)
				} else {
					super.right = super.hook(root)
				}
				if stop {
					break
				}
			}
			root, state = super, super.state
			if victim.key < root.key {
				root.state--
			} else {
				root.state++
			}
		}
	}
}
