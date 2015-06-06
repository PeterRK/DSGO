package list

import (
	"unsafe"
)

type NodeX struct {
	Prev, Next *NodeX
	Val        int
}
type Ring struct {
	tail, head *NodeX
}

func (r *Ring) asNodeX() *NodeX {
	return (*NodeX)(unsafe.Pointer(r))
}
func (r *Ring) Initialize() {
	r.head, r.tail = r.asNodeX(), r.asNodeX()
}
func (r *Ring) IsEmpty() bool {
	return r.head == r.asNodeX()
}

func (r *Ring) Head() *NodeX {
	if r.IsEmpty() {
		return nil
	}
	return r.head
}
func (r *Ring) PopHead() *NodeX {
	if r.IsEmpty() {
		return nil
	}
	var node = r.head
	r.head = node.Next
	r.head.Prev = r.asNodeX()
	return node
}
func (r *Ring) InsertHead(node *NodeX) {
	node.Next = r.head
	node.Next.Prev = node
	node.Prev = r.asNodeX()
	r.head = node
}

func (r *Ring) Tail() *NodeX {
	if r.IsEmpty() {
		return nil
	}
	return r.tail
}
func (r *Ring) PopTail() *NodeX {
	if r.IsEmpty() {
		return nil
	}
	var node = r.tail
	r.tail = node.Prev
	r.tail.Next = r.asNodeX()
	return node
}
func (r *Ring) InsertTail(node *NodeX) {
	node.Prev = r.tail
	node.Prev.Next = node
	node.Next = r.asNodeX()
	r.tail = node
}

func Release(node *NodeX) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}
