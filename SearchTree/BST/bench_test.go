package bst

import (
	"testing"
)

func benchInsert(b *testing.B, hint int) {
	b.StopTimer()
	var tree = newTree(hint)
	var list = ramdomArray(b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(list[i])
	}
}
func benchFind(b *testing.B, hint int) {
	b.StopTimer()
	var tree = newTree(hint)
	var list = ramdomArray(b.N)
	for i := 0; i < b.N; i++ {
		tree.Insert(list[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.Search(list[i])
	}
}
func benchRemove(b *testing.B, hint int) {
	b.StopTimer()
	var tree = newTree(hint)
	var list = ramdomArray(b.N)
	for i := 0; i < b.N; i++ {
		tree.Insert(list[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.Remove(list[i])
	}
}

func Benchmark_SimpleInsert(b *testing.B) {
	benchInsert(b, SIMPLE_BST)
}
func Benchmark_SimpleFind(b *testing.B) {
	benchFind(b, SIMPLE_BST)
}
func Benchmark_SimpleRemove(b *testing.B) {
	benchRemove(b, SIMPLE_BST)
}
func Benchmark_AVLtreeInsert(b *testing.B) {
	benchInsert(b, AVL_TREE)
}
func Benchmark_AVLtreeFind(b *testing.B) {
	benchFind(b, AVL_TREE)
}
func Benchmark_AVLtreeRemove(b *testing.B) {
	benchRemove(b, AVL_TREE)
}
func Benchmark_RBtreeInsert(b *testing.B) {
	benchInsert(b, RB_TREE)
}
func Benchmark_RBtreeFind(b *testing.B) {
	benchFind(b, RB_TREE)
}
func Benchmark_RBtreeRemove(b *testing.B) {
	benchRemove(b, RB_TREE)
}
