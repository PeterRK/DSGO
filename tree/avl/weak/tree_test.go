package weak

import (
	"DSGO/utils"
	"math/rand"
	"testing"
	"time"
)

type elem int32

func genRand(size int) []elem {
	rand.Seed(time.Now().UnixNano())
	list := make([]elem, size)
	for i := 0; i < size; i++ {
		list[i] = elem(rand.Uint64())
	}
	return list
}

func (root *node[T]) check(t *testing.T) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		utils.Assert(t, root.lDiff == 1 &&  root.rDiff == 1)
	} else {
		utils.Assert(t, root.lDiff == 1 || root.lDiff == 2)
		utils.Assert(t, root.rDiff == 1 || root.rDiff == 2)
		root.left.check(t)
		root.right.check(t)
	}
}

func Test_Tree(t *testing.T) {
	defer utils.FailInPanic(t)

	const size = 2000
	list := genRand(size)

	var tree Tree[elem]
	cnt := 0
	for i := 0; i < size; i++ {
		if tree.Insert(list[i]) {
			cnt++
		}
		tree.root.check(t)
	}
	utils.Assert(t, tree.Size() == cnt)

	for i := 0; i < size; i++ {
		utils.Assert(t, tree.Search(list[i]))
		utils.Assert(t, !tree.Insert(list[i]))
	}
	for i := 0; i < size; i++ {
		if tree.Remove(list[i]) {
			cnt--
		}
		tree.root.check(t)
		utils.Assert(t, !tree.Search(list[i]))
	}
	utils.Assert(t, tree.IsEmpty() && tree.Size() == 0 && cnt == 0)
	utils.Assert(t, !tree.Remove(0))
}


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