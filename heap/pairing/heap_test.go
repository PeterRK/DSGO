package pairing

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

func Test_FloatUpAndRemove(t *testing.T) {
	defer utils.FailInPanic(t)

	const size = 200
	list, mark := genRand(size)

	var hp Heap[int]

	//插入
	fake := Node[int]{key: math.MinInt}
	node := &fake
	for i := 0; i < size; i++ {
		unit := new(Node[int])
		unit.key = list[i]
		if unit.key > node.key {
			node = unit
		}
		hp.PushNode(unit)
	}

	super := node.prev
	hp.FloatUp(node, super.key)
	utils.Assert(t, node.prev == super && node.key == super.key)
	hp.Remove(node)
	hp.PushNode(node)

	mark--
	hp.FloatUp(node, mark)
	utils.Assert(t, hp.Top() == mark && mark == node.key)
	hp.Remove(node)
	hp.PushNode(node)

	kid := node.child
	hp.Remove(kid)
	kid = node.child
	hp.FloatUp(kid, mark-1)
}
