package binomial

import (
	"DSGO/utils"
	"math"
	"math/rand"
	"testing"
	"time"
)

func genRand(size int) ([]int, int) {
	rand.Seed(time.Now().UnixNano())
	list := make([]int, size)
	min := math.MaxInt
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
		if list[i] < min {
			min = list[i]
		}
	}
	return list, min
}

func Test_Heap(t *testing.T) {
	defer utils.FailInPanic(t)

	const size = 200
	list, mark := genRand(size)

	var hp Heap[int]

	for i := 0; i < size; i++ {
		hp.Push(list[i])
	}
	for i := 0; i < size; i++ {
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

func Test_Merge(t *testing.T) {
	defer utils.FailInPanic(t)

	var hp1, hp2 Heap[int]
	hp1.Merge(&hp2)
	utils.Assert(t, hp1.IsEmpty())
	hp1.Merge(&hp1)

	hp2.Push(999)
	utils.Assert(t, hp2.Size() == 1)
	hp1.Merge(&hp2)
	utils.Assert(t, hp1.Size() == 1)
	utils.Assert(t, hp2.IsEmpty())

	hp1.Push(100)
	hp2.Push(101)
	hp1.Merge(&hp2)
	utils.Assert(t, hp1.Size() == 3)
	utils.Assert(t, hp1.Top() == 100)

	hp2.Push(11)
	hp2.Push(10)
	hp1.Merge(&hp2)
	utils.Assert(t, hp1.Size() == 5)
	utils.Assert(t, hp1.Top() == 10)

	hp1.Pop()
	utils.Assert(t, hp1.Size() == 4)
	utils.Assert(t, hp2.Size() == 0)
}
