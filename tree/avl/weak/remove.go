package weak

//成功返回true，没有返回false。
//弱AVL树删除过程包括：O(logN)的搜索，O(1)的旋转，O(logN)的高度调整。
func (tr *Tree[T]) Remove(key T) bool {
	target := tr.findRemoveTarget(key)
	if target == nil {
		return false
	}
	victim, orphan := target.findVictim()

	if root := victim.parent; root == nil { //此时victim==target
		tr.root = root.Hook(orphan)
	} else {
		if victim.key < root.key {
			root.left = root.Hook(orphan)
		} else {
			root.right = root.Hook(orphan)
		}
		if root.left == nil && root.right == nil {
			//root.height = 0 //当节点为叶节点时，高度必须为0
			root.lDiff, root.rDiff = 1, 1
			if root.parent == nil {
				tr.root = ((*node[T])(nil)).Hook(root)
				goto Ldone
			}
			root = root.parent
		}
		tr.rebalanceAfterRemove(root, victim.key)
		target.key = victim.key //调整好了再修正值
	}
Ldone:
	tr.size--
	return true
}

func (tr *Tree[T]) findRemoveTarget(key T) *node[T] {
	target := tr.root
	for target != nil && key != target.key {
		if key < target.key {
			target = target.left
		} else {
			target = target.right
		}
	}
	return target
}

func (target *node[T]) findVictim() (victim, orphan *node[T]) {
	switch {
	case target.left == nil:
		victim, orphan = target, target.right
	case target.right == nil:
		victim, orphan = target, target.left
	default:
		//if target.left.height > target.right.height {
		if target.lDiff < target.rDiff {
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
	return victim, orphan
}

//------------直降-------------
//|    P       =       P      |
//|   / \      =      / \     |
//| S2   \     =    S1   \    |
//|       \    =   /  \   \   |
//|        X3  = Y2    Z2  X3 |

//----------------LR型----------------
//|       P        |        Z        |
//|      / \       |       / \       |
//|     S1  \      |      /   \      |
//|    / \   \     |     /     \     |
//|   /   Z1  \    |    S2      P2   |
//|  /   / \   \   |   /  \    / \   |
//| Y2  a   b   X3 |  Y1   a  b   X1 |

//----------------LL型----------------
//|       P        |       S         |
//|      / \       |      / \        |
//|    S1   \      |     /   \       |
//|   /  \   \     |    /     P1+    |
//| Y1    \   \    |  Y2     /  \    |
//|       Z?   \   |        Z?   \   |
//|             X3 |             X2- |

//P != nil
func (tr *Tree[T]) rebalanceAfterRemove(P *node[T], key T) {
	for {
		super, root := P.parent, (*node[T])(nil)
		if key > P.key {
			if P.rDiff++; P.rDiff == 2 {
				break
			}
			if P.lDiff == 2 {
				P.lDiff, P.rDiff = 1, 2
				goto Lup
			}
			if S := P.left; S.lDiff == 2 {
				if S.rDiff == 2 {
					S.lDiff, S.rDiff, P.rDiff = 1, 1, 2
					goto Lup
				}
				Z := S.right
				S.right, P.left = S.Hook(Z.left), P.Hook(Z.right)
				Z.left, Z.right = Z.hook(S), Z.hook(P)
				S.rDiff, P.lDiff = Z.lDiff, Z.rDiff
				S.lDiff, P.rDiff = 1, 1
				Z.lDiff, Z.rDiff = 2, 2
				root = Z
			} else {
				X, Z := P.right, S.right
				S.right = S.hook(P)
				P.lDiff, S.rDiff = S.rDiff, 1
				S.lDiff, P.rDiff = 2, 2
				if Z == nil {
					P.left = nil
					if X == nil {
						S.rDiff, P.lDiff, P.rDiff = 2, 1, 1
					}
				} else {
					P.left = P.hook(Z)
				}
				root = S
			}
		} else {
			if P.lDiff++; P.lDiff == 2 {
				break
			}
			if P.rDiff == 2 {
				P.rDiff, P.lDiff = 1, 2
				goto Lup
			}
			if S := P.right; S.rDiff == 2 {
				if S.lDiff == 2 {
					S.rDiff, S.lDiff, P.lDiff = 1, 1, 2
					goto Lup
				}
				Z := S.left
				S.left, P.right = S.Hook(Z.right), P.Hook(Z.left)
				Z.right, Z.left = Z.hook(S), Z.hook(P)
				S.lDiff, P.rDiff = Z.rDiff, Z.lDiff
				S.rDiff, P.lDiff = 1, 1
				Z.rDiff, Z.lDiff = 2, 2
				root = Z
			} else {
				X, Z := P.left, S.left
				S.left = S.hook(P)
				P.rDiff, S.lDiff = S.lDiff, 1
				S.rDiff, P.lDiff = 2, 2
				if Z == nil {
					P.right = nil
					if X == nil {
						S.lDiff, P.rDiff, P.lDiff = 2, 1, 1
					}
				} else {
					P.right = P.hook(Z)
				}
				root = S
			}
		}
		tr.hookSubTree(super, root)
		break
	Lup:
		if P = super; P == nil {
			break
		}
	}
}

/*
func (tr *Tree[T]) rebalanceAfterRemove(P *node[T], key T) {
	for {
		super, root := P.parent, (*node[T])(nil)
		if key > P.key {
			S, X := P.left, P.right
			if P.height-X.Height() <= 2 {
				break
			}
			if P.height-S.height > 1 {
				P.height--
				if P = super; P == nil {
					break
				}
				continue
			}
			Y, Z := S.left, S.right
			if S.height-Y.Height() > 1 {
				if S.height-Z.Height() > 1 {
					P.height--
					S.height--
					if P = super; P == nil {
						break
					}
					continue
				}
				Z.height += 2
				P.height -= 2
				S.height--
				S.right, P.left = S.Hook(Z.left), P.Hook(Z.right)
				Z.left, Z.right = Z.hook(S), Z.hook(P)
				root = Z
			} else {
				P.height--
				S.height++
				S.right = S.hook(P)
				P.left = P.Hook(Z)
				if X == nil && Z == nil {
					P.height--
				}
				root = S
			}
		} else {
			S, X := P.right, P.left
			if P.height-X.Height() <= 2 {
				break
			}
			if P.height-S.height > 1 {
				P.height--
				if P = super; P == nil {
					break
				}
				continue
			}
			Y, Z := S.right, S.left
			if S.height-Y.Height() > 1 {
				if S.height-Z.Height() > 1 {
					P.height--
					S.height--
					if P = super; P == nil {
						break
					}
					continue
				}
				Z.height += 2
				P.height -= 2
				S.height--
				S.left, P.right = S.Hook(Z.right), P.Hook(Z.left)
				Z.right, Z.left = Z.hook(S), Z.hook(P)
				root = Z
			} else {
				P.height--
				S.height++
				S.left = S.hook(P)
				P.right = P.Hook(Z)
				if X == nil && Z == nil {
					P.height--
				}
				root = S
			}
		}
		tr.hookSubTree(super, root)
		break
	}
}
*/
