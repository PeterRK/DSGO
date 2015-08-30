package binary

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

	var heap Heap
	const size = 200
	var list, list2 [size]int

	const INT_MAX = int(^uint(0) >> 1)
	var mark, mark2 = INT_MAX, INT_MAX

	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
		if list[i] < mark {
			mark = list[i]
		}
		list2[i] = rand.Int()
		if list2[i] < mark2 {
			mark2 = list2[i]
		}
	}

	//建堆
	heap.Build(list[:])
	var key, err = heap.Top()
	assert(t, err == nil && key == mark)

	//插入
	if mark > mark2 {
		mark = mark2
	}
	for i := 0; i < size; i++ {
		heap.Push(list2[i])
	}
	key, err = heap.Top()
	assert(t, err == nil && key == mark)

	//删除
	for i := 0; i < size*2; i++ {
		key, err = heap.Pop()
		assert(t, err == nil && key >= mark)
	}
}
