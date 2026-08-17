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
	utils.Assert(t, hp.Top() == min(m1, m2))

	//删除
	mark := min(m1, m2)
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

func Test_NodeHeapFloatUp(t *testing.T) {
	defer utils.FailInPanic(t)

	less := func(a, b *int) bool { return *a < *b }
	hp := New(less)
	nodes := []NodeG[int]{{Val: 5}, {Val: 10}, {Val: 8}}
	for i := range nodes {
		hp.Push(&nodes[i])
	}

	utils.Assert(t, hp.Pop() == &nodes[0])
	nodes[1].Val = 1
	hp.FloatUp(&nodes[1])
	utils.Assert(t, hp.Pop() == &nodes[1])
	utils.Assert(t, hp.Pop() == &nodes[2])
	utils.Assert(t, nodes[0].pos == -1 && nodes[1].pos == -1 && nodes[2].pos == -1)

	hp.Push(&nodes[0])
	hp.Push(&nodes[1])
	hp.Clear()
	utils.Assert(t, nodes[0].pos == -1 && nodes[1].pos == -1)
}
