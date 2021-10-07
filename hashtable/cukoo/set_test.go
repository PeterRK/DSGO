package cukoo

import (
	"DSGO/utils"
	"testing"
)

func Test_HashSet(t *testing.T) {
	defer utils.FailInPanic(t)

	tpl := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var keys [52]string
	for i := 0; i < len(keys); i++ {
		keys[i] = tpl[i : i+26]
	}

	set := NewSet()
	for i := 0; i < len(keys); i++ {
		utils.Assert(t, set.Insert(keys[i]))
	}
	utils.Assert(t, set.Size() == len(keys))
	utils.Assert(t, !set.Insert(keys[0]))
	for i := 0; i < len(keys); i++ {
		utils.Assert(t, set.Search(keys[i]))
	}
	for i := 0; i < len(keys); i++ {
		utils.Assert(t, set.Remove(keys[i]))
	}
	utils.Assert(t, set.IsEmpty())
	utils.Assert(t, !set.Search(keys[0]))
	utils.Assert(t, !set.Remove(keys[0]))
}
