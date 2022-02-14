package skiplist

import (
	"golang.org/x/exp/constraints"
)

func newLeaf[T constraints.Ordered](level int) *lNode[T] {
	switch level {
	case 1:
		node1 := new(lNode1[T])
		node1.next = node1.space[:]
		return &node1.lNode
	case 2:
		node2 := new(lNode2[T])
		node2.next = node2.space[:]
		return &node2.lNode
	case 3:
		node3 := new(lNode3[T])
		node3.next = node3.space[:]
		return &node3.lNode
	case 4:
		node4 := new(lNode4[T])
		node4.next = node4.space[:]
		return &node4.lNode
	case 5:
		node5 := new(lNode5[T])
		node5.next = node5.space[:]
		return &node5.lNode
	case 6:
		node6 := new(lNode6[T])
		node6.next = node6.space[:]
		return &node6.lNode
	case 7:
		node7 := new(lNode7[T])
		node7.next = node7.space[:]
		return &node7.lNode
	case 8:
		node8 := new(lNode8[T])
		node8.next = node8.space[:]
		return &node8.lNode
	case 9:
		node9 := new(lNode9[T])
		node9.next = node9.space[:]
		return &node9.lNode
	case 10:
		node10 := new(lNode10[T])
		node10.next = node10.space[:]
		return &node10.lNode
	case 11:
		node11 := new(lNode11[T])
		node11.next = node11.space[:]
		return &node11.lNode
	case 12:
		node12 := new(lNode12[T])
		node12.next = node12.space[:]
		return &node12.lNode
	case 13:
		node13 := new(lNode13[T])
		node13.next = node13.space[:]
		return &node13.lNode
	case 14:
		node14 := new(lNode14[T])
		node14.next = node14.space[:]
		return &node14.lNode
	case 15:
		node15 := new(lNode15[T])
		node15.next = node15.space[:]
		return &node15.lNode
	case 16:
		node16 := new(lNode16[T])
		node16.next = node16.space[:]
		return &node16.lNode
	case 17:
		node17 := new(lNode17[T])
		node17.next = node17.space[:]
		return &node17.lNode
	case 18:
		node18 := new(lNode18[T])
		node18.next = node18.space[:]
		return &node18.lNode
	case 19:
		node19 := new(lNode19[T])
		node19.next = node19.space[:]
		return &node19.lNode
	case 20:
		node20 := new(lNode20[T])
		node20.next = node20.space[:]
		return &node20.lNode
	case 21:
		node21 := new(lNode21[T])
		node21.next = node21.space[:]
		return &node21.lNode
	case 22:
		node22 := new(lNode22[T])
		node22.next = node22.space[:]
		return &node22.lNode
	case 23:
		node23 := new(lNode23[T])
		node23.next = node23.space[:]
		return &node23.lNode
	case 24:
		node24 := new(lNode24[T])
		node24.next = node24.space[:]
		return &node24.lNode
	case 25:
		node25 := new(lNode25[T])
		node25.next = node25.space[:]
		return &node25.lNode
	case 26:
		node26 := new(lNode26[T])
		node26.next = node26.space[:]
		return &node26.lNode
	case 27:
		node27 := new(lNode27[T])
		node27.next = node27.space[:]
		return &node27.lNode
	case 28:
		node28 := new(lNode28[T])
		node28.next = node28.space[:]
		return &node28.lNode
	case 29:
		node29 := new(lNode29[T])
		node29.next = node29.space[:]
		return &node29.lNode
	case 30:
		node30 := new(lNode30[T])
		node30.next = node30.space[:]
		return &node30.lNode
	default:
		node := new(lNode[T])
		node.next = make([]*bNode[T], level)
		return node
	}
}

type lNode1[T constraints.Ordered] struct {
	lNode[T]
	space [1]*bNode[T]
}
type lNode2[T constraints.Ordered] struct {
	lNode[T]
	space [2]*bNode[T]
}
type lNode3[T constraints.Ordered] struct {
	lNode[T]
	space [3]*bNode[T]
}
type lNode4[T constraints.Ordered] struct {
	lNode[T]
	space [4]*bNode[T]
}
type lNode5[T constraints.Ordered] struct {
	lNode[T]
	space [5]*bNode[T]
}
type lNode6[T constraints.Ordered] struct {
	lNode[T]
	space [6]*bNode[T]
}
type lNode7[T constraints.Ordered] struct {
	lNode[T]
	space [7]*bNode[T]
}
type lNode8[T constraints.Ordered] struct {
	lNode[T]
	space [8]*bNode[T]
}
type lNode9[T constraints.Ordered] struct {
	lNode[T]
	space [9]*bNode[T]
}
type lNode10[T constraints.Ordered] struct {
	lNode[T]
	space [10]*bNode[T]
}
type lNode11[T constraints.Ordered] struct {
	lNode[T]
	space [11]*bNode[T]
}
type lNode12[T constraints.Ordered] struct {
	lNode[T]
	space [12]*bNode[T]
}
type lNode13[T constraints.Ordered] struct {
	lNode[T]
	space [13]*bNode[T]
}
type lNode14[T constraints.Ordered] struct {
	lNode[T]
	space [14]*bNode[T]
}
type lNode15[T constraints.Ordered] struct {
	lNode[T]
	space [15]*bNode[T]
}
type lNode16[T constraints.Ordered] struct {
	lNode[T]
	space [16]*bNode[T]
}
type lNode17[T constraints.Ordered] struct {
	lNode[T]
	space [17]*bNode[T]
}
type lNode18[T constraints.Ordered] struct {
	lNode[T]
	space [18]*bNode[T]
}
type lNode19[T constraints.Ordered] struct {
	lNode[T]
	space [19]*bNode[T]
}
type lNode20[T constraints.Ordered] struct {
	lNode[T]
	space [20]*bNode[T]
}
type lNode21[T constraints.Ordered] struct {
	lNode[T]
	space [21]*bNode[T]
}
type lNode22[T constraints.Ordered] struct {
	lNode[T]
	space [22]*bNode[T]
}
type lNode23[T constraints.Ordered] struct {
	lNode[T]
	space [23]*bNode[T]
}
type lNode24[T constraints.Ordered] struct {
	lNode[T]
	space [24]*bNode[T]
}
type lNode25[T constraints.Ordered] struct {
	lNode[T]
	space [25]*bNode[T]
}
type lNode26[T constraints.Ordered] struct {
	lNode[T]
	space [26]*bNode[T]
}
type lNode27[T constraints.Ordered] struct {
	lNode[T]
	space [27]*bNode[T]
}
type lNode28[T constraints.Ordered] struct {
	lNode[T]
	space [28]*bNode[T]
}
type lNode29[T constraints.Ordered] struct {
	lNode[T]
	space [29]*bNode[T]
}
type lNode30[T constraints.Ordered] struct {
	lNode[T]
	space [30]*bNode[T]
}
