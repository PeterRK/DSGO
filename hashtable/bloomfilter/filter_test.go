package bloomfilter

import (
	"DSGO/utils"
	"testing"
)

func Test_BloomFliter(t *testing.T) {
	defer utils.FailInPanic(t)

	tpl := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var keys [52 * 25]string
	for k := 0; k < 25; k++ {
		for i := 0; i < 52; i++ {
			keys[i+52*k] = tpl[i : i+k+1]
		}
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
