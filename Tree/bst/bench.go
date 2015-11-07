package bst

import (
	"Tree/bst/avltree"
	"Tree/bst/rbtree"
	"Tree/bst/simplebst"
	"math/rand"
	"time"
	"unsafe"
)

type BST interface {
	IsEmpty() bool
	Insert(key int32) bool
	Search(key int32) bool
	Remove(key int32) bool
}

const (
	SIMPLE_BST = iota
	AVL_TREE
	RB_TREE
)

func newTree(hint int) BST {
	switch hint {
	case SIMPLE_BST:
		return new(simplebst.Tree)
	case AVL_TREE:
		return new(avltree.Tree)
	case RB_TREE:
		return new(rbtree.Tree)
	default:
		panic("Illegal BST type")
	}
}

func mixedArray(size int) []int32 {
	var list = make([]int32, size)

	const bits_of_int = uint(unsafe.Sizeof(list[0])) * 8
	var tmp = uint(size)
	var cnt = uint(0)
	for cnt < bits_of_int && tmp != 0 {
		cnt++
		tmp >>= 1
	}
	cnt = bits_of_int - cnt - 11
	var mask = ^((^0) << cnt)

	var num = int32(0)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		if i%32 == 0 { //局部摻入有序数列
			num += int32(rand.Int() & mask)
			list[i] = num
		} else {
			list[i] = int32(rand.Int())
		}
	}
	return list
}
