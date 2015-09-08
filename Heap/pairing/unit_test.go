package pairing

import (
	"math/rand"
	"testing"
	"time"
)

func assert(t *testing.T, state bool) {
	if !state {
		t.Fail()
	}
}
func guard_ut(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = -INT_MAX - 1

func Test_Heap(t *testing.T) {
	defer guard_ut(t)

	var heap Heap
	const size = 200
	var list = new([size]int)

	var mark = INT_MAX
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		var key = rand.Int()
		if key < mark {
			mark = key
		}
		list[i] = key
		heap.Push(key)
	}
	var key, err = heap.Top()
	assert(t, err == nil && key == mark)

	for i := 0; i < size; i++ {
		key, err = heap.Pop()
		assert(t, err == nil && key >= mark)
		mark = key
	}
	assert(t, heap.IsEmpty())

	_, err = heap.Top()
	assert(t, err != nil)
	_, err = heap.Pop()
	assert(t, err != nil)
	heap.Push(99)
	assert(t, !heap.IsEmpty())
	heap.Clear()
	assert(t, heap.IsEmpty())
}

func Test_Merge(t *testing.T) {
	defer guard_ut(t)

	var hp1, hp2 Heap
	hp1.Merge(&hp2)
	assert(t, hp1.IsEmpty())
	hp1.Merge(&hp1)

	hp2.Push(999)
	assert(t, !hp2.IsEmpty())
	hp1.Merge(&hp2)
	assert(t, !hp1.IsEmpty())
	assert(t, hp2.IsEmpty())

	hp1.Push(100)
	hp2.Push(101)
	hp1.Merge(&hp2)
	var key, err = hp1.Top()
	assert(t, err == nil && key == 100)

	hp1.Push(11)
	hp2.Push(10)
	hp1.Merge(&hp2)
	key, err = hp1.Top()
	assert(t, err == nil && key == 10)
}

func Test_FloatUpAndRemove(t *testing.T) {
	defer guard_ut(t)

	var heap Heap
	const size = 200
	var list = new([size]int)

	var mark = INT_MAX
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
		if list[i] < mark {
			mark = list[i]
		}
	}

	//插入
	var fake = Node{key: INT_MIN}
	var node = &fake
	for i := 0; i < size; i++ {
		var unit = new(Node)
		unit.key = list[i]
		if unit.key > node.key {
			node = unit
		}
		heap.PushNode(unit)
	}

	heap.Remove(node)
	heap.PushNode(node)

	mark--
	heap.FloatUp(node, mark)
	var key, err = heap.Top()
	assert(t, err == nil && key == mark && key == node.key)

	heap.Remove(node)
	heap.PushNode(node)

	var kid = node.child
	heap.Remove(kid)
	kid = node.child
	heap.FloatUp(kid, mark-1)
}
