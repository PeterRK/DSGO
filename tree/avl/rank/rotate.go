package rank

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
func (G *node[T]) rotate() (root *node[T], stop bool) {
	stop = false
	//root = nil
	if G.state == 2 { //左倾右旋
		P := G.left
		if P.state == -1 { //LR
			C := P.right //一定非nil
			w := C.right.realWeight()
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
			root = C

			C.weight = G.weight
			G.weight -= P.weight - w
			P.weight -= w + 1
		} else { //LL
			w := P.right.realWeight()
			G.left, P.right = G.tryHook(P.right), P.hook(G)

			if P.state == 1 { //真LL
				G.state, P.state = 0, 0
			} else { //伪LL，保持高度
				stop = true
				G.state, P.state = 1, -1
			}
			root = P

			p := P.weight
			P.weight = G.weight
			G.weight -= p - w
		}
	} else { //右倾左旋(P.state==-2)
		P := G.right
		if P.state == 1 { //RL
			C := P.left //一定非nil
			w := C.left.realWeight()
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
			root = C

			C.weight = G.weight
			G.weight -= P.weight - w
			P.weight -= w + 1
		} else { //RR
			w := P.left.realWeight()
			G.right, P.left = G.tryHook(P.left), P.hook(G)

			if P.state == -1 { //真RR
				G.state, P.state = 0, 0
			} else { //伪RR，保持高度
				stop = true
				G.state, P.state = -1, 1
			}
			root = P

			p := P.weight
			P.weight = G.weight
			G.weight -= p - w
		}
	}
	return root, stop
}
