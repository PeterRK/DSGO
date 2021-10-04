package sort

import (
	"DSGO/utils"
	"math/rand"
	"testing"
	"time"
)

type Unit struct {
	val int
}

func UnitLess(a, b *Unit) bool {
	return a.val < b.val
}

func Test_Ranker(t *testing.T) {
	defer utils.FailInPanic(t)

	const size = 2000
	rand.Seed(time.Now().Unix())
	lst1 := make([]int, size)
	lst2 := make([]Unit, size)
	for i := 0; i < size; i++ {
		num := rand.Int()
		lst1[i] = num
		lst2[i].val = num
	}

	MergeSort(lst1)
	utils.Assert(t, IsSorted(lst1))

	ranker := Ranker[Unit]{Less: UnitLess}
	ranker.Sort(lst2)
	for i := 0; i < size; i++ {
		utils.Assert(t, lst1[i] == lst2[i].val)
	}
}
