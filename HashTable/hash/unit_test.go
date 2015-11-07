package hash

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

func Test_BloomFliter(t *testing.T) {
	var tpl = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var book [52][]byte
	for i := 0; i < 52; i++ {
		book[i] = tpl[i : i+26]
	}

	var fliter BloomFliter
	for i := 0; i < 52; i++ {
		fliter.Insert(book[i])
	}
	for i := 0; i < 52; i++ {
		assert(t, fliter.Search(book[i]))
	}

	assert(t, !fliter.Search([]byte("1234567890987654321")))
}
