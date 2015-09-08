package chained

import (
	"HashTable/hash"
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

func Test_HashTable(t *testing.T) {
	defer guard_ut(t)

	var tpl = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var book [52]string
	for i := 0; i < 52; i++ {
		book[i] = string(tpl[i : i+26])
	}

	var table = NewHashTable(hash.BKDRhash)
	for i := 0; i < 52; i++ {
		assert(t, table.Insert([]byte(book[i])))
	}
	assert(t, table.Size() == 52)
	assert(t, !table.Insert([]byte(book[0])))
	for i := 0; i < 52; i++ {
		assert(t, table.Search([]byte(book[i])))
	}
	for i := 0; i < 52; i++ {
		assert(t, table.Remove([]byte(book[i])))
	}
	assert(t, table.IsEmpty())
	assert(t, !table.Search([]byte(book[0])))
	assert(t, !table.Remove([]byte(book[0])))
}

func Test_nextSize(t *testing.T) {
	defer guard_ut(t)

	var sz, ok = nextSize(primes[0])
	assert(t, ok && sz == primes[1])

	var last = primes[len(primes)-1]
	for sz = 1; sz <= last; sz *= 2 {
		var next, ok = nextSize(sz)
		assert(t, ok && next > sz)
	}
	next, ok := nextSize(sz)
	assert(t, !ok && next == sz)
}
