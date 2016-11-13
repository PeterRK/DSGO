package lrucache

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

func Test_LRUCache(t *testing.T) {
	guardUT(t)

	var c = New(5, 4)
	assert(t, c != nil)

	c.Insert(1, "A")
	c.Insert(2, "B")
	c.Insert(3, "C")
	c.Insert(4, "D")
	c.Insert(5, "E")
	c.Insert(6, "F")
	c.Insert(5, "e")
	c.Remove(4)
	var val, ok = c.Search(4)
	assert(t, !ok)

	val, ok = c.Search(6)
	assert(t, ok && val == "F")
	val, ok = c.Search(5)
	assert(t, ok && val == "e")
	val, ok = c.Search(1)
	assert(t, !ok)
	val, ok = c.Search(2)
	assert(t, !ok)
	val, ok = c.Search(3)
	assert(t, ok && val == "C")

	c.Insert(7, "G")
	c.Insert(8, "H")
	c.Insert(9, "I")
	c.Insert(10, "J")

	val, ok = c.Search(7)
	assert(t, ok && val == "G")
	val, ok = c.Search(8)
	assert(t, ok && val == "H")
	val, ok = c.Search(9)
	assert(t, ok && val == "I")
	val, ok = c.Search(10)
	assert(t, ok && val == "J")

	val, ok = c.Search(6)
	assert(t, !ok)
	val, ok = c.Search(5)
	assert(t, !ok)

	c.Insert(1, "A")
	c.Insert(2, "B")

	c.Insert(9, "i")
	c.Remove(10)

	c.Insert(6, "F")
	val, ok = c.Search(6)
	c.Insert(5, "E")
	c.Insert(4, "D")

	val, ok = c.Search(3)
	assert(t, ok && val == "C")

	val, ok = c.Search(9)
	assert(t, ok && val == "i")
	val, ok = c.Search(8)
	assert(t, ok && val == "H")
	val, ok = c.Search(7)
	assert(t, ok && val == "G")
	val, ok = c.Search(6)
	assert(t, ok && val == "F")
	val, ok = c.Search(5)
	assert(t, ok && val == "E")
	val, ok = c.Search(4)
	assert(t, ok && val == "D")

	val, ok = c.Search(2)
	assert(t, ok && val == "B")
	val, ok = c.Search(1)
	assert(t, ok && val == "A")
}
