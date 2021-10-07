package bloomfilter

import (
	"DSGO/utils"
	"testing"
)

func Test_BloomFliter(t *testing.T) {
	defer utils.FailInPanic(t)

	tpl := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var keys [52]string
	for i := 0; i < len(keys); i++ {
		keys[i] = tpl[i : i+26]
	}

	bf := New(uint32(len(keys)))
	for i := 0; i < len(keys); i++ {
		bf.Insert(keys[i])
	}
	utils.Assert(t, bf.Size() > 0 && bf.Size() <= len(keys))
	for i := 0; i < len(keys); i++ {
		utils.Assert(t, bf.Search(keys[i]))
	}
}
