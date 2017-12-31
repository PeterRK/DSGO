package bptree

import (
	"math/rand"
	"testing"
	"time"
)

func Benchmark_Insert(b *testing.B) {
	b.StopTimer()
	var tree Tree
	list := make([]int, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		list[i] = rand.Int()
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Insert(list[i])
	}
}

func Benchmark_Remove(b *testing.B) {
	b.StopTimer()
	var tree Tree
	list := make([]int, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		list[i] = rand.Int()
		tree.Insert(list[i])
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Remove(list[i])
	}
}

func Benchmark_Search(b *testing.B) {
	b.StopTimer()
	var tree Tree
	list := make([]int, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		list[i] = rand.Int()
		tree.Insert(list[i])
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Search(list[i])
	}
}
