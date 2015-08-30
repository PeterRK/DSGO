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

func Test_Heap(t *testing.T) {
	defer guard_ut(t)

	var heap, another Heap
	const size = 200
	var list [size * 2]int

	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = -INT_MAX - 1
	var mark = INT_MAX

	rand.Seed(time.Now().Unix())
	for i := 0; i < size*2; i++ {
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
	for i := 0; i < size; i++ {
		another.Push(list[size+i])
	}

	//部分删除
	var unit = another.PopNode()
	assert(t, unit != nil)
	if unit.key > node.key {
		node = unit
	}
	heap.PushNode(unit)

	//合并
	heap.Merge(&another)
	var key, err = heap.Top()
	assert(t, err == nil && key == mark && another.IsEmpty())

	//部分删除
	for i := 0; i < size; i++ {
		key, err = heap.Pop()
		assert(t, err == nil && key >= mark)
		mark = key
	}

	//节点上浮
	mark--
	heap.FloatUp(node, mark)
	key, err = heap.Top()
	assert(t, err == nil && key == mark && key == node.key)

	//删除
	for i := 0; i < size; i++ {
		key, err = heap.Pop()
		assert(t, err == nil && key >= mark)
		mark = key
	}
	assert(t, heap.IsEmpty())
}
