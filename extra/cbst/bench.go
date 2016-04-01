package cbst

import (
	"DSGO/Tree/bst/avltree"
	"DSGO/Tree/bst/rbtree"
	"DSGO/extra/cbst/avlt"
	"DSGO/extra/cbst/rbt"
	"fmt"
	"math/rand"
	"time"
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
	PLAIN_AVL
	PLAIN_RB
)

func showName(hint int) string {
	switch hint {
	case AVL_TREE:
		return "AVL tree"
	case RB_TREE:
		return "red-black tree"
	case PLAIN_AVL:
		return "AVL tree [no uplink]"
	case PLAIN_RB:
		return "red-black [no uplink]"
	default:
		panic("Illegal BST type")
	}
}
func newTree(hint int) BST {
	switch hint {
	case AVL_TREE:
		return new(avltree.Tree)
	case RB_TREE:
		return new(rbtree.Tree)
	case PLAIN_AVL:
		return new(avlt.Tree)
	case PLAIN_RB:
		return new(rbt.Tree)
	default:
		panic("Illegal BST type")
	}
}

func ramdomArray(size int) []int32 {
	rand.Seed(time.Now().Unix())
	var list = make([]int32, size)
	for i := 0; i < size; i++ {
		list[i] = int32(rand.Int())
	}
	return list
}
func benchMark(list []int32, hint int) {
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
	//for i := 0; i < size; i++ { //对非平衡树而言删除顺序重要
	for i := size - 1; i >= 0; i-- {
		tree.Remove(list[i])
	}
	fmt.Println("remove:", time.Since(start))
}
func BenchMark(size int) {
	if size < 10000 {
		fmt.Println("too small")
		return
	}
	var list = ramdomArray(size)

	benchMark(list, AVL_TREE)
	benchMark(list, RB_TREE)
	benchMark(list, PLAIN_AVL)
	benchMark(list, PLAIN_RB)
}
