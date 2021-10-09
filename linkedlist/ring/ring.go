package ring

import (
	"unsafe"
)

type node struct {
	prev, next *node
}

func (n *node) Release() {
	if n.next != nil {
		n.next.prev = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
	n.prev = nil
	n.next = nil
}

type ring struct {
	node
}

func (r *ring) init() {
	r.prev = &r.node
	r.next = &r.node
}

func (r *ring) IsEmpty() bool {
	return r.next == &r.node
}

func (r *ring) head() *node {
	if r.IsEmpty() {
		return nil
	}
	return r.next
}

func (r *ring) popHead() *node {
	unit := r.head()
	if unit != nil {
		unit.Release()
	}
	return unit
}

func (r *ring) unsafePopHead() *node {
	unit := r.next
	unit.Release()
	return unit
}

func (r *ring) pushHead(unit *node) {
	unit.next = r.next
	unit.next.prev = unit
	unit.prev = &r.node
	r.next = unit
}

func (r *ring) tail() *node {
	if r.IsEmpty() {
		return nil
	}
	return r.prev
}

func (r *ring) popTail() *node {
	unit := r.tail()
	if unit != nil {
		unit.Release()
	}
	return unit
}

func (r *ring) unsafePopTail() *node {
	unit := r.prev
	unit.Release()
	return unit
}

func (r *ring) pushTail(unit *node) {
	unit.prev = r.prev
	unit.prev.next = unit
	unit.next = &r.node
	r.prev = unit
}

type Node[T any] struct {
	node
	Val T
}

func cast[T any](unit *node) *Node[T] {
	return (*Node[T])(unsafe.Pointer(unit))
}

type Ring[T any] interface {
	IsEmpty() bool
	Head() *Node[T]
	PopHead() *Node[T]
	PushHead(*Node[T])
	Tail() *Node[T]
	PopTail() *Node[T]
	PushTail(*Node[T])
}

type rinG[T any] struct {
	ring
}

func New[T any]() Ring[T] {
	r := new(rinG[T])
	r.init()
	return r
}

func (r *rinG[T]) Head() *Node[T] {
	return cast[T](r.head())
}

func (r *rinG[T]) PopHead() *Node[T] {
	return cast[T](r.popHead())
}

func (r *rinG[T]) PushHead(unit *Node[T]) {
	r.pushHead(&unit.node)
}

func (r *rinG[T]) Tail() *Node[T] {
	return cast[T](r.tail())
}

func (r *rinG[T]) PopTail() *Node[T] {
	return cast[T](r.popTail())
}

func (r *rinG[T]) PushTail(unit *Node[T]) {
	r.pushTail(&unit.node)
}
