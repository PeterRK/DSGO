package redblack

//成功返回true，没有返回false。
//红黑树删除过程包括：O(logN)的搜索，O(1)的旋转，O(logN)的平衡因子调整。
func (tr *Tree[T]) Remove(key T) bool {
	target := tr.findRemoveTarget(key)
	if target == nil {
		return false
	}
	victim, orphan := target.findVictim()

	if victim.parent == nil { //此时victim==target
		tr.root = orphan
		if tr.root != nil {
			tr.root.black = true
			tr.root.parent = nil
		}
	} else {
		root := victim.parent
		if victim.key < root.key {
			root.left = root.Hook(orphan)
		} else {
			root.right = root.Hook(orphan)
		}
		if victim.black { //红victim随便删，黑的要考虑
			if orphan != nil && !orphan.black {
				orphan.black = true //红子变黑顶上
			} else {
				tr.rebalanceAfterRemove(root, victim.key)
			}
		}
		target.key = victim.key
	}
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

func (target *node[T])findVictim() (victim, orphan *node[T]) {
	switch {
	case target.left == nil:
		victim, orphan = target, target.right
	case target.right == nil:
		victim, orphan = target, target.left
	default:
		victim = target.right
		for victim.left != nil {
			victim = victim.left
		}
		orphan = victim.right
	}
	return victim, orphan
}

//----------------红叔模式----------------
//|        bG        |        bU        |
//|       /  \       |       /  \       |
//|     bO    rU     |     rG    bR     |
//|          /  \    |    /  \          |
//|        bL    bR  |  bO    bL        |

//------------------双黑------------------
//|        xG        |        bG        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bO    rU     |
//|          /  \    |          /  \    |
//|        bL    bR  |        bL    bR  |

//----------------中黑外红----------------
//|        xG        |        xU        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bG    bR     |
//|          /  \    |    /  \          |
//|        bL    rR  |  bO    bL        |

//------------------中红------------------
//|        xG        |        xL        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bG    bU     |
//|          /  \    |    /  \  /  \    |
//|        rL    xR  |  bO   u  v   xR  |
//|       /  \       |                  |
//|      u    v      |                  |

func (tr *Tree[T]) rebalanceAfterRemove(G *node[T], key T) {
	for { //剩下情况：victim黑，orphan也黑，此时victim(orphan顶替)的兄弟必然存在
		super := G.parent
		if key < G.key {
			U := G.right //U != nil
			L, R := U.left, U.right
			if !U.black { //红U下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
				G.right, U.left = G.hook(L), U.hook(G)
				U.black, G.black = true, false
				tr.hookSubTree(super, U)
				continue //变出黑U后再行解决
			} else {
				if L == nil || L.black {
					if R == nil || R.black { //双黑，变色解决
						U.black = false
						if G.black && super != nil {
							G = super
							continue //上溯
						}
						G.black = true
					} else { //中黑外红
						G.right, U.left = G.Hook(L), U.hook(G)
						U.black, G.black, R.black = G.black, true, true
						tr.hookSubTree(super, U)
					}
				} else { //中红
					U.left, G.right = U.Hook(L.right), G.Hook(L.left)
					L.right, L.left = L.hook(U), L.hook(G)
					L.black, G.black = G.black, true
					tr.hookSubTree(super, L)
				}
			}
		} else {
			U := G.left //U != nil
			R, L := U.right, U.left
			if !U.black { //红U下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
				G.left, U.right = G.hook(R), U.hook(G)
				U.black, G.black = true, false
				tr.hookSubTree(super, U)
				continue //变出黑U后再行解决
			} else {
				if R == nil || R.black {
					if L == nil || L.black { //双黑，变色解决
						U.black = false
						if G.black && super != nil {
							G = super
							continue //上溯
						}
						G.black = true
					} else { //中黑外红
						G.left, U.right = G.Hook(R), U.hook(G)
						U.black, G.black, L.black = G.black, true, true
						tr.hookSubTree(super, U)
					}
				} else { //中红
					U.right, G.left = U.Hook(R.left), G.Hook(R.right)
					R.left, R.right = R.hook(U), R.hook(G)
					R.black, G.black = G.black, true
					tr.hookSubTree(super, R)
				}
			}
		}
		break //个别情况需要循环
	}
}
