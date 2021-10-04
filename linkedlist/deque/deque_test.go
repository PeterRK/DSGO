package deque

import (
	"DSGO/utils"
	"testing"
)

func Test_Deque(t *testing.T) {
	defer utils.FailInPanic(t)
	const size = PieceSize + 1

	dq := New[int]()

	//后增一段
	for i := 0; i < size; i++ {
		dq.PushBack(size*1 + i)
	}

	//前增两段
	for i := 0; i < size; i++ {
		dq.PushFront(size*2 + i)
	}
	for i := 0; i < size; i++ {
		dq.PushFront(size*3 + i)
	}

	//后删两段
	for i := 0; i < size; i++ {
		utils.Assert(t, dq.PopBack() == (size*1+(size-1)-i))
	}
	for i := 0; i < size; i++ {
		utils.Assert(t, dq.PopBack() == (size*2+i))
	}

	//前增一段
	for i := 0; i < size; i++ {
		dq.PushFront(size*4 + i)
	}

	//前删两段
	for i := 0; i < size; i++ {
		utils.Assert(t, dq.PopFront() == (size*4+(size-1)-i))
	}
	for i := 0; i < size; i++ {
		utils.Assert(t, dq.PopFront() == (size*3+(size-1)-i))
	}

	//后增一段
	for i := 0; i < size; i++ {
		dq.PushBack(size*5 + i)
	}
	//前增一段
	for i := 0; i < size; i++ {
		dq.PushFront(size*6 + i)
	}
	//后删一段
	for i := 0; i < size; i++ {
		utils.Assert(t, dq.PopBack() == (size*5+(size-1)-i))
	}
	//前删一段
	for i := 0; i < size; i++ {
		utils.Assert(t, dq.PopFront() == (size*6+(size-1)-i))
	}

	utils.Assert(t, dq.IsEmpty())
}

func Test_Stack(t *testing.T) {
	defer utils.FailInPanic(t)
	const size = 100
	c := NewStack[int]()
	for i := 0; i < size; i++ {
		c.Push(i)
	}
	utils.Assert(t, c.Top() == size-1)
	for i := 0; i < size; i++ {
		utils.Assert(t, c.Pop() == (size-1)-i)
	}
	utils.Assert(t, c.IsEmpty())
}

func Test_Queue(t *testing.T) {
	defer utils.FailInPanic(t)
	const size = 100
	c := NewQueue[int]()
	for i := 0; i < size; i++ {
		c.Push(i)
	}
	utils.Assert(t, c.Front() == 0)
	utils.Assert(t, c.Back() == size-1)
	for i := 0; i < size; i++ {
		utils.Assert(t, c.Pop() == i)
	}
}
