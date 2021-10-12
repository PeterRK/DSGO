package avl

import (
	"constraints"
	"fmt"
)

const (
	Left  = 0
	Right = 1
)

type node[T constraints.Ordered] struct {
	key    T
	state  int8 //(-2), -1, 0, 1, (2)
	parent *node[T]
	kids   [2]*node[T]
}

func newNode[T constraints.Ordered](parent *node[T], key T) (unit *node[T]) {
	unit = new(node[T])
	unit.key = key
	//unit.state = 0
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
	fmt.Println(root.state, root.key)
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
func (G *node[T]) rotate() (*node[T], bool) {
	//G.state == -2 || G.state == 2
	Pside := (G.state + 2) / 4
	Uside := 1 - Pside
	P := G.kids[Pside]
	if direct := G.state / 2; P.state == -direct { //LR(RL)
		C := P.kids[Uside]
		P.kids[Uside] = P.Hook(C.kids[Pside])
		G.kids[Pside] = G.Hook(C.kids[Uside])
		C.kids[Pside], C.kids[Uside] = C.hook(P), C.hook(G)
		G.state, P.state = 0, 0
		if C.state == direct {
			G.state = -direct
		} else if C.state == -direct {
			P.state = direct
		}
		C.state = 0
		return C, false
	} else { //LL(RR)
		G.kids[Pside] = G.Hook(P.kids[Uside])
		P.kids[Uside] = P.hook(G)
		if P.state == direct { //真LL(RR)
			G.state, P.state = 0, 0
			return P, false
		} else { //伪LL(RR)，保持高度
			G.state, P.state = direct, -direct
			return P, true
		}
	}
}

//成功返回true，冲突返回false。
//AVL树插入过程包括：O(logN)的搜索，O(1)的旋转，O(logN)的平衡因子调整。
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

func (tr *Tree[T]) rebalanceAfterInsert(root *node[T], trace uint64) {
	for {
		state := root.state
		root.state += int8(trace&1)*2 - 1
		if state == 0 {
			if root.parent != nil {
				root = root.parent
				trace >>= 1
				continue
			}
		} else if root.state != 0 {
			super := root.parent
			root, _ = root.rotate()
			if super == nil {
				tr.root = super.hook(root)
			} else {
				super.kids[(trace>>1)&1] = super.hook(root)
			}
		}
		break
	}
}

//成功返回true，没有返回false。
//AVL树删除过程包括：O(logN)的搜索，O(logN)的旋转，O(logN)的平衡因子调整。
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
		if target.state > 0 {
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

//root != nil
func (tr *Tree[T]) rebalanceAfterRemove(root *node[T], trace uint64) {
	state, stop := root.state, false
	root.state -= int8(trace&1)*2 - 1
	for state != 0 { //如果原平衡因子为0则子树高度不变
		super := root.parent
		if super == nil {
			if root.state != 0 { //2 || -2
				root, _ = root.rotate()
				tr.root = super.hook(root)
			}
			break
		}
		if root.state != 0 { //2 || -2
			root, stop = root.rotate()
			super.kids[(trace>>1)&1] = super.hook(root)
			if stop {
				break
			}
		}
		root, state = super, super.state
		trace >>= 1
		root.state -= int8(trace&1)*2 - 1
	}
}
