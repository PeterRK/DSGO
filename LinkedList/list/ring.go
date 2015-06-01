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

func (r *Ring) Initialize() {
	r.self = (*NodeX)(unsafe.Pointer(r))
	r.prev, r.next = r.self, r.self
}
func (r *Ring) IsEmpty() bool {
	return r.next == r.self
}

func (r *Ring) InsertHead(node *NodeX) {
	node.next = r.next
	node.next.prev = node
	node.prev = r.self
	r.next = node
}
func (r *Ring) PopHead() *NodeX {
	if r.IsEmpty() {
		return nil
	}
	var node = r.next
	r.next = node.next
	r.next.prev = r.self
	return node
}

func (r *Ring) InsertTail(node *NodeX) {
	node.prev = r.prev
	node.prev.next = node
	node.next = r.self
	r.prev = node
}
func (r *Ring) PopTail() *NodeX {
	if r.IsEmpty() {
		return nil
	}
	var node = r.prev
	r.prev = node.prev
	r.prev.next = r.self
	return node
}

func Release(node *NodeX) {
	node.next.prev = node.prev
	node.prev.next = node.next
}
