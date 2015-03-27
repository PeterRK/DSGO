package avltree

func (parent *node) rotate() (root *node, keep_height bool) {
	keep_height = false
	//	root = nil
	if parent.balance == 2 { //左倾右旋
		var child = parent.left
		if child.balance == -1 { //LR
			var grandchild = child.right //一定非nil
			child.right, parent.left = grandchild.left, grandchild.right
			grandchild.left, grandchild.right = child, parent
			switch grandchild.balance {
			case -1:
				parent.balance, child.balance = 0, 1
			case 1:
				parent.balance, child.balance = -1, 0
			default:
				parent.balance, child.balance = 0, 0
			}
			grandchild.balance = 0
			root = grandchild
		} else { //LL
			parent.left, child.right = child.right, parent
			if child.balance == 0 { //不降高旋转
				parent.balance, child.balance = 1, -1
				keep_height = true
			} else { //child.balance == 1
				parent.balance, child.balance = 0, 0
			}
			root = child
		}
	} else { //右倾左旋(parent.balance==-2)
		var child = parent.right
		if child.balance == 1 { //RL
			var grandchild = child.left //一定非nil
			child.left, parent.right = grandchild.right, grandchild.left
			grandchild.right, grandchild.left = child, parent
			switch grandchild.balance {
			case -1:
				parent.balance, child.balance = 1, 0
			case 1:
				parent.balance, child.balance = 0, -1
			default:
				parent.balance, child.balance = 0, 0
			}
			grandchild.balance = 0
			root = grandchild
		} else { //RR
			parent.right, child.left = child.left, parent
			if child.balance == 0 { //不降高旋转
				parent.balance, child.balance = -1, 1
				keep_height = true
			} else { //child.balance == -1
				parent.balance, child.balance = 0, 0
			}
			root = child
		}
	}
	return
}
