package array

import (
	"DSGO/utils"
	"testing"
)

func Test_CyclicQueue(t *testing.T) {
	defer utils.FailInPanic(t)

	q := NewQueue[int](5)
	for i := 1; i < 8; i++ {
		q.Push(i)
	}
	utils.Assert(t, q.IsFull())

	for i := 1; i < 5; i++ {
		utils.Assert(t, q.Pop() == i)
	}

	for i := 8; i < 12; i++ {
		q.Push(i)
	}
	utils.Assert(t, q.IsFull())

	utils.Assert(t, q.Front() == 5)
	utils.Assert(t, q.Back() == 11)

	for i := 5; i < 12; i++ {
		utils.Assert(t, q.Pop() == i)
	}
	utils.Assert(t, q.IsEmpty())
}
