package linkedlist

import (
	"unsafe"
)

type BiNode struct {
	prev, next *BiNode
	val        int
}

type Ring struct {
	prev, next *BiNode
	self       *BiNode
}

func (ring *Ring) Initialize() {
	ring.self = (*BiNode)(unsafe.Pointer(ring))
	ring.prev, ring.next = ring.self, ring.self
}
func (ring *Ring) IsEmpty() bool {
	return ring.next == ring.self
}

func (ring *Ring) InsertHead(node *BiNode) {
	node.next = ring.next
	node.next.prev = node
	node.prev = ring.self
	ring.next = node
}
func (ring *Ring) PopHead() *BiNode {
	var node = ring.next
	ring.next = node.next
	ring.next.prev = ring.self
	return node
}

func (ring *Ring) InsertTail(node *BiNode) {
	node.prev = ring.prev
	node.prev.next = node
	node.next = ring.self
	ring.prev = node
}
func (ring *Ring) PopTail() *BiNode {
	var node = ring.prev
	ring.prev = node.prev
	ring.prev.next = ring.self
	return node
}

func Release(node *BiNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}
