package tree

import (
	"DSGO/utils"
	"testing"
)

type elem int32


func Test_NaiveBST(t *testing.T) {
	defer utils.FailInPanic(t)

	const size = 200
	list := utils.RandomArray[elem](size)

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
