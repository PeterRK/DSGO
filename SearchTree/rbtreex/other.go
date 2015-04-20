//红黑树的实现
package rbtreex

type node struct {
	key   int
	black bool
	left  *node
	right *node
}

var nil_node = node{black: true, left: nil, right: nil}
var NULL = &nil_node

type stackNode struct {
	pt *node
	lf bool
}
type stack struct {
	core []stackNode
}
type Tree struct {
	root *node
	path stack
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

func (tree *Tree) Find(key int) bool {
	var target = tree.root
	for target != NULL {
		if key == target.key {
			return true
		}
		if key < target.key {
			target = target.left
		} else {
			target = target.right
		}
	}
	return false
}

func (tree *Tree) hookSubTree(subtree *node) {
	if tree.path.isEmpty() {
		tree.root = subtree
	} else {
		if super, lf := tree.path.top(); lf {
			super.left = subtree
		} else {
			super.right = subtree
		}
	}
}
