package array

import (
	"testing"
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

func Test_BinarySearch(t *testing.T) {
	defer guard_ut(t)
	var list = []int{2, 2, 4, 6, 6, 6, 8, 8}

	assert(t, SearchFirst(list, 1) == 0)
	assert(t, SearchFirst(list, 2) == 0)
	assert(t, SearchFirst(list, 3) == 2)
	assert(t, SearchFirst(list, 6) == 3)
	assert(t, SearchFirst(list, 9) == 8)
	assert(t, SearchLast(list, 1) == -1)
	assert(t, SearchLast(list, 5) == 2)
	assert(t, SearchLast(list, 6) == 5)
	assert(t, SearchLast(list, 8) == 7)
	assert(t, SearchLast(list, 9) == 7)
}

func Test_CyclicQueue(t *testing.T) {
	defer guard_ut(t)

	var queue, _ = NewQueue(5)
	for i := 1; i < 8; i++ {
		var err = queue.Push(i)
		if err != nil {
			t.Fail()
		}
	}
	var err = queue.Push(9)
	assert(t, err != nil)
	//1, 2, 3, 4, 5, 6, 7

	for i := 1; i < 5; i++ {
		var key, err = queue.Pop()
		assert(t, err == nil && key == i)
	}
	//5, 6, 7
	for i := 8; i < 12; i++ {
		var err = queue.Push(i)
		assert(t, err == nil)
	}
	assert(t, queue.IsFull())
	//5, 6, 7, 8, 9, 10, 11

	for i := 5; i < 12; i++ {
		var key, err = queue.Pop()
		assert(t, err == nil && key == i)
	}
	assert(t, queue.IsEmpty())
}
