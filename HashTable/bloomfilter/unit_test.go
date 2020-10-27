package bloomfilter

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
	defer guardUT(t)

	tpl := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var book [52][]byte
	for i := 0; i < 52; i++ {
		book[i] = tpl[i : i+26]
	}

	var bf BloomFliter
	bf.init(52)
	for i := 0; i < 52; i++ {
		bf.Insert(book[i])
	}
	assert(t, bf.Item() > 0 && bf.Item() <= 52)
	for i := 0; i < 52; i++ {
		assert(t, bf.Search(book[i]))
		assert(t, !bf.Insert(book[i]))
	}
}
