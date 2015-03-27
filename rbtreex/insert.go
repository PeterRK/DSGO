//红黑树的实现
package rbtreex

func newNode(key int) (unit *node) {
	unit = new(node)
	unit.key, unit.black = key, false
	unit.left, unit.right = NULL, NULL
	return
}

//红黑树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整
func (tree *Tree) Insert(key int) bool {
	if tree.root == NULL {
		tree.root = newNode(key) //默认为红
		tree.root.black = true
		return true
	}

	tree.path.clear()
	var parent = tree.root
	var kid *node
	var kid_is_left bool
	for {
		if key == parent.key {
			return false
		}
		if key < parent.key {
			kid_is_left = true
			if parent.left == NULL {
				kid = newNode(key) //默认为红
				parent.left = kid
				break
			}
			tree.path.push(parent, kid_is_left)
			parent = parent.left
		} else {
			kid_is_left = false
			if parent.right == NULL {
				kid = newNode(key) //默认为红
				parent.right = kid
				break
			}
			tree.path.push(parent, kid_is_left)
			parent = parent.right
		}
	}

	for !parent.black { //违法双红禁
		var grandparent, parent_is_left = tree.path.pop() //必然存在，根为黑，parent非根
		if parent_is_left {
			var another = grandparent.right
			if !another.black { //红叔模式，变色解决
				parent.black, another.black = true, true
				if !tree.path.isEmpty() {
					grandparent.black = false
					kid = grandparent
					parent, kid_is_left = tree.path.pop() //上溯，检查双红禁
					continue
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if kid_is_left { //LL
					grandparent.left, parent.right = parent.right, grandparent
					grandparent.black, parent.black = false, true
					tree.hookSubTree(parent)
				} else { //LR
					parent.right, grandparent.left = kid.left, kid.right
					kid.left, kid.right = parent, grandparent
					grandparent.black, kid.black = false, true
					tree.hookSubTree(kid)
				}
			}
		} else {
			var another = grandparent.left
			if !another.black { //红叔模式，变色解决
				parent.black = true
				another.black = true
				if !tree.path.isEmpty() {
					grandparent.black = false
					kid = grandparent
					parent, kid_is_left = tree.path.pop() //上溯，检查双红禁
					continue
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if kid_is_left { //RL
					parent.left, grandparent.right = kid.right, kid.left
					kid.right, kid.left = parent, grandparent
					grandparent.black, kid.black = false, true
					tree.hookSubTree(kid)
				} else { //RR
					grandparent.right, parent.left = parent.left, grandparent
					grandparent.black, parent.black = false, true
					tree.hookSubTree(parent)
				}
			}
		}
		break //变色时才需要循环
	}
	return true
}
