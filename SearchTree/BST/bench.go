package bst

import (
	"SearchTree/BST/avltree"
	"SearchTree/BST/rbtree"
	"SearchTree/BST/simplebst"
	"fmt"
	"math/rand"
	"time"
)

type BST interface {
	Insert(key int) bool
	Search(key int) bool
	Remove(key int) bool
}

const (
	SIMPLE_BST = iota
	AVL_TREE
	RB_TREE
)

func showName(hint int) string {
	switch hint {
	case SIMPLE_BST:
		return "simple BST"
	case AVL_TREE:
		return "AVL tree"
	case RB_TREE:
		return "red-black tree"
	default:
		panic("Illegal BST type")
	}
}
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

func ramdomArray(size int) []int {
	rand.Seed(time.Now().Unix())
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
	}
	return list
}

func DoOneTry(list []int, hint int) {
	var tree = newTree(hint)
	var size = len(list)
	fmt.Println(showName(hint))

	var start = time.Now()
	for i := 0; i < size; i++ {
		tree.Insert(list[i])
	}
	fmt.Println("insert:", time.Since(start))
	start = time.Now()
	for i := 0; i < size; i++ {
		tree.Search(list[i])
	}
	fmt.Println("search:", time.Since(start))
	start = time.Now()
	for i := 0; i < size; i++ {
		tree.Remove(list[i])
	}
	fmt.Println("remove:", time.Since(start))

	//	fmt.Println()
}

func DoBenchMark() {
	const size = 2000000
	var list = ramdomArray(size)

	DoOneTry(list, SIMPLE_BST)
	DoOneTry(list, AVL_TREE)
	DoOneTry(list, RB_TREE)
}
