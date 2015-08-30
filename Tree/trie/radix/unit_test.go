package radix

import (
	"testing"
	"unsafe"
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

func Test_Tree(t *testing.T) {
	defer guard_ut(t)

	var tree Tree
	for i := 1; i < 18; i++ {
		var ptr = new(int)
		*ptr = i
		assert(t, tree.Insert(uint(i), unsafe.Pointer(ptr)))
	}
	assert(t, !tree.Insert(uint(16), unsafe.Pointer(nil)))

	for i := 1; i < 18; i++ {
		var ptr = (*int)(tree.Search(uint(i)))
		assert(t, *ptr == i)
	}
	assert(t, !tree.Remove(uint(0)))
	assert(t, !tree.Remove(uint(32)))
	for i := 1; i < 18; i++ {
		assert(t, tree.Remove(uint(i)))
	}
}
