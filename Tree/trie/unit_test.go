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

	var trie = NewTrie()
	trie.Insert("ABCDEFG")
	trie.Insert("ABCD1234")
	trie.Insert("A987654")
	trie.Insert("ABC")
	trie.Insert("abcdefghijklmn")
	trie.Insert("abcdef0987654321")
	trie.Insert("abcde")
	trie.Insert("ABCDEFG")

	assert(t, trie.Search("ABCD") == 0)
	assert(t, trie.Search("ABCD123") == 0)
	assert(t, trie.Search("ABCDEFG") == 2)
	assert(t, trie.Search("ABCD1234") == 1)
	assert(t, trie.Search("ABC") == 1)
	assert(t, trie.Search("abcdefghijklmn") == 1)

	trie.Remove("ABCD", false)
	trie.Remove("ABCD123", false)
	trie.Remove("ABC", false)
	assert(t, trie.Search("ABCDE") == 0)

	trie.Remove("abcdefghijklmn", false)
	assert(t, trie.Search("abcdef0987654321") == 1)

	trie.Remove("abcdef0987654321", false)
	trie.Remove("ABCDEFG", true)
	trie.Remove("ABCD1234", false)
	trie.Remove("zzzzz", true)

	assert(t, trie.Search("") == 1)
	trie.Insert("")
	trie.Remove("", false)
}
