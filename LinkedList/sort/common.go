package sort

import (
	"unsafe"
)

type Node struct {
	key  int
	next *Node
}

func FakeHead(this **Node) *Node {
	var base = uintptr(unsafe.Pointer(this))
	var off = unsafe.Offsetof((*this).next)
	return (*Node)(unsafe.Pointer(base - off))
}
