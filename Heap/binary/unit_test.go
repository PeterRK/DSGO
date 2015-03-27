package binary

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
	var key, fail = heap.Top()
	if fail || key != mark {
		t.Fail()
	}

	//插入
	if mark > mark2 {
		mark = mark2
	}
	for i := 0; i < size; i++ {
		heap.Push(list2[i])
	}
	key, fail = heap.Top()
	if fail || key != mark {
		t.Fail()
	}

	//删除
	for i := 0; i < size*2; i++ {
		key, fail = heap.Pop()
		if fail || key < mark {
			t.Fail()
		}
	}
}
