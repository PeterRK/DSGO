package ring

import (
	"DSGO/utils"
	"testing"
)

func Test_Ring(t *testing.T) {
	defer utils.FailInPanic(t)

	var space [10]Node[int]
	for i := 0; i < 10; i++ {
		space[i].Val = i
	}

	r := New[int]()

	for i := 4; i >= 0; i-- {
		r.PushHead(&space[i])
	}
	for i := 5; i < 10; i++ {
		r.PushTail(&space[i])
	}

	space[8].Release()
	r.PushTail(&space[8])
	space[9].Release()
	r.PushHead(&space[9])
	space[8].Release()
	r.PushTail(&space[8])
	space[0].Release()
	r.PushTail(&space[0])

	//9 1 2 3 4 5 6 7 8 0
	unit := r.Head()
	utils.Assert(t, unit != nil && unit.Val == 9)
	unit = r.Tail()
	utils.Assert(t, unit != nil && unit.Val == 0)

	unit = r.PopHead()
	utils.Assert(t, unit != nil && unit.Val == 9)
	unit = r.PopTail()
	utils.Assert(t, unit != nil && unit.Val == 0)
	for i := 1; i < 9; i++ {
		unit = r.PopHead()
		utils.Assert(t, unit != nil && unit.Val == i)
	}
	utils.Assert(t, r.IsEmpty())
	utils.Assert(t, r.Head() == nil && r.Tail() == nil)
	utils.Assert(t, r.PopHead() == nil && r.PopTail() == nil)
}
