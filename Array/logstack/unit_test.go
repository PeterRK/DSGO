package logstack

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

func Test_LogStack(t *testing.T) {
	guardUT(t)

	var ls = logStack{limit: 4}
	for i := 0; i < 20; i += 2 {
		ls.Insert(i)
	}
	for i := 0; i < 10; i++ {
		ls.Delete(i)
	}
	for i := 10; i < 20; i++ {
		ls.Insert(i)
	}
	for i := 1; i < 20; i += 2 {
		ls.Delete(i)
	}
	ls.Insert(20)

	assert(t, ls.Search(20))
	assert(t, !ls.Search(19))
	assert(t, ls.Search(18))
	assert(t, ls.Search(12))
	assert(t, !ls.Search(7))
	assert(t, !ls.Search(0))
	assert(t, !ls.Search(99))
}
