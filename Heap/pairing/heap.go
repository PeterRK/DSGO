package pairing

import (
	"errors"
)

//虽然Fibonacci理论复杂度更好，但配对堆实际上更为实用。
type Heap struct {
	root *Node
}
type Node struct {
	key   int
	child *Node
	prev  *Node //父兄节点
	next  *Node //弟节点
}

func (unit *Node) hook(peer *Node) *Node {
	if peer != nil {
		peer.prev = unit
	}
	return peer
}

func (hp *Heap) IsEmpty() bool {
	return hp.root == nil
}
func (hp *Heap) Clear() {
	hp.root = nil
}

func (hp *Heap) Top() (int, error) {
	if hp.IsEmpty() {
		return 0, errors.New("empty")
	}
	return hp.root.key, nil
}

func merge(one *Node, another *Node) *Node { //one != nil && another != nil
	if one.key > another.key {
		one, another = another, one
	}
	another.next = another.hook(one.child)
	one.child, another.prev = another, one
	return one
}
func (hp *Heap) Merge(victim *Heap) {
	if hp != victim && !victim.IsEmpty() {
		if hp.IsEmpty() {
			*hp = *victim
		} else {
			hp.root = merge(hp.root, victim.root)
		}
		victim.Clear()
	}
}

//这货Push时不怎么管整理，到Pop时再做
func (hp *Heap) PushNode(unit *Node) {
	if unit != nil {
		unit.prev, unit.next, unit.child = nil, nil, nil
		if hp.root == nil {
			hp.root = unit
		} else {
			hp.root = merge(hp.root, unit)
		}
	}
}
func (hp *Heap) Push(key int) *Node {
	var unit = new(Node)
	unit.key = key
	hp.PushNode(unit)
	return unit
}
