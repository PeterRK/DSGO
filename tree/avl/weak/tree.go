package weak

import (
	"constraints"
	"fmt"
)

const (
	Left  = 0
	Right = 1
)

//弱AVL树满足如下三条约束：
//nil的高度为-1
//当节点为叶节点（子节点皆为nil）时，高度必须为0
//节点和子节点的高度差为1或2
type node[T constraints.Ordered] struct {
	diff   [2]uint8
	key    T
	parent *node[T]
	kids   [2]*node[T]
}

func newNode[T constraints.Ordered](parent *node[T], key T) *node[T] {
	unit := new(node[T])
	unit.key = key
	unit.diff[Left], unit.diff[Right] = 1, 1
	unit.parent = parent
	//unit.kids = [2]*node[T]{nil, nil}
	return unit
}

func (root *node[T]) debug(indent int) {
	if root == nil {
		return
	}
	root.kids[Left].debug(indent + 1)
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
	fmt.Println(root.diff[Left], root.diff[Right], root.key)
	root.kids[Right].debug(indent + 1)
}

func (parent *node[T]) Hook(child *node[T]) *node[T] {
	if child != nil {
		child.parent = parent
	}
	return child
}
func (parent *node[T]) hook(child *node[T]) *node[T] {
	child.parent = parent
	return child
}

type Tree[T constraints.Ordered] struct {
	root *node[T]
	size int
}

func (tr *Tree[T]) Size() int {
	return tr.size
}

func (tr *Tree[T]) IsEmpty() bool {
	return tr.root == nil
}

func (tr *Tree[T]) Clear() {
	tr.root = nil
	tr.size = 0
}

func (tr *Tree[T]) Search(key T) bool {
	target := tr.root
	for target != nil {
		switch {
		case key < target.key:
			target = target.kids[Left]
		case key > target.key:
			target = target.kids[Right]
		default: //key == root.key
			return true
		}
	}
	return false
}

//成功返回true，冲突返回false。
//弱AVL树插入过程包括：O(logN)的搜索，O(1)的旋转，O(logN)的高度调整。
func (tr *Tree[T]) Insert(key T) bool {
	if tr.root == nil {
		tr.root = newNode(nil, key)
	} else {
		root, trace := tr.insert(key)
		if root == nil {
			return false
		}
		tr.rebalanceAfterInsert(root, trace)
	}
	tr.size++
	return true
}

//tr.root != nil
func (tr *Tree[T]) insert(key T) (*node[T], uint64) {
	root, depth, trace := tr.root, 0, uint64(0)
	for {
		depth++
		if depth > 64 {
			panic("too deep")
		}
		switch {
		case key < root.key:
			trace = (trace << 1) | Left
			if root.kids[Left] == nil {
				root.kids[Left] = newNode(root, key)
				return root, trace
			}
			root = root.kids[Left]
		case key > root.key:
			trace = (trace << 1) | Right
			if root.kids[Right] == nil {
				root.kids[Right] = newNode(root, key)
				return root, trace
			}
			root = root.kids[Right]
		default: //key == root.key
			return nil, trace
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
func (tr *Tree[T]) rebalanceAfterInsert(P *node[T], trace uint64) {
	for {
		Xside := trace & 1
		Sside := 1 - Xside
		if P.diff[Xside]--; P.diff[Xside] > 0 {
			break
		}
		super, root := P.parent, (*node[T])(nil)
		if P.diff[Sside] == 1 {
			P.diff[Xside], P.diff[Sside] = 1, 2
			if P = super; P == nil {
				break
			}
			trace >>= 1
			continue
		}
		if X := P.kids[Xside]; X.diff[Xside] == 1 {
			P.kids[Xside] = P.Hook(X.kids[Sside])
			X.kids[Sside] = X.hook(P)
			P.diff[Xside], P.diff[Sside], X.diff[Sside] = 1, 1, 1
			root = X
		} else {
			Z := X.kids[Sside]
			X.kids[Sside], P.kids[Xside] = X.Hook(Z.kids[Xside]), P.Hook(Z.kids[Sside])
			Z.kids[Xside], Z.kids[Sside] = Z.hook(X), Z.hook(P)
			X.diff[Sside], P.diff[Xside] = Z.diff[Xside], Z.diff[Sside]
			X.diff[Xside], P.diff[Sside] = 1, 1
			Z.diff[Xside], Z.diff[Sside] = 1, 1
			root = Z
		}
		if super == nil {
			tr.root = super.hook(root)
		} else {
			super.kids[(trace>>1)&1] = super.hook(root)
		}
		break
	}
}

//成功返回true，没有返回false。
//弱AVL树删除过程包括：O(logN)的搜索，O(1)的旋转，O(logN)的高度调整。
func (tr *Tree[T]) Remove(key T) bool {
	victim, orphan, trace := tr.findRemoveTarget(key)
	if victim == nil {
		return false
	}
	if root := victim.parent; root == nil {
		tr.root = root.Hook(orphan)
	} else {
		root.kids[trace&1] = root.Hook(orphan)
		tr.rebalanceAfterRemove(root, trace)
	}
	tr.size--
	return true
}

func (tr *Tree[T]) findRemoveTarget(key T) (victim, orphan *node[T], trace uint64) {
	target := tr.root
	for target != nil {
		switch {
		case key < target.key:
			trace = (trace << 1) | Left
			target = target.kids[Left]
		case key > target.key:
			trace = (trace << 1) | Right
			target = target.kids[Right]
		default: //key == root.key
			goto Lfound
		}
	}
	return nil, nil, 0
Lfound:
	switch {
	case target.kids[Left] == nil:
		victim, orphan = target, target.kids[Right]
	case target.kids[Right] == nil:
		victim, orphan = target, target.kids[Left]
	default:
		fst, snd := uint64(Left), uint64(Right)
		if target.diff[Left] > target.diff[Right] {
			fst, snd = Right, Left
		}
		trace = (trace << 1) | fst
		victim = target.kids[fst]
		for victim.kids[snd] != nil {
			trace = (trace << 1) | snd
			victim = victim.kids[snd]
		}
		orphan = victim.kids[fst]
		target.key = victim.key
	}
	return victim, orphan, trace
}

//------------直降-------------
//|    P       =       P      |
//|   / \      =      / \     |
//| S2   \     =    S1   \    |
//|       \    =   /  \   \   |
//|        X3  = Y2    Z2  X3 |

//----------------LR型----------------
//|       P        |        Z        |
//|      / \       |       / \       |
//|     S1  \      |      /   \      |
//|    / \   \     |     /     \     |
//|   /   Z1  \    |    S2      P2   |
//|  /   / \   \   |   /  \    / \   |
//| Y2  a   b   X3 |  Y1   a  b   X1 |

//----------------LL型----------------
//|       P        |       S         |
//|      / \       |      / \        |
//|    S1   \      |     /   \       |
//|   /  \   \     |    /     P1+    |
//| Y1    \   \    |  Y2     /  \    |
//|       Z?   \   |        Z?   \   |
//|             X3 |             X2- |

//P != nil
func (tr *Tree[T]) rebalanceAfterRemove(P *node[T], trace uint64) {
	if P.kids[Left] == nil && P.kids[Right] == nil { //叶节点需要特别处理
		P.diff[Left], P.diff[Right] = 1, 1
		if super := P.parent; super == nil {
			tr.root = super.Hook(P)
			return
		} else {
			P = super
			trace >>= 1
		}
	}
	for {
		Xside := trace & 1
		Sside := 1 - Xside
		if P.diff[Xside]++; P.diff[Xside] == 2 {
			break
		}
		super, root := P.parent, (*node[T])(nil)
		if P.diff[Sside] == 2 {
			P.diff[Sside], P.diff[Xside] = 1, 2
			goto Lcascade
		}
		if S := P.kids[Sside]; S.diff[Sside] == 2 {
			if S.diff[Xside] == 2 {
				S.diff[Sside], S.diff[Xside], P.diff[Xside] = 1, 1, 2
				goto Lcascade
			}
			Z := S.kids[Xside]
			S.kids[Xside], P.kids[Sside] = S.Hook(Z.kids[Sside]), P.Hook(Z.kids[Xside])
			Z.kids[Sside], Z.kids[Xside] = Z.hook(S), Z.hook(P)
			S.diff[Xside], P.diff[Sside] = Z.diff[Sside], Z.diff[Xside]
			S.diff[Sside], P.diff[Xside] = 1, 1
			Z.diff[Sside], Z.diff[Xside] = 2, 2
			root = Z
		} else {
			X, Z := P.kids[Xside], S.kids[Xside]
			S.kids[Xside] = S.hook(P)
			P.diff[Sside], S.diff[Xside] = S.diff[Xside], 1
			S.diff[Sside], P.diff[Xside] = 2, 2
			if Z == nil {
				P.kids[Sside] = nil
				if X == nil { //叶节点需要特别处理
					S.diff[Xside], P.diff[Sside], P.diff[Xside] = 2, 1, 1
				}
			} else {
				P.kids[Sside] = P.hook(Z)
			}
			root = S
		}
		if super == nil {
			tr.root = super.hook(root)
		} else {
			super.kids[(trace>>1)&1] = super.hook(root)
		}
		break
	Lcascade:
		if P = super; P == nil {
			break
		}
		trace >>= 1
	}
}
