package binary

import (
	dshp "DSGO/heap"
)

type NodeG[T any] struct {
	Val T
	pos int
}

type heap[T any] struct {
	less func(a, b *T) bool
	vec  []*NodeG[T]
}

func New[T any](less func(a, b *T) bool) dshp.NodeHeap[NodeG[T]] {
	hp := new(heap[T])
	hp.less = less
	return hp
}

func (hp *heap[T]) Size() int {
	return len(hp.vec)
}

func (hp *heap[T]) IsEmpty() bool {
	return len(hp.vec) == 0
}

func (hp *heap[T]) Clear() {
	hp.vec = hp.vec[:0]
}

func (hp *heap[T]) Push(node *NodeG[T]) {
	if node == nil {
		return
	}
	place := len(hp.vec)
	hp.vec = append(hp.vec, node)
	hp.up(place)
}

func (hp *heap[T]) Top() *NodeG[T] {
	if hp.IsEmpty() {
		return nil
	}
	return hp.vec[0]
}

func (hp *heap[T]) Pop() *NodeG[T] {
	size := hp.Size()
	if size == 0 {
		return nil
	}
	node := hp.vec[0]
	if size == 1 {
		hp.vec = hp.vec[:0]
	} else {
		hp.vec[0] = hp.vec[size-1]
		hp.vec = hp.vec[:size-1]
		hp.down(0)
	}
	return node
}

func (hp *heap[T]) down(pos int) {
	tgt := hp.vec[pos]
	kid, last := pos*2+1, len(hp.vec)-1
	for kid < last {
		if hp.less(&hp.vec[kid+1].Val, &hp.vec[kid].Val) {
			kid++
		}
		if !hp.less(&hp.vec[kid].Val, &tgt.Val) {
			break
		}
		hp.vec[pos] = hp.vec[kid]
		pos, kid = kid, kid*2+1
	}
	if kid == last && hp.less(&hp.vec[kid].Val, &tgt.Val) {
		hp.vec[pos], pos = hp.vec[kid], kid
	}
	hp.vec[pos] = tgt
}

func (hp *heap[T]) up(pos int) {
	tgt := hp.vec[pos]
	for pos > 0 {
		parent := (pos - 1) / 2
		if !hp.less(&tgt.Val, &hp.vec[parent].Val) {
			break
		}
		hp.vec[pos], pos = hp.vec[parent], parent
	}
	hp.vec[pos] = tgt
}

func (hp *heap[T]) FloatUp(node *NodeG[T]) {
	if node == nil || hp.vec[node.pos] != node {
		return
	}
	hp.floatUp(node.pos)
}

func (hp *heap[T]) floatUp(pos int) {
	node := hp.vec[pos]
	for pos > 0 {
		parent := (pos - 1) / 2
		if !hp.less(&node.Val, &hp.vec[parent].Val) {
			break
		}
		hp.vec[pos] = hp.vec[parent]
		hp.vec[pos].pos, pos = pos, parent
	}
	hp.vec[pos], node.pos = node, pos
}
