package skiplist

import (
	"math/rand"
	"testing"
	"time"
)

func Benchmark_Insert(b *testing.B) {
	b.StopTimer()
	var dict = NewSkipList()
	var list = make([]int, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		list[i] = rand.Int()
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		dict.Insert(list[i])
	}
}

func Benchmark_Remove(b *testing.B) {
	b.StopTimer()
	var dict = NewSkipList()
	var list = make([]int, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		list[i] = rand.Int()
		dict.Insert(list[i])
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		dict.Remove(list[i])
	}
}

func Benchmark_Search(b *testing.B) {
	b.StopTimer()
	var dict = NewSkipList()
	var list = make([]int, b.N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		list[i] = rand.Int()
		dict.Insert(list[i])
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		dict.Search(list[i])
	}
}
