package ring

import (
	"DSGO/utils"
	"testing"
)

func Test_LruCache(t *testing.T) {
	defer utils.FailInPanic(t)

	c := NewCacheDetail[int, string](5, 4, 0)
	utils.Assert(t, c != nil)

	c.Put(1, "A")
	c.Put(2, "B")
	c.Put(3, "C")
	c.Put(4, "D")
	c.Put(5, "E")
	c.Put(6, "F")
	c.Put(5, "e")
	c.Discard(4)
	_, got := c.Get(4)
	utils.Assert(t, !got)

	val, got := c.Get(6)
	utils.Assert(t, got && val == "F")
	val, got = c.Get(5)
	utils.Assert(t, got && val == "e")
	val, got = c.Get(1)
	utils.Assert(t, !got)
	val, got = c.Get(2)
	utils.Assert(t, !got)
	val, got = c.Get(3)
	utils.Assert(t, got && val == "C")

	c.Put(7, "G")
	c.Put(8, "H")
	c.Put(9, "I")
	c.Put(10, "J")

	val, got = c.Get(7)
	utils.Assert(t, got && val == "G")
	val, got = c.Get(8)
	utils.Assert(t, got && val == "H")
	val, got = c.Get(9)
	utils.Assert(t, got && val == "I")
	val, got = c.Get(10)
	utils.Assert(t, got && val == "J")

	c.Put(1, "A")
	c.Put(2, "B")

	c.Put(9, "i")
	c.Discard(10)

	c.Put(6, "F")
	val, got = c.Get(6)
	c.Put(5, "E")
	c.Put(4, "D")

	val, got = c.Get(3)
	utils.Assert(t, got && val == "C")

	val, got = c.Get(9)
	utils.Assert(t, got && val == "i")
	val, got = c.Get(8)
	utils.Assert(t, got && val == "H")
	val, got = c.Get(7)
	utils.Assert(t, got && val == "G")
	val, got = c.Get(6)
	utils.Assert(t, got && val == "F")
	val, got = c.Get(5)
	utils.Assert(t, got && val == "E")
	val, got = c.Get(4)
	utils.Assert(t, got && val == "D")

	val, got = c.Get(2)
	utils.Assert(t, got && val == "B")
	val, got = c.Get(1)
	utils.Assert(t, got && val == "A")
}
