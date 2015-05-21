package list

import (
	"unsafe"
)

type NodeX struct {
	prev, next *NodeX
	val        int
}
type Ring struct {
	prev, next *NodeX
	self       *NodeX
}

func (ring *Ring) Initialize() {
	ring.self = (*NodeX)(unsafe.Pointer(ring))
	ring.prev, ring.next = ring.self, ring.self
}
func (ring *Ring) IsEmpty() bool {
	return ring.next == ring.self
}

func (ring *Ring) InsertHead(node *NodeX) {
	node.next = ring.next
	node.next.prev = node
	node.prev = ring.self
	ring.next = node
}
func (ring *Ring) PopHead() *NodeX {
	var node = ring.next
	ring.next = node.next
	ring.next.prev = ring.self
	return node
}

func (ring *Ring) InsertTail(node *NodeX) {
	node.prev = ring.prev
	node.prev.next = node
	node.next = ring.self
	ring.prev = node
}
func (ring *Ring) PopTail() *NodeX {
	var node = ring.prev
	ring.prev = node.prev
	ring.prev.next = ring.self
	return node
}

func Release(node *NodeX) {
	node.next.prev = node.prev
	node.prev.next = node.next
}
