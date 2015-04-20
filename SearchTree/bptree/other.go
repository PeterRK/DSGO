package bptree

import (
	"fmt"
	"unsafe"
)

const LEAF_HALF = 8
const LEAF_FULL = LEAF_HALF*2 - 1
const LEAF_QUARTER = LEAF_HALF / 2

type leaf struct {
	cnt  int //用正数表示leaf节点
	data [LEAF_FULL]int
	next *leaf
}

func (node *leaf) locate(key int) int {
	var end = node.cnt - 1
	var start = 0
	for start < end {
		var mid = (start + end) / 2
		if key > node.data[mid] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

const INDEX_HALF = 4
const INDEX_FULL = INDEX_HALF*2 - 1
const INDEX_QUARTER = INDEX_HALF / 2

type index struct {
	cnt  int //用负数表示index节点
	data [INDEX_FULL]int
	kids [INDEX_FULL]*index
}

func (node *index) locate(key int) int {
	var end = (-node.cnt) - 1
	var start = 0
	for start < end {
		var mid = (start + end) / 2
		if key > node.data[mid] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

type stackNode struct {
	pt  *index
	idx int
}
type stack struct {
	core []stackNode
}
type Tree struct {
	root *index
	head *leaf
	path stack
}

func (unit *index) ceil() int {
	if unit.cnt < 0 {
		return unit.data[-unit.cnt-1]
	} else {
		var node = (*leaf)(unsafe.Pointer(unit))
		return node.data[node.cnt-1]
	}
}

func (tree *Tree) Find(key int) bool {
	if tree.root == nil ||
		key > tree.root.ceil() {
		return false
	}

	var target = tree.root
	for target.cnt < 0 { //index节点
		var idx = target.locate(key)
		if key == target.data[idx] {
			return true
		}
		target = target.kids[idx]
	}
	var node = (*leaf)(unsafe.Pointer(target)) //叶节点
	var place = node.locate(key)
	return key == node.data[place]
}

func (tree *Tree) Print() {
	for node := tree.head; node != nil; node = node.next {
		for i := 0; i < node.cnt; i++ {
			fmt.Printf("%d ", node.data[i])
		}
	}
	fmt.Println()
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
