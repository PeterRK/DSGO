package bptree

type stackNode struct {
	pt  *index
	idx int
}
type stack struct {
	core []stackNode
}

func (this *stack) clear() {
	this.core = this.core[:0]
}
func (this *stack) isEmpty() bool {
	return len(this.core) == 0
}
func (this *stack) push(pt *index, idx int) {
	this.core = append(this.core, stackNode{pt, idx})
}
func (this *stack) pop() (pt *index, idx int) {
	var sz = len(this.core) - 1
	var unit = this.core[sz]
	this.core = this.core[:sz]
	return unit.pt, unit.idx
}
