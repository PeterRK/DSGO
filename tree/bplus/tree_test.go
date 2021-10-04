package bplus

import (
	"DSGO/utils"
	"math"
	"math/rand"
	"testing"
	"time"
)

func genRand(size int) []int32 {
	rand.Seed(time.Now().Unix())
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
