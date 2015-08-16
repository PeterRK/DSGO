package cuckoo

import (
	"HashTable/hash"
	"testing"
)

func Test_HashTable(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
	var tpl = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var book [52]string
	for i := 0; i < 52; i++ {
		book[i] = string(tpl[i : i+26])
	}

	var fn = [WAYS]func([]byte) uint{hash.APhash, hash.FNVhash, hash.JShash}
	var table = NewHashTable(fn)
	for i := 0; i < 52; i++ {
		if !table.Insert([]byte(book[i])) {
			t.Fail()
		}
	}
	for i := 0; i < 52; i++ {
		if !table.Search([]byte(book[i])) {
			t.Fail()
		}
	}
	for i := 0; i < 52; i++ {
		if !table.Remove([]byte(book[i])) {
			t.Fail()
		}
	}
	if !table.IsEmpty() {
		t.Fail()
	}
}
