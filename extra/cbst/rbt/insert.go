package rbt

//成功返回true，冲突返回false。
//红黑树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Insert(key int32) bool {
	if tr.root == nil {
		tr.root = newNode(key) //默认为红
		tr.root.black = true
	} else {
		root := tr.insert(key)
		if root == nil {
			return false
		}
		tr.rebalanceAfterInsert()
	}
	return true
}

//插入节点，root != nil
func (tr *Tree) insert(key int32) *node {
	tr.path.clear()
	for root := tr.root; ; {
		switch {
		case key < root.key:
			tr.path.push(root, true)
			if root.left == nil {
				root.left = newNode(key) //默认为红
				return root
			}
			root = root.left
		case key > root.key:
			tr.path.push(root, false)
			if root.right == nil {
				root.right = newNode(key) //默认为红
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

func (tr *Tree) rebalanceAfterInsert() {
	P, klf := tr.path.pop()
	for !P.black { //违法双红禁
		G, plf := tr.path.pop() //必然存在，根为黑，P非根
		if plf {
			U := G.right
			if !U.isBlack() { //红叔模式，变色解决
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
					C := P.right
					P.right, G.left = C.left, C.right
					C.left, C.right = P, G
					G.black, C.black = false, true
					tr.hookSubTree(C)
				}
			}
		} else {
			U := G.left
			if !U.isBlack() { //红叔模式，变色解决
				P.black, U.black = true, true
				if !tr.path.isEmpty() {
					G.black = false
					P, klf = tr.path.pop()
					continue //上溯，检查双红禁
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if klf { //RL
					C := P.left
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
	unit.key = key
	//unit.black = false
	//unit.left, unit.right = nil, nil
	return unit
}
