package wavl

//---------------LR型--------------
//|       G       |       C       |
//|      / \      |      / \      |
//|     P         |     P   G     |
//|    / \        |    / \ / \    |
//|       C       |      u v      |
//|      / \      |               |
//|     u   v     |               |

//---------------LL型--------------
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
			var v = C.right.realWeight()
			P.right, G.left = P.tryHook(C.left), G.tryHook(C.right)
			C.left, C.right = C.hook(P), C.hook(G)

			switch C.state {
			case 1:
				G.state, P.state = -1, 0
			case -1:
				G.state, P.state = 0, 1
			default:
				G.state, P.state = 0, 0
			}
			C.state = 0

			C.weight = G.weight
			G.weight -= P.weight - v
			P.weight -= v + 1
			root = C
		} else { //LL
			var x = P.right.realWeight()
			G.left, P.right = G.tryHook(P.right), P.hook(G)

			if P.state == 1 { //真LL
				G.state, P.state = 0, 0
			} else { //伪LL，保持高度
				stop = true
				G.state, P.state = 1, -1
			}

			var p = P.weight
			P.weight = G.weight
			G.weight -= p - x
			root = P
		}
	} else { //右倾左旋(P.state==-2)
		var P = G.right
		if P.state == 1 { //RL
			var C = P.left //一定非nil
			var v = C.left.realWeight()
			P.left, G.right = P.tryHook(C.right), G.tryHook(C.left)
			C.right, C.left = C.hook(P), C.hook(G)

			switch C.state {
			case -1:
				G.state, P.state = 1, 0
			case 1:
				G.state, P.state = 0, -1
			default:
				G.state, P.state = 0, 0
			}
			C.state = 0

			C.weight = G.weight
			G.weight -= P.weight - v
			P.weight -= v + 1
			root = C
		} else { //RR
			var x = P.left.realWeight()
			G.right, P.left = G.tryHook(P.left), P.hook(G)

			if P.state == -1 { //真RR
				G.state, P.state = 0, 0
			} else { //伪RR，保持高度
				stop = true
				G.state, P.state = -1, 1
			}

			var p = P.weight
			P.weight = G.weight
			G.weight -= p - x
			root = P
		}
	}
	return root, stop
}
