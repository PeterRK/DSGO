package deque

import (
	"testing"
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

func Test_Deque(t *testing.T) {
	defer guardUT(t)
	const size = PIECE_SIZE + 1

	var con = NewDeque()

	//后增一段
	for i := 0; i < size; i++ {
		con.PushBack(size*1 + i)
	}

	//前增两段
	for i := 0; i < size; i++ {
		con.PushFront(size*2 + i)
	}
	for i := 0; i < size; i++ {
		con.PushFront(size*3 + i)
	}

	//后删两段
	for i := 0; i < size; i++ {
		var key, err = con.PopBack()
		assert(t, err == nil && key == (size*1+(size-1)-i))
	}
	for i := 0; i < size; i++ {
		var key, err = con.PopBack()
		assert(t, err == nil && key == (size*2+i))
	}

	//前增一段
	for i := 0; i < size; i++ {
		con.PushFront(size*4 + i)
	}

	//前删两段
	for i := 0; i < size; i++ {
		var key, err = con.PopFront()
		assert(t, err == nil && key == (size*4+(size-1)-i))
	}
	for i := 0; i < size; i++ {
		var key, err = con.PopFront()
		assert(t, err == nil && key == (size*3+(size-1)-i))
	}

	//后增一段
	for i := 0; i < size; i++ {
		con.PushBack(size*5 + i)
	}
	//前增一段
	for i := 0; i < size; i++ {
		con.PushFront(size*6 + i)
	}
	//后删一段
	for i := 0; i < size; i++ {
		var key, err = con.PopBack()
		assert(t, err == nil && key == (size*5+(size-1)-i))
	}
	//前删一段
	for i := 0; i < size; i++ {
		var key, err = con.PopFront()
		assert(t, err == nil && key == (size*6+(size-1)-i))
	}

	var _, err = con.PopFront()
	assert(t, err != nil)
	_, err = con.PopBack()
	assert(t, err != nil)
	_, err = con.Front()
	assert(t, err != nil)
	_, err = con.Back()
	assert(t, err != nil)
	con.Clear()
	assert(t, con.IsEmpty())
}

func Test_Stack(t *testing.T) {
	defer guardUT(t)
	const size = 200
	var con = NewStack()
	for i := 0; i < size; i++ {
		con.Push(i)
	}
	var key, err = con.Top()
	assert(t, err == nil && key == size-1)
	for i := 0; i < size; i++ {
		key, err = con.Pop()
		assert(t, err == nil && key == (size-1)-i)
	}
	assert(t, con.IsEmpty())
}

func Test_Queue(t *testing.T) {
	defer guardUT(t)
	const size = 20
	var con = NewQueue()
	for i := 0; i < size; i++ {
		con.Push(i)
	}
	var key, err = con.Back()
	assert(t, err == nil && key == size-1)
	for i := 0; i < size; i++ {
		var key, err = con.Pop()
		assert(t, err == nil && key == i)
	}
}
