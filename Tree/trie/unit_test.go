package trie

import (
	"testing"
)

func assert(t *testing.T, state bool) {
	if !state {
		t.Fail()
	}
}
func guard_ut(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

func Test_Trie(t *testing.T) {
	defer guard_ut(t)

	var trie = NewTrie()
	trie.Insert([]byte("ABCDEFG"))
	trie.Insert([]byte("ABCD1234"))
	trie.Insert([]byte("A987654"))
	trie.Insert([]byte("ABC"))
	trie.Insert([]byte("abcdefghijklmn"))
	trie.Insert([]byte("abcdef0987654321"))
	trie.Insert([]byte("abcde"))
	trie.Insert([]byte("ABCDEFG"))

	assert(t, trie.Search([]byte("ABCD")) == 0)
	assert(t, trie.Search([]byte("ABCD123")) == 0)
	assert(t, trie.Search([]byte("ABCDEFG")) == 2)
	assert(t, trie.Search([]byte("ABCD1234")) == 1)
	assert(t, trie.Search([]byte("ABC")) == 1)
	assert(t, trie.Search([]byte("abcdefghijklmn")) == 1)

	trie.Remove([]byte("ABCD"), false)
	trie.Remove([]byte("ABCD123"), false)
	trie.Remove([]byte("ABC"), false)
	assert(t, trie.Search([]byte("ABCDE")) == 0)

	trie.Remove([]byte("abcdefghijklmn"), false)
	assert(t, trie.Search([]byte("abcdef0987654321")) == 1)

	trie.Remove([]byte("abcdef0987654321"), false)
	trie.Remove([]byte("ABCDEFG"), true)
	trie.Remove([]byte("ABCD1234"), false)
	trie.Remove([]byte("zzzzz"), true)

	assert(t, trie.Search([]byte{}) == 1)
	trie.Insert([]byte{})
	trie.Remove([]byte{}, false)
}
