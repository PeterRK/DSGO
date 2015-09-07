package list

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

func Test_Ring(t *testing.T) {
	defer guard_ut(t)

	var space [10]NodeX
	for i := 0; i < 10; i++ {
		space[i].Val = i
	}

	var ring Ring
	ring.Initialize()

	for i := 4; i >= 0; i-- {
		ring.InsertHead(&space[i])
	}
	for i := 5; i < 10; i++ {
		ring.InsertTail(&space[i])
	}

	Release(&space[8])
	ring.InsertTail(&space[8])
	Release(&space[9])
	ring.InsertHead(&space[9])
	Release(&space[8])
	ring.InsertTail(&space[8])
	Release(&space[0])
	ring.InsertTail(&space[0])

	//9 1 2 3 4 5 6 7 8 0
	var node = ring.Head()
	assert(t, node != nil && node.Val == 9)
	node = ring.Tail()
	assert(t, node != nil && node.Val == 0)

	node = ring.PopHead()
	assert(t, node != nil && node.Val == 9)
	node = ring.PopTail()
	assert(t, node != nil && node.Val == 0)
	for i := 1; i < 9; i++ {
		node = ring.PopHead()
		assert(t, node != nil && node.Val == i)
	}
	assert(t, ring.IsEmpty())
	assert(t, ring.Head() == nil && ring.Tail() == nil)
	assert(t, ring.PopHead() == nil && ring.PopTail() == nil)
}
