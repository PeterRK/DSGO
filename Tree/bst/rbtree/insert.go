package rbtree

//成功返回true，冲突返回false。
//红黑树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Insert(key int32) bool {
	if tr.root == nil {
		tr.root = newNode(nil, key) //默认为红
		tr.root.black = true
	} else {
		var root = tr.insert(key)
		if root == nil {
			return false
		}
		tr.rebalanceAfterInsert(root, key)
	}
	return true
}

//插入节点，root != nil
func (tr *Tree) insert(key int32) *node {
	var root = tr.root
	for {
		switch {
		case key < root.key:
			if root.left == nil {
				root.left = newNode(root, key) //默认为红
				return root
			}
			root = root.left
		case key > root.key:
			if root.right == nil {
				root.right = newNode(root, key) //默认为红
				return root
			}
			root = root.right
		default: //key == root.key
			return nil
		}
	}
}

//------------红叔模式------------
//|      bG      |      rG      |
//|     /  \     |     /  \     |
//|   rP    rU   |   bP    bU   |
//|   |          |   |          |
//|   rC         |   rC         |

//------------------LL型-----------------
//|        bG        |        bP        |
//|       /  \       |       /  \       |
//|     rP    bU     |     rC     rG    |
//|    /  \          |          /  \    |
//|  rC    x         |         x    bU  |

//------------------LR型-----------------
//|        bG        |        bC        |
//|       /  \       |       /  \       |
//|     rP    bU     |     rP    rG     |
//|    / \           |    / \    / \    |
//|      rC          |       u  v   bU  |
//|     /  \         |                  |
//|    u    v        |                  |
func (tr *Tree) rebalanceAfterInsert(P *node, key int32) {
	for !P.black { //违反双红禁
		var G = P.parent //必然存在，根为黑，P非根
		var super = G.parent
		if key < G.key {
			var U = G.right
			if U != nil && !U.black { //红叔模式，变色解决
				P.black, U.black = true, true
				if super != nil {
					G.black = false
					P = G.parent
					continue //上溯，检查双红禁
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if key < P.key { //LL
					G.left, P.right = G.tryHook(P.right), P.hook(G)
					G.black, P.black = false, true
					tr.hookSubTree(super, P)
				} else { //LR
					var C = P.right
					P.right, G.left = P.tryHook(C.left), G.tryHook(C.right)
					C.left, C.right = C.hook(P), C.hook(G)
					G.black, C.black = false, true
					tr.hookSubTree(super, C)
				}
			}
		} else {
			var U = G.left
			if U != nil && !U.black { //红叔模式，变色解决
				P.black, U.black = true, true
				if super != nil {
					G.black = false
					P = G.parent
					continue //上溯，检查双红禁
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if key > P.key { //RR
					G.right, P.left = G.tryHook(P.left), P.hook(G)
					G.black, P.black = false, true
					tr.hookSubTree(super, P)
				} else { //RL
					var C = P.left
					P.left, G.right = P.tryHook(C.right), G.tryHook(C.left)
					C.right, C.left = C.hook(P), C.hook(G)
					G.black, C.black = false, true
					tr.hookSubTree(super, C)
				}
			}
		}
		break //变色时才需要循环
	}
}

func newNode(parent *node, key int32) (unit *node) {
	unit = new(node)
	unit.key, unit.black = key, false
	unit.parent = parent
	unit.left, unit.right = nil, nil
	return unit
}
