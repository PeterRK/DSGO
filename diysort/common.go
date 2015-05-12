package diysort

import (
	"unsafe"
)

const sz_limit = 7

type Node struct {
	key  int
	next *Node
}

func FakeHead(this **Node) *Node {
	var base = uintptr(unsafe.Pointer(this))
	var off = unsafe.Offsetof((*this).next)
	return (*Node)(unsafe.Pointer(base - off))
}

type pair struct {
	start int
	end   int
}
type stack struct {
	core []pair
}

func (this *stack) clear() {
	this.core = this.core[:0]
}
func (this *stack) size() int {
	return len(this.core)
}
func (this *stack) isEmpty() bool {
	return len(this.core) == 0
}
func (this *stack) push(start int, end int) {
	this.core = append(this.core, pair{start, end})
}
func (this *stack) pop() (start int, end int) {
	var sz = len(this.core) - 1
	var unit = this.core[sz]
	this.core = this.core[:sz]
	return unit.start, unit.end
}
