package bplus

import (
	"DSGO/array"
	"constraints"
	"fmt"
)

type trait[T constraints.Ordered] struct {
	node *iNode[T]
	idx  int
}

type Tree[T constraints.Ordered] struct {
	root *iNode[T]
	head *lNode[T]
	size int
	path array.Stack[trait[T]]
}

func (tr *Tree[T]) Clear() {
	*tr = Tree[T]{}
}

func (tr *Tree[T]) IsEmpty() bool {
	return tr.root == nil
}

func (tr *Tree[T]) Size() int {
	return tr.size
}

func (tr *Tree[T]) Travel(doit func(T)) {
	for node := tr.head; node != nil; node = node.next {
		for _, key := range node.view {
			doit(key)
		}
	}
}

func (tr *Tree[T]) Search(key T) bool {
	if tr.root == nil ||
		key > tr.root.ceil() {
		return false
	}
	node := tr.root
	for node.inner {
		pos := array.SearchFirstGE(node.view, key)
		if key == node.view[pos] {
			return true
		}
		node = node.kids[pos]
	}
	return key == node.view[array.SearchFirstGE(node.view, key)]
}

func (tr *Tree[T]) Insert(key T) bool {
	if tr.insert(key) {
		tr.size++
		return true
	}
	return false
}

//成功返回true，冲突返回false。
func (tr *Tree[T]) insert(key T) bool {
	if tr.root == nil {
		node := newLeaf[T]()
		node.data[0] = key
		node.view = node.data[:1]
		tr.head, tr.root = node, node.asIndex()
		return true
	}

	tr.path.Clear()
	node, pos := tr.root, 0
	if key > tr.root.ceil() { //右界拓展
		for node.inner {
			pos = len(node.view) - 1
			node.data[pos] = key //之后难以修改，现在先改掉
			tr.path.Push(trait[T]{node, pos})
			node = node.kids[pos]
		}
		pos = len(node.view)
	} else {
		for node.inner {
			pos = array.SearchFirstGE(node.view, key)
			if key == node.view[pos] {
				return false
			}
			tr.path.Push(trait[T]{node, pos})
			node = node.kids[pos]
		}
		pos = array.SearchFirstGE(node.view, key)
		if key == node.view[pos] {
			return false
		}
	}

	peer := node.asLeaf().insert(pos, key).asIndex()
	for peer != nil {
		if tr.path.IsEmpty() {
			unit := newIndex[T]()
			unit.view = unit.data[:2]
			unit.data[0], unit.data[1] = node.ceil(), peer.ceil()
			unit.kids[0], unit.kids[1] = node, peer
			tr.root, peer = unit, nil
		} else {
			last := tr.path.Pop()
			last.node.view[last.idx] = node.ceil()
			node, peer = last.node, last.node.insert(last.idx+1, peer)
		}
	}
	return true
}

func (tr *Tree[T]) Remove(key T) bool {
	if tr.remove(key) {
		tr.size--
		return true
	}
	return false
}

//B+树的删除比较复杂，本实现采用积极合并策略。
//成功返回true，没有返回false。
func (tr *Tree[T]) remove(key T) bool {
	if tr.root == nil ||
		key > tr.root.ceil() {
		return false
	}

	tr.path.Clear()
	node := tr.root
	for node.inner {
		pos := array.SearchFirstGE(node.view, key)
		tr.path.Push(trait[T]{node, pos})
		node = node.kids[pos]
	}
	pos := array.SearchFirstGE(node.view, key)
	if key != node.view[pos] {
		return false
	}

	node.asLeaf().remove(pos)
	if tr.path.IsEmpty() {
		if len(node.view) == 0 {
			tr.root, tr.head = nil, nil
		}
		return true
	} //除了到根节点，len(node.view) >= 2
	shrink, ceil := (pos == len(node.view)), node.ceil()

	last := tr.path.Pop()
	for limit, isLeaf := lQuarterSize, true; len(node.view) < limit; limit, isLeaf = iQuarterSize, false {
		peer := node
		if last.idx == len(last.node.view)-1 {
			node = last.node.kids[last.idx-1]
		} else {
			last.idx++
			peer, shrink = last.node.kids[last.idx], false
		}
		combined := false
		if isLeaf {
			combined = node.asLeaf().combine(peer.asLeaf())
		} else {
			combined = node.combine(peer)
		}
		last.node.view[last.idx-1] = node.ceil()
		if !combined {
			break
		}
		last.node.remove(last.idx)
		if tr.path.IsEmpty() {
			if len(last.node.view) == 1 {
				tr.root = last.node.kids[0]
			}
			return true
		}
		node = last.node
		last = tr.path.Pop()
	}
	if shrink {
		last.node.view[last.idx] = ceil
		for last.idx == len(last.node.view)-1 && //级联
			!tr.path.IsEmpty() {
			last = tr.path.Pop()
			last.node.view[last.idx] = ceil
		}
	}
	return true
}

func (node *iNode[T]) debug(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
	fmt.Println(node.view)
	if node.inner {
		for i := 0; i < len(node.view); i++ {
			node.kids[i].debug(indent + 1)
		}
	}
}

func (tr *Tree[T]) debug() {
	if tr.root != nil {
		tr.root.debug(0)
	}
	fmt.Println("================")
}
