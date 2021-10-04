package tree

import (
	"DSGO/utils"
	"math/rand"
	"testing"
	"time"
)

type elem int32

func genRand(size int) []elem {
	rand.Seed(time.Now().Unix())
	list := make([]elem, size)
	for i := 0; i < size; i++ {
		list[i] = elem(rand.Uint64())
	}
	return list
}

func Test_NaiveBST(t *testing.T) {
	defer utils.FailInPanic(t)

	const size = 200
	list := genRand(size)

	var tree NaiveBST[elem]
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
	for i := 0; i < size; i++ {
		if tree.Remove(list[i]) {
			cnt--
		}
		utils.Assert(t, !tree.Search(list[i]))
	}
	utils.Assert(t, tree.IsEmpty() && cnt == 0)
	utils.Assert(t, !tree.Remove(0))
}
