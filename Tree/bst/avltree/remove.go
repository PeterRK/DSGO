package avltree

//成功返回true，没有返回false。
//AVL树删除过程包括：O(log N)的搜索，O(log N)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Remove(key int32) bool {
	var target = tr.root
	for target != nil && key != target.key {
		if key < target.key {
			target = target.left
		} else {
			target = target.right
		}
	}
	if target == nil {
		return false
	}

	var victim, orphan *node
	switch {
	case target.left == nil:
		victim, orphan = target, target.right
	case target.right == nil:
		victim, orphan = target, target.left
	default:
		if target.state == 1 {
			victim = target.left
			for victim.right != nil {
				victim = victim.right
			}
			orphan = victim.left
		} else {
			victim = target.right
			for victim.left != nil {
				victim = victim.left
			}
			orphan = victim.right
		}
	}

	var root = victim.parent
	if root == nil { //此时victim==target
		tr.root = root.tryHook(orphan)
		return true
	}
	key = victim.key

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
	target.key = key
	return true
}
