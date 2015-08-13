package deque

import (
	"testing"
)

func Test_Deque(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
	const size = piece_sz + 1

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
		if err != nil || key != (size*1+(size-1)-i) {
			t.Fail()
		}
	}
	for i := 0; i < size; i++ {
		var key, err = con.PopBack()
		if err != nil || key != (size*2+i) {
			t.Fail()
		}
	}

	//前增一段
	for i := 0; i < size; i++ {
		con.PushFront(size*4 + i)
	}

	//前删两段
	for i := 0; i < size; i++ {
		var key, err = con.PopFront()
		if err != nil || key != (size*4+(size-1)-i) {
			t.Fail()
		}
	}
	for i := 0; i < size; i++ {
		var key, err = con.PopFront()
		if err != nil || key != (size*3+(size-1)-i) {
			t.Fail()
		}
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
		if err != nil || key != (size*5+(size-1)-i) {
			t.Fail()
		}
	}
	//前删一段
	for i := 0; i < size; i++ {
		var key, err = con.PopFront()
		if err != nil || key != (size*6+(size-1)-i) {
			t.Fail()
		}
	}
}

func Test_Stack(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
	const size = 200
	var con = NewStack()
	for i := 0; i < size; i++ {
		con.Push(i)
	}
	for i := 0; i < size; i++ {
		var key, err = con.Pop()
		if err != nil || key != (size-1)-i {
			t.Fail()
		}
	}
}

func Test_Queue(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
	const size = 20
	var con = NewQueue()
	for i := 0; i < size; i++ {
		con.Push(i)
	}
	for i := 0; i < size; i++ {
		var key, err = con.Pop()
		if err != nil || key != i {
			t.Fail()
		}
	}
}
