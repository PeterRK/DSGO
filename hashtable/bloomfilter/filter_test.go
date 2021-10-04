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
	for i := 0; i < 52; i++ {
		keys[i] = tpl[i : i+26]
	}

	var bf BloomFliter
	bf.Init(52)
	for i := 0; i < 52; i++ {
		bf.Insert(keys[i])
	}
	utils.Assert(t, bf.Item() > 0 && bf.Item() <= 52)
	for i := 0; i < 52; i++ {
		utils.Assert(t, bf.Search(keys[i]))
		utils.Assert(t, !bf.Insert(keys[i]))
	}
}
