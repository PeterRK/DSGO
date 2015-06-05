package list

import (
	"unsafe"
)

type NodeX struct {
	Prev, Next *NodeX
	Val        int
}
type Ring struct {
	Prev, Next *NodeX
}

func (r *Ring) asNodeX() *NodeX {
	return (*NodeX)(unsafe.Pointer(r))
}
func (r *Ring) Initialize() {
	r.Prev, r.Next = r.asNodeX(), r.asNodeX()
}
func (r *Ring) IsEmpty() bool {
	return r.Next == r.asNodeX()
}

func (r *Ring) InsertHead(node *NodeX) {
	node.Next = r.Next
	node.Next.Prev = node
	node.Prev = r.asNodeX()
	r.Next = node
}
func (r *Ring) PopHead() *NodeX {
	if r.IsEmpty() {
		return nil
	}
	var node = r.Next
	r.Next = node.Next
	r.Next.Prev = r.asNodeX()
	return node
}

func (r *Ring) InsertTail(node *NodeX) {
	node.Prev = r.Prev
	node.Prev.Next = node
	node.Next = r.asNodeX()
	r.Prev = node
}
func (r *Ring) PopTail() *NodeX {
	if r.IsEmpty() {
		return nil
	}
	var node = r.Prev
	r.Prev = node.Prev
	r.Prev.Next = r.asNodeX()
	return node
}

func Release(node *NodeX) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}
