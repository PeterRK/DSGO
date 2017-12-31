package trie

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

func Test_Trie(t *testing.T) {
	defer guardUT(t)

	tr := NewTrie()
	tr.Insert("ABCDEFG")
	tr.Insert("ABCD1234")
	tr.Insert("A987654")
	tr.Insert("ABC")
	tr.Insert("abcdefghijklmn")
	tr.Insert("abcdef0987654321")
	tr.Insert("abcde")
	tr.Insert("ABCDEFG")

	assert(t, tr.Search("ABCD") == 0)
	assert(t, tr.Search("ABCD123") == 0)
	assert(t, tr.Search("ABCDEFG") == 2)
	assert(t, tr.Search("ABCD1234") == 1)
	assert(t, tr.Search("ABC") == 1)
	assert(t, tr.Search("abcdefghijklmn") == 1)

	tr.Remove("ABCD", false)
	tr.Remove("ABCD123", false)
	tr.Remove("ABC", false)
	assert(t, tr.Search("ABCDE") == 0)

	tr.Remove("abcdefghijklmn", false)
	assert(t, tr.Search("abcdef0987654321") == 1)

	tr.Remove("abcdef0987654321", false)
	tr.Remove("ABCDEFG", true)
	tr.Remove("ABCD1234", false)
	tr.Remove("zzzzz", true)

	assert(t, tr.Search("") == 1)
	tr.Insert("")
	tr.Remove("", false)
}
