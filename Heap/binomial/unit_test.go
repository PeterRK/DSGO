package binomial

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
func guardUT(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

func Test_Heap(t *testing.T) {
	defer guardUT(t)

	var heap Heap
	const size = 200
	var list = new([size]int)

	const INT_MAX = int(^uint(0) >> 1)
	var mark = INT_MAX

	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
		if list[i] < mark {
			mark = list[i]
		}
	}

	for i := 0; i < size; i++ {
		heap.Push(list[i])
	}
	for i := 0; i < size; i++ {
		var key, err = heap.Pop()
		assert(t, err == nil && key >= mark)
		mark = key
	}
	assert(t, heap.IsEmpty())

	var _, err = heap.Top()
	assert(t, err != nil)
	_, err = heap.Pop()
	assert(t, err != nil)
	heap.Push(99)
	assert(t, !heap.IsEmpty())
	heap.Clear()
	assert(t, heap.IsEmpty())
}

func Test_Merge(t *testing.T) {
	defer guardUT(t)

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
