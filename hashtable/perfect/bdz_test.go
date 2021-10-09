package perfect

import (
	"DSGO/utils"
	"testing"
)

func Test_Hasher(t *testing.T) {
	defer utils.FailInPanic(t)

	tpl := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var keys [52 * 25]string
	for k := 0; k < 25; k++ {
		for i := 0; i < 52; i++ {
			keys[i+52*k] = tpl[i : i+k+1]
		}
	}

	hasher := New(keys[:])
	utils.Assert(t, hasher != nil)

	book := make(map[uint32]bool)
	for i := 0; i < len(keys); i++ {
		code := hasher.Hash(keys[i])
		utils.Assert(t, !book[code])
		book[code] = true
	}

	msg := "GoodLuck"
	hasher = New([]string{})
	utils.Assert(t, hasher != nil)
	hasher.Hash(msg)
	hasher = New([]string{msg})
	utils.Assert(t, hasher != nil)
	hasher.Hash(msg)
}
