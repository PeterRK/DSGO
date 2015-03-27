package pairing

import (
	"math/rand"
	"testing"
	"time"
)

func Test_Heap(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
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
	if unit == nil {
		t.Fail()
	}
	if unit.key > node.key {
		node = unit
	}
	heap.PushNode(unit)

	//合并
	heap.Merge(&another)
	var key, fail = heap.Top()
	if fail || key != mark || !another.IsEmpty() {
		t.Fail()
	}

	//部分删除
	for i := 0; i < size; i++ {
		key, fail = heap.Pop()
		if fail || key < mark {
			t.Fail()
		}
		mark = key
	}

	//值上浮
	mark--
	heap.FloatUpX(node, mark)
	key, fail = heap.Top()
	if fail || key != mark || key == node.key {
		t.Fail()
	}

	//节点上浮
	mark--
	heap.FloatUp(node, mark)
	key, fail = heap.Top()
	if fail || key != mark || key != node.key {
		t.Fail()
	}

	//删除
	for i := 0; i < size; i++ {
		key, fail = heap.Pop()
		if fail || key < mark {
			t.Fail()
		}
		mark = key
	}
	if !heap.IsEmpty() {
		t.Fail()
	}
}
