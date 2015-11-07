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
	var lst2 = new([size]int)

	const INT_MAX = int(^uint(0) >> 1)
	var mark, mark2 = INT_MAX, INT_MAX

	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
		if list[i] < mark {
			mark = list[i]
		}
		lst2[i] = rand.Int()
		if lst2[i] < mark2 {
			mark2 = lst2[i]
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
		heap.Push(lst2[i])
	}
	key, err = heap.Top()
	assert(t, err == nil && key == mark)

	//删除
	for i := 0; i < size*2; i++ {
		key, err = heap.Pop()
		assert(t, err == nil && key >= mark)
	}

	key, err = heap.Top()
	assert(t, err != nil)
	key, err = heap.Pop()
	assert(t, err != nil)
	heap.Push(99)
	assert(t, !heap.IsEmpty())
	heap.Clear()
	assert(t, heap.IsEmpty())
}
