package avlt

//--------------LR形式--------------
//|       G       |       C       |
//|      / \      |      / \      |
//|     P         |     P   G     |
//|    / \        |    / \ / \    |
//|       C       |      u v      |
//|      / \      |               |
//|     u   v     |               |

//--------------LL形式--------------
//|       G       |       P       |
//|      / \      |      / \      |
//|     P         |     C   G     |
//|    / \        |    / \ / \    |
//|   C   x       |        x      |
//|  / \          |               |
//|               |               |

//旋转后高度不发生变化时stop为true
func (G *node) rotate() (root *node, stop bool) {
	stop = false
	//root = nil
	if G.state == 2 { //左倾右旋
		var P = G.left
		if P.state == -1 { //LR
			var C = P.right //一定非nil
			P.right, G.left = C.left, C.right
			C.left, C.right = P, G
			switch C.state {
			case 1:
				G.state, P.state = -1, 0
			case -1:
				G.state, P.state = 0, 1
			default:
				G.state, P.state = 0, 0
			}
			C.state = 0
			root = C
		} else { //LL
			G.left, P.right = P.right, G
			if P.state == 0 { //不降高旋转
				G.state, P.state = 1, -1
				stop = true
			} else { //P.state == 1
				G.state, P.state = 0, 0
			}
			root = P
		}
	} else { //右倾左旋(P.state==-2)
		var P = G.right
		if P.state == 1 { //RL
			var C = P.left //一定非nil
			P.left, G.right = C.right, C.left
			C.right, C.left = P, G
			switch C.state {
			case -1:
				G.state, P.state = 1, 0
			case 1:
				G.state, P.state = 0, -1
			default:
				G.state, P.state = 0, 0
			}
			C.state = 0
			root = C
		} else { //RR
			G.right, P.left = P.left, G
			if P.state == 0 { //不降高旋转
				G.state, P.state = -1, 1
				stop = true
			} else { //P.state == -1
				G.state, P.state = 0, 0
			}
			root = P
		}
	}
	return root, stop
}
