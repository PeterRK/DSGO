package bst

import (
	"SearchTree/BST/avltree"
	"SearchTree/BST/rbtree"
	"SearchTree/BST/simplebst"
	"fmt"
	"math/rand"
	"time"
	"unsafe"
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

func mixArray(size int) []int {
	var list = make([]int, size)

	const bits_of_int = uint(unsafe.Sizeof(size)) * 8
	var tmp = uint(size)
	var cnt uint = 0
	for cnt < bits_of_int && tmp != 0 {
		cnt++
		tmp >>= 1
	}
	cnt = bits_of_int - cnt - 11
	var mask = ^((^0) << cnt)

	var num = 0
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		if i%32 == 0 { //局部参入有序数列
			num += rand.Int() & mask
			list[i] = num
		} else {
			list[i] = rand.Int()
		}
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
	//for i := 0; i < size; i++ { //对非平衡树而言删除顺序重要
	for i := size - 1; i >= 0; i-- {
		tree.Remove(list[i])
	}
	fmt.Println("remove:", time.Since(start))
}

func DoBenchMark(size int) {
	if size < 1000 {
		fmt.Println("too small")
		return
	}
	var list = mixArray(size)

	DoOneTry(list, SIMPLE_BST)
	DoOneTry(list, AVL_TREE)
	DoOneTry(list, RB_TREE)
}
