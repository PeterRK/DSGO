package rank

import (
	"DSGO/array"
	"DSGO/utils"
	"testing"
)

type elem int32

func Test_Tree(t *testing.T) {
	defer utils.FailInPanic(t)

	const size = 2000
	list := utils.RandomArray[elem](size)

	var tree Tree[elem]
	cnt := 0
	for i := 0; i < size; i++ {
		if tree.Insert(list[i]) > 0 {
			cnt++
		}
	}
	utils.Assert(t, tree.Size() == cnt)

	for i := 0; i < size; i++ {
		utils.Assert(t, tree.Search(list[i]) != 0)
		utils.Assert(t, tree.Insert(list[i]) < 0)
	}

	for i := 0; i < size; i++ {
		if tree.Remove(list[i]) > 0 {
			cnt--
		}
		utils.Assert(t, tree.Search(list[i]) == 0)
	}

	utils.Assert(t, tree.IsEmpty() && tree.Size() == 0 && cnt == 0)
	utils.Assert(t, tree.Remove(0) == 0)
}

func Test_Rank(t *testing.T) {
	defer utils.FailInPanic(t)

	const size = 200
	list := make([]elem, size)

	for i := 0; i < size; i++ {
		list[i] = elem(i + 1)
	}
	array.RandomShuffle(list)
	shadow := make([]elem, size)
	copy(shadow, list)

	var tree Tree[elem]

	utils.Assert(t, tree.Insert(shadow[0]) == 1)
	for i := 1; i < size; i++ {
		key := shadow[i]
		rank := tree.Insert(key)
		pos := array.SearchSuccessor(shadow[:i], key)
		array.InsertTo(shadow[:i], pos, key)
		utils.Assert(t, rank == pos+1)
		utils.Assert(t, tree.Insert(key) == -(pos+1))
	}

	for i := 0; i < size; i++ {
		utils.Assert(t, tree.Search(elem(i+1)) == i+1)
	}

	for i := 0; i < size; i++ {
		key := list[i]
		rank := tree.Remove(list[i])
		pos := array.Search(shadow[:size-i], key)
		array.EraseFrom(shadow[:size-i], pos, true)
		utils.Assert(t, rank == pos+1)
	}
}