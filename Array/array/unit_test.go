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

	var queue = NewQueue(5)
	for i := 1; i < 8; i++ {
		var fail = queue.Push(i)
		if fail {
			t.Fail()
		}
	}
	var fail = queue.Push(9)
	if !fail {
		t.Fail()
	}

	for i := 1; i < 5; i++ {
		var key, fail = queue.Pop()
		if fail || key != i {
			t.Fail()
		}
	}
	for i := 8; i < 12; i++ {
		var fail = queue.Push(i)
		if fail {
			t.Fail()
		}
	}
	if !queue.IsFull() {
		t.Fail()
	}

	for i := 5; i < 12; i++ {
		var key, fail = queue.Pop()
		if fail || key != i {
			t.Fail()
		}
	}
	if !queue.IsEmpty() {
		t.Fail()
	}
}
