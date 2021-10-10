package rank

import (
	"constraints"
)

//为了方便查询排行，加入了数量记录

type node[T constraints.Ordered] struct {
	state  int8 //(2), 1, 0, -1, (-2)
	weight int32
	key    T
	parent *node[T]
	left   *node[T]
	right  *node[T]
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

//找到返回序号（从1开始），没有返回0
func (tr *Tree[T]) Search(key T) int {
	target, base := tr.root, int32(0)
	for target != nil {
		if key == target.key {
			return int(base + target.rank())
		}
		if key < target.key {
			target = target.left
		} else {
			base += target.rank()
			target = target.right
		}
	}
	return 0
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

func (unit *node[T]) Weight() int32 {
	if unit == nil {
		return 0
	}
	return unit.weight
}
func (unit *node[T]) rank() int32 {
	return unit.left.Weight() + 1
}

func newNode[T constraints.Ordered](parent *node[T], key T) (unit *node[T]) {
	unit = new(node[T])
	//unit.state = 0
	unit.weight = 1
	//unit.left, unit.right = nil, nil
	unit.parent, unit.key = parent, key
	return unit
}
