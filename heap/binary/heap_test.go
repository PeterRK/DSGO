package binary

import (
	"DSGO/utils"
	"math"
	"testing"
)

func genRand(size int) ([]int, int) {
	list := utils.RandomArray[int](size)
	min := math.MaxInt
	for i := 0; i < size; i++ {
		if list[i] < min {
			min = list[i]
		}
	}
	return list, min
}

func Test_Heap(t *testing.T) {
	defer utils.FailInPanic(t)

	const size = 200
	lst1, m1 := genRand(size)
	lst2, m2 := genRand(size)
	var hp Heap[int]

	//建堆
	hp.BuildInPlace(lst1)
	utils.Assert(t, hp.Top() == m1)

	//插入
	for i := 0; i < size; i++ {
		hp.Push(lst2[i])
	}
	utils.Assert(t, hp.Top() == utils.Min(m1, m2))

	//删除
	mark := utils.Min(m1, m2)
	for i := 0; i < size*2; i++ {
		key := hp.Pop()
		utils.Assert(t, key >= mark)
		mark = key
	}

	utils.Assert(t, hp.IsEmpty())
	hp.Push(99)
	utils.Assert(t, !hp.IsEmpty())
	hp.Clear()
	utils.Assert(t, hp.IsEmpty())
}
