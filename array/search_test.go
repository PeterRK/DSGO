package array

import (
	"DSGO/utils"
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	defer utils.FailInPanic(t)
	list := []int{2, 2, 4, 6, 6, 6, 8, 8}

	utils.Assert(t, SearchFirstGE(list, 1) == 0)
	utils.Assert(t, SearchFirstGE(list, 2) == 0)
	utils.Assert(t, SearchFirstGE(list, 3) == 2)
	utils.Assert(t, SearchFirstGE(list, 6) == 3)
	utils.Assert(t, SearchFirstGE(list, 9) == 8)
	utils.Assert(t, SearchLastLE(list, 1) == -1)
	utils.Assert(t, SearchLastLE(list, 5) == 2)
	utils.Assert(t, SearchLastLE(list, 6) == 5)
	utils.Assert(t, SearchLastLE(list, 8) == 7)
	utils.Assert(t, SearchLastLE(list, 9) == 7)
}
