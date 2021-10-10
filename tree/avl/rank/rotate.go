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
			w := C.right.Weight()
			P.right, G.left = P.Hook(C.left), G.Hook(C.right)
			C.left, C.right = C.hook(P), C.hook(G)

			G.state, P.state = 0, 0
			if C.state > 0 {
				G.state = -1
			}
			if C.state < 0 {
				P.state = 1
			}
			C.state = 0
			root = C

			C.weight = G.weight
			G.weight -= P.weight - w
			P.weight -= w + 1
		} else { //LL
			w := P.right.Weight()
			G.left, P.right = G.Hook(P.right), P.hook(G)

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
			w := C.left.Weight()
			P.left, G.right = P.Hook(C.right), G.Hook(C.left)
			C.right, C.left = C.hook(P), C.hook(G)

			G.state, P.state = 0, 0
			if C.state < 0 {
				G.state = 1
			}
			if C.state > 0 {
				P.state = -1
			}
			C.state = 0
			root = C

			C.weight = G.weight
			G.weight -= P.weight - w
			P.weight -= w + 1
		} else { //RR
			w := P.left.Weight()
			G.right, P.left = G.Hook(P.left), P.hook(G)

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
