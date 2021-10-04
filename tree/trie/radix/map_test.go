package radix

import (
	"DSGO/utils"
	"testing"
	"unsafe"
)

func Test_Map(t *testing.T) {
	defer utils.FailInPanic(t)

	var m Map
	for i := 1; i < 18; i++ {
		ptr := new(int)
		*ptr = i
		utils.Assert(t, m.Insert(uint(i), unsafe.Pointer(ptr)))
	}
	utils.Assert(t, !m.Insert(uint(16), unsafe.Pointer(nil)))

	for i := 1; i < 18; i++ {
		ptr := (*int)(m.Search(uint(i)))
		utils.Assert(t, *ptr == i)
	}
	utils.Assert(t, !m.Remove(uint(0)))
	utils.Assert(t, !m.Remove(uint(32)))
	for i := 1; i < 18; i++ {
		utils.Assert(t, m.Remove(uint(i)))
	}
}
