package chained

import (
	"DSGO/HashTable/hash"
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

func Test_HashTable(t *testing.T) {
	defer guardUT(t)

	var tpl = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var book [52][]byte
	for i := 0; i < 52; i++ {
		book[i] = tpl[i : i+26]
	}

	var table = NewHashTable(hash.BKDRhash)
	for i := 0; i < 52; i++ {
		assert(t, table.Insert(book[i]))
	}
	assert(t, table.Size() == 52)
	assert(t, !table.Insert(book[0]))
	for i := 0; i < 52; i++ {
		assert(t, table.Search(book[i]))
	}
	for i := 0; i < 52; i++ {
		assert(t, table.Remove(book[i]))
	}
	assert(t, table.IsEmpty())
	assert(t, !table.Search(book[0]))
	assert(t, !table.Remove(book[0]))
}
