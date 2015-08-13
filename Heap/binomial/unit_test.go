package binomial

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
	var mark = INT_MAX

	rand.Seed(time.Now().Unix())
	for i := 0; i < size*2; i++ {
		list[i] = rand.Int()
		if list[i] < mark {
			mark = list[i]
		}
	}

	//插入
	for i := 0; i < size; i++ {
		heap.Push(list[i])
		another.Push(list[size+i])
	}

	//合并
	heap.Merge(&another)
	var key, err = heap.Top()
	if err != nil || key != mark || !another.IsEmpty() {
		t.Fail()
	}

	//删除
	for i := 0; i < size*2; i++ {
		key, err = heap.Pop()
		if err != nil || key < mark {
			t.Fail()
		}
		mark = key
	}
	if !heap.IsEmpty() {
		t.Fail()
	}
}
