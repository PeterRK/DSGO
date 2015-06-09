package rbt

//成功返回true，冲突返回false。
//红黑树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Insert(key int32) bool {
	if tr.root == nil {
		tr.root = newNode(key) //默认为红
		tr.root.black = true
		return true
	}
	tr.path.clear()
	for root := tr.root; ; {
		if key < root.key {
			tr.path.push(root, true)
			if root.left == nil {
				root.left = newNode(key) //默认为红
				break
			}
			root = root.left
		} else if key > root.key {
			tr.path.push(root, false)
			if root.right == nil {
				root.right = newNode(key) //默认为红
				break
			}
			root = root.right
		} else { //key == root.key
			return false
		}
	}
	tr.adjustAfterInsert()
	return true
}

//------------红叔模式------------
//|      bG      |      rG      |
//|     /  \     |     /  \     |
//|   rP    rU   |   bP    bU   |
//|   |          |   |          |
//|   rC         |   rC         |

//-----------------LL形式-----------------
//|        bG        |        bP        |
//|       /  \       |       /  \       |
//|     rP    bU     |     rC     rG    |
//|    /  \          |          /  \    |
//|  rC    x         |         x    bU  |

//-----------------LR形式-----------------
//|        bG        |        bC        |
//|       /  \       |       /  \       |
//|     rP    bU     |     rP    rG     |
//|    / \           |    / \    / \    |
//|      rC          |       u  v   bU  |
//|     /  \         |                  |
//|    u    v        |                  |

func (tr *Tree) adjustAfterInsert() {
	var P, klf = tr.path.pop()
	for !P.black { //违法双红禁
		var G, plf = tr.path.pop() //必然存在，根为黑，P非根
		if plf {
			var U = G.right
			if U != nil && !U.black { //红叔模式，变色解决
				P.black, U.black = true, true
				if !tr.path.isEmpty() {
					G.black = false
					P, klf = tr.path.pop()
					continue //上溯，检查双红禁
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if klf { //LL
					G.left, P.right = P.right, G
					G.black, P.black = false, true
					tr.hookSubTree(P)
				} else { //LR
					var C = P.right
					P.right, G.left = C.left, C.right
					C.left, C.right = P, G
					G.black, C.black = false, true
					tr.hookSubTree(C)
				}
			}
		} else {
			var U = G.left
			if U != nil && !U.black { //红叔模式，变色解决
				P.black, U.black = true, true
				if !tr.path.isEmpty() {
					G.black = false
					P, klf = tr.path.pop()
					continue //上溯，检查双红禁
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if klf { //RL
					var C = P.left
					P.left, G.right = C.right, C.left
					C.right, C.left = P, G
					G.black, C.black = false, true
					tr.hookSubTree(C)
				} else { //RR
					G.right, P.left = P.left, G
					G.black, P.black = false, true
					tr.hookSubTree(P)
				}
			}
		}
		break //变色时才需要循环
	}
}

func newNode(key int32) (unit *node) {
	unit = new(node)
	unit.key, unit.black = key, false
	unit.left, unit.right = nil, nil
	return unit
}
