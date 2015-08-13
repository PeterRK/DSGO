package array

import (
	"testing"
)

func guard_ut(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

func Test_BinarySearch(t *testing.T) {
	defer guard_ut(t)
	var list = []int{2, 2, 4, 6, 6, 6, 8, 8}

	if SearchFirst(list, 1) != 0 ||
		SearchFirst(list, 2) != 0 ||
		SearchFirst(list, 3) != 2 ||
		SearchFirst(list, 6) != 3 ||
		SearchFirst(list, 9) != 8 ||
		SearchLast(list, 1) != -1 ||
		SearchLast(list, 5) != 2 ||
		SearchLast(list, 6) != 5 ||
		SearchLast(list, 8) != 7 ||
		SearchLast(list, 9) != 7 {
		t.Fail()
	}
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
	if err == nil {
		t.Fail()
	}
	//1, 2, 3, 4, 5, 6, 7

	for i := 1; i < 5; i++ {
		var key, err = queue.Pop()
		if err != nil || key != i {
			t.Fail()
		}
	}
	//5, 6, 7
	for i := 8; i < 12; i++ {
		var err = queue.Push(i)
		if err != nil {
			t.Fail()
		}
	}
	if !queue.IsFull() {
		t.Fail()
	}
	//5, 6, 7, 8, 9, 10, 11

	for i := 5; i < 12; i++ {
		var key, err = queue.Pop()
		if err != nil || key != i {
			t.Fail()
		}
	}
	if !queue.IsEmpty() {
		t.Fail()
	}
}
