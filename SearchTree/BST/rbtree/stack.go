package rbtree

type stackNode struct {
	pt *node
	lf bool
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
func (this *stack) push(pt *node, lf bool) {
	this.core = append(this.core, stackNode{pt, lf})
}
func (this *stack) pop() (pt *node, lf bool) {
	var sz = len(this.core) - 1
	var unit = this.core[sz]
	this.core = this.core[:sz]
	return unit.pt, unit.lf
}
func (this *stack) top() (pt *node, lf bool) {
	var unit = this.core[len(this.core)-1]
	return unit.pt, unit.lf
}
