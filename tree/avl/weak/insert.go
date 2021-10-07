package weak

//成功返回true，冲突返回false。
//弱AVL树插入过程包括：O(logN)的搜索，O(1)的旋转，O(logN)的高度调整。
func (tr *Tree[T]) Insert(key T) bool {
	if tr.root == nil {
		tr.root = newNode(nil, key)
	} else {
		root := tr.insert(key)
		if root == nil {
			return false
		}
		tr.rebalanceAfterInsert(root, key)
	}
	tr.size++
	return true
}

//插入节点，root != nil
func (tr *Tree[T]) insert(key T) *node[T] {
	root := tr.root
	for {
		switch {
		case key < root.key:
			if root.left == nil {
				root.left = newNode[T](root, key)
				return root
			}
			root = root.left
		case key > root.key:
			if root.right == nil {
				root.right = newNode[T](root, key)
				return root
			}
			root = root.right
		default: //key == root.key
			return nil
		}
	}
}

//-------------直升-------------
//|             |       P      |
//|             |      / \     |
//|   X0 - P    |    X1   \    |
//|         \   |          \   |
//|          S1 |           S2 |

//----------------LL型---------------
//|    X0 - P      |       X        |
//|   / \    \     |      / \       |
//| Y1   \    \    |    Y1   P1     |
//|       \    \   |        /  \    |
//|        Z2   S2 |      Z1    S1  |

//-----------------LR型----------------
//|      X0 - P     |        Z        |
//|     / \    \    |      /   \      |
//|    /   Z1   \   |    X1     P1    |
//|   /   /  \   \  |   /  \   /  \   |
//| Y2   a    b  S2 | Y1    a b    S1 |

//P != nil
func (tr *Tree[T]) rebalanceAfterInsert(P *node[T], key T) {
	for {
		super, root := P.parent, (*node[T])(nil)
		if key < P.key {
			X, S := P.left, P.right
			if X.height < P.height {
				break
			}
			if P.height-S.Height() == 1 {
				P.height++
				goto Lup
			}
			//P.height - S.Height() == 2
			Y, Z := X.left, X.right
			if X.height-Y.Height() == 1 {
				P.height--
				X.right = X.hook(P)
				P.left = P.Hook(Z)
				root = X
			} else {
				Z.height++
				X.height--
				P.height--
				X.right, P.left = X.Hook(Z.left), P.Hook(Z.right)
				Z.left, Z.right = Z.hook(X), Z.hook(P)
				root = Z
			}
		} else {
			X, S := P.right, P.left
			if X.height < P.height {
				break
			}
			if P.height-S.Height() == 1 {
				P.height++
				goto Lup
			}
			//P.height - S.Height() == 2
			Y, Z := X.right, X.left
			if X.height-Y.Height() == 1 {
				P.height--
				X.left = X.hook(P)
				P.right = P.Hook(Z)
				root = X
			} else {
				Z.height++
				X.height--
				P.height--
				X.left, P.right = X.Hook(Z.right), P.Hook(Z.left)
				Z.right, Z.left = Z.hook(X), Z.hook(P)
				root = Z
			}
		}
		tr.hookSubTree(super, root)
		break
	Lup:
		if P = super; P == nil {
			break
		}
	}
}
