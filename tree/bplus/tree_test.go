package bplus

import (
	"DSGO/utils"
	"math"
	"math/rand"
	"testing"
	"time"
)

func genRand(size int) []int32 {
	rand.Seed(time.Now().UnixNano())
	rand.Seed(0)
	list := make([]int32, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Int31()
	}
	return list
}

func Test_Tree(t *testing.T) {
	defer utils.FailInPanic(t)

	const size = 5000
	list := genRand(size)

	var tree Tree[int32]
	cnt := 0
	for i := 0; i < size; i++ {
		if tree.Insert(list[i]) {
			cnt++
		}
	}

	for i := 0; i < size; i++ {
		utils.Assert(t, tree.Search(list[i]))
		utils.Assert(t, !tree.Insert(list[i]))
	}

	mark := int32(math.MinInt32)
	tree.Travel(func(val int32) {
		if val < mark {
			panic(val)
		}
		mark = val
	})

	for i := 0; i < size; i++ {
		if tree.Remove(list[i]) {
			cnt--
		}
		utils.Assert(t, !tree.Search(list[i]))
	}
	utils.Assert(t, tree.IsEmpty() && cnt == 0)
	utils.Assert(t, !tree.Remove(0))
}

type elem int32

func genPseudo(size int) []elem {
	rand.Seed(999)
	list := make([]elem, size)
	for i := 0; i < size; i++ {
		list[i] = elem(rand.Uint64())
	}
	return list
}

func Benchmark_Insert(b *testing.B) {
	b.StopTimer()
	var tree Tree[elem]
	list := genPseudo(b.N)
	b.StartTimer()
	for i:= 0; i < len(list); i++ {
		tree.Insert(list[i])
	}
}

func Benchmark_Search(b *testing.B) {
	b.StopTimer()
	var tree Tree[elem]
	list := genPseudo(b.N)
	for i:= 0; i < len(list); i++ {
		tree.Insert(list[i])
	}
	b.StartTimer()
	for i:= 0; i < len(list); i++ {
		tree.Search(list[i])
	}
}

func Benchmark_Remove(b *testing.B) {
	b.StopTimer()
	var tree Tree[elem]
	list := genPseudo(b.N)
	for i:= 0; i < len(list); i++ {
		tree.Insert(list[i])
	}
	b.StartTimer()
	for i:= 0; i < len(list); i++ {
		tree.Remove(list[i])
	}
}

func Benchmark_Mix(b *testing.B) {
	b.StopTimer()
	var tree Tree[elem]
	list := genPseudo(b.N)
	b.StartTimer()
	len1, len2, len3 := len(list)/3, len(list)*2/3, len(list)
	for i:= 0; i < len2; i++ {
		tree.Insert(list[i])
	}
	for i:= len1; i < len2; i++ {
		tree.Remove(list[i])
	}
	for i:= 0; i < len2; i++ {
		tree.Search(list[i])
	}
	for i:= len2; i < len3; i++ {
		tree.Insert(list[i])
	}
	for i:= 0; i < len1; i++ {
		tree.Remove(list[i])
	}
	for i:= len2; i < len3; i++ {
		tree.Search(list[i])
	}
}