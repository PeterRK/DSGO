package perfect

import (
	"DSGO/utils"
	//"fmt"
	"testing"
)

func Test_Hasher(t *testing.T) {
	//defer utils.FailInPanic(t)

	tpl := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var keys [52]string
	for i := 0; i < len(keys); i++ {
		keys[i] = tpl[i : i+26]
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
