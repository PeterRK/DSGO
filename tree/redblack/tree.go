package redblack

import (
	"golang.org/x/exp/constraints"
	"fmt"
	"unsafe"
)

const (
	Left    = 0
	Right   = 1
	Red     = 0
	Black   = 1
	HackBit = uintptr(1)
	//HackBit = ^(^uintptr(0) >> 1)
)

/*
type node[T constraints.Ordered] struct {
	key   T
	black  bool
	parent *node[T]
	kids  [2]*node[T]
}

func (unit *node[T]) setParent(parent *node[T]) {
	unit.parent = parent
}

func (unit *node[T]) getParent() *node[T] {
	return unit.parent
}

func (unit *node[T]) isRed() bool {
	return !unit.black
}

func (unit *node[T]) isBlack() bool {
	return unit.black
}

func (unit *node[T]) setRed() {
	unit.black = false
}

func (unit *node[T]) setBlack() {
	unit.black = true
}

func (unit *node[T]) copyColor(peer *node[T]) {
	unit.black = peer.black
}
*/

type node[T constraints.Ordered] struct {
	key   T
	trait uintptr
	kids  [2]*node[T]
}

func (unit *node[T]) setParent(parent *node[T]) {
	pt := uintptr(unsafe.Pointer(parent))
	if (pt & HackBit) != 0 {
		panic("unexpected high pointer")
	}
	unit.trait = pt | (unit.trait & HackBit)
}

func (unit *node[T]) getParent() *node[T] {
	pt := unit.trait & ^HackBit
	return (*node[T])(unsafe.Pointer(pt))
}

func (unit *node[T]) isRed() bool {
	return (unit.trait & HackBit) == 0
}

func (unit *node[T]) isBlack() bool {
	return (unit.trait & HackBit) != 0
}

func (unit *node[T]) setRed() {
	unit.trait &= ^HackBit
}

func (unit *node[T]) setBlack() {
	unit.trait |= HackBit
}

func (unit *node[T]) copyColor(peer *node[T]) {
	mark := peer.trait & HackBit
	unit.trait &= ^HackBit
	unit.trait |= mark
}

func newNode[T constraints.Ordered](parent *node[T], key T) (unit *node[T]) {
	unit = new(node[T])
	unit.key = key
	//unit.setRead()
	unit.setParent(parent)
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
	mark := "R"
	if root.isBlack() {
		mark = "B"
	}
	fmt.Println(mark, root.key)
	root.kids[Right].debug(indent + 1)
}

func (parent *node[T]) Hook(child *node[T]) *node[T] {
	if child != nil {
		child.setParent(parent)
	}
	return child
}
func (parent *node[T]) hook(child *node[T]) *node[T] {
	child.setParent(parent)
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
func (tr *Tree[T]) Insert(key T) bool {
	if tr.root == nil {
		tr.root = newNode(nil, key)
		tr.root.setBlack()
	} else {
		parent, trace := tr.insert(key)
		if parent == nil {
			return false
		}
		tr.rebalanceAfterInsert(parent, trace)
	}
	tr.size++
	return true
}

//tr.root != nil
func (tr *Tree[T]) insert(key T) (*node[T], uint64) {
	parent, depth, trace := tr.root, 0, uint64(0)
	for {
		depth++
		if depth > 64 {
			panic("too deep")
		}
		switch {
		case key < parent.key:
			trace = (trace << 1) | Left
			if parent.kids[Left] == nil {
				parent.kids[Left] = newNode(parent, key) //默认为红
				return parent, trace
			}
			parent = parent.kids[Left]
		case key > parent.key:
			trace = (trace << 1) | Right
			if parent.kids[Right] == nil {
				parent.kids[Right] = newNode(parent, key) //默认为红
				return parent, trace
			}
			parent = parent.kids[Right]
		default: //key == root.key
			return nil, trace
		}
	}
}

//------------红叔模式------------
//|      bG      |      rG      |
//|     /  \     |     /  \     |
//|   rP    rU   |   bP    bU   |
//|   |          |   |          |
//|   rC         |   rC         |

//------------------LL型-----------------
//|        bG        |        bP        |
//|       /  \       |       /  \       |
//|     rP    bU     |     rC     rG    |
//|    /  \          |          /  \    |
//|  rC    x         |         x    bU  |

//------------------LR型-----------------
//|        bG        |        bC        |
//|       /  \       |       /  \       |
//|     rP    bU     |     rP    rG     |
//|    / \           |    / \    / \    |
//|      rC          |       u  v   bU  |
//|     /  \         |                  |
//|    u    v        |                  |

func (tr *Tree[T]) rebalanceAfterInsert(P *node[T], trace uint64) {
	for P.isRed() { //违反双红禁
		G := P.getParent() //必然存在，根为黑，P非根
		super := G.getParent()
		Pside := (trace >> 1) & 1
		Uside := 1 - Pside
		U := G.kids[Uside]
		if U != nil && U.isRed() { //红叔模式，变色解决
			P.setBlack()
			U.setBlack()
			if super != nil {
				G.setRed()
				P = super
				trace >>= 2
				continue //上溯，检查双红禁
			} //遇根终止
		} else { //黑叔模式，旋转解决
			var root *node[T]
			if (trace & 1) == Pside {
				G.kids[Pside] = G.Hook(P.kids[Uside])
				P.kids[Uside] = P.hook(G)
				root = P
			} else {
				C := P.kids[Uside]
				P.kids[Uside] = P.Hook(C.kids[Pside])
				G.kids[Pside] = G.Hook(C.kids[Uside])
				C.kids[Pside], C.kids[Uside] = C.hook(P), C.hook(G)
				root = C
			}
			G.setRed()
			root.setBlack()
			if super == nil {
				tr.root = super.hook(root)
			} else {
				super.kids[(trace>>2)&1] = super.hook(root)
			}
		}
		break
	}
}

//成功返回true，没有返回false。
func (tr *Tree[T]) Remove(key T) bool {
	victim, orphan, trace := tr.findRemoveTarget(key)
	if victim == nil {
		return false
	}
	if root := victim.getParent(); root == nil {
		if orphan != nil {
			orphan.setBlack()
			orphan.setParent(nil)
		}
		tr.root = orphan
	} else {
		root.kids[trace&1] = root.Hook(orphan)
		if victim.isBlack() { //红victim随便删，黑的要考虑
			if orphan != nil && orphan.isRed() {
				orphan.setBlack() //红子变黑顶上
			} else {
				tr.rebalanceAfterRemove(root, trace)
			}
		}
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
		trace = (trace << 1) | Right
		victim = target.kids[Right]
		for victim.kids[Left] != nil {
			trace = (trace << 1) | Left
			victim = victim.kids[Left]
		}
		orphan = victim.kids[Right]
		target.key = victim.key
	}
	return victim, orphan, trace
}

//----------------红叔模式----------------
//|        bG        |        bU        |
//|       /  \       |       /  \       |
//|     bO    rU     |     rG    bZ     |
//|          /  \    |    /  \          |
//|        bY    bZ  |  bO    bY        |

//------------------双黑------------------
//|        xG        |        bG        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bO    rU     |
//|          /  \    |          /  \    |
//|        bY    bZ  |        bY    bZ  |

//----------------中黑外红----------------
//|        xG        |        xU        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bG    bZ     |
//|          /  \    |    /  \          |
//|        bY    rZ  |  bO    bY        |

//------------------中红------------------
//|        xG        |        xY        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bG    bU     |
//|          /  \    |    /  \  /  \    |
//|        rY    xZ  |  bO   u  v   xZ  |
//|       /  \       |                  |
//|      u    v      |                  |

func (tr *Tree[T]) rebalanceAfterRemove(G *node[T], trace uint64) {
	for { //剩下情况：victim黑，orphan也黑，此时victim(orphan顶替)的兄弟必然存在
		super := G.getParent()
		Oside := trace & 1
		Uside := 1 - Oside
		U := G.kids[Uside]
		Y, Z := U.kids[Oside], U.kids[Uside]
		if U.isRed() { //红U下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
			G.kids[Uside] = G.hook(Y)
			U.kids[Oside] = U.hook(G)
			U.setBlack()
			G.setRed()
			if super == nil {
				tr.root = super.hook(U)
			} else {
				super.kids[(trace>>1)&1] = super.hook(U)
			}
			trace = (trace << 1) | Oside
			continue //变出黑U后再行解决
		}
		var root *node[T]
		if Y == nil || Y.isBlack() {
			if Z == nil || Z.isBlack() { //双黑，变色解决
				U.setRed()
				if G.isRed() {
					G.setBlack()
				} else if super != nil {
					G = super
					trace >>= 1
					continue //上溯
				}
				break
			}
			G.kids[Uside] = G.Hook(Y)
			U.kids[Oside] = U.hook(G)
			U.copyColor(G)
			G.setBlack()
			Z.setBlack()
			root = U //中黑外红
		} else { //中红
			G.kids[Uside] = G.Hook(Y.kids[Oside])
			U.kids[Oside] = U.Hook(Y.kids[Uside])
			Y.kids[Oside], Y.kids[Uside] = Y.hook(G), Y.hook(U)
			Y.copyColor(G)
			G.setBlack()
			root = Y
		}
		if super == nil {
			tr.root = super.hook(root)
		} else {
			super.kids[(trace>>1)&1] = super.hook(root)
		}
		break //个别情况需要循环
	}
}
