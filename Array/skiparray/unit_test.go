package skiparray

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

type Object struct {
	val int
}

func Test_SkipArray(t *testing.T) {
	guardUT(t)

	obj1, obj2 := Object{val: 1}, Object{val: 2}

	space := NewSkipArray(64 + 8 + 1)
	for i := 0; i < space.Capacity(); i++ {
		assert(t, space.Insert(&obj1) == i)
	}
	assert(t, space.Insert(&obj2) < 0)

	space.Remove(0)
	space.Remove(30)
	space.Remove(63)
	space.Remove(72)
	assert(t, space.Search(0) == nil)
	assert(t, space.Search(30) == nil)
	assert(t, space.Search(63) == nil)
	assert(t, space.Search(72) == nil)

	assert(t, space.Insert(&obj2) == 0)
	assert(t, space.Insert(&obj2) == 30)
	assert(t, space.Insert(&obj2) == 63)
	assert(t, space.Insert(&obj2) == 72)

	it := space.Search(0)
	assert(t, it != nil && it.(*Object) == &obj2)
	it = space.Search(1)
	assert(t, it != nil && it.(*Object) == &obj1)
	it = space.Search(63)
	assert(t, it != nil && it.(*Object) == &obj2)
	it = space.Search(64)
	assert(t, it != nil && it.(*Object) == &obj1)
}
