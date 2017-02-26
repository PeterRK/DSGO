package perfect

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

func Test_Nothing(t *testing.T) {
	defer guardUT(t)

	var tpl = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var book [52]string
	for i := 0; i < 52; i++ {
		book[i] = tpl[i : i+26]
	}

	var table Table
	assert(t, table.Build(book[:]) == nil)
	for i := 0; i < len(book); i++ {
		assert(t, table.Serach(book[i]))
	}
	assert(t, !table.Serach("GoodLuck"))
}
