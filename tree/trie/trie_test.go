package trie

import (
	"DSGO/utils"
	"testing"
)

func Test_Trie(t *testing.T) {
	defer utils.FailInPanic(t)

	tr := New()
	utils.Assert(t, tr.Insert("ABCDEFG"))
	utils.Assert(t, tr.Insert("ABCD1234"))
	utils.Assert(t, tr.Insert("A987654"))
	utils.Assert(t, tr.Insert("ABC"))
	utils.Assert(t, tr.Insert("abcdefghijklmn"))
	utils.Assert(t, tr.Insert("abcdef0987654321"))
	utils.Assert(t, tr.Insert("abcde"))
	utils.Assert(t, !tr.Insert("ABCDEFG"))

	utils.Assert(t, !tr.Search(""))
	utils.Assert(t, tr.Insert(""))
	utils.Assert(t, tr.Search(""))
	utils.Assert(t, tr.Remove(""))

	utils.Assert(t, !tr.Search("ABCD"))
	utils.Assert(t, !tr.Search("ABCD123"))
	utils.Assert(t, tr.Search("ABCDEFG"))
	utils.Assert(t, tr.Search("ABCD1234"))
	utils.Assert(t, tr.Search("ABC"))
	utils.Assert(t, tr.Search("abcdefghijklmn"))

	utils.Assert(t, !tr.Remove("ABCD"))
	utils.Assert(t, !tr.Remove("ABCD123"))
	utils.Assert(t, tr.Remove("ABC"))
	utils.Assert(t, !tr.Search("ABCDE"))

	utils.Assert(t, tr.Remove("abcdefghijklmn"))
	utils.Assert(t, tr.Search("abcdef0987654321"))
	utils.Assert(t, tr.Remove("abcdef0987654321"))
	utils.Assert(t, tr.Remove("ABCDEFG"))
	utils.Assert(t, tr.Remove("abcde"))
	utils.Assert(t, !tr.Remove("zzzzz"))
}
