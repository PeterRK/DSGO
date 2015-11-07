package avlt

import (
	"math/rand"
	"testing"
	"time"
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

func Test_Tree(t *testing.T) {
	defer guardUT(t)

	var tree Tree
	var cnt = 0
	const size = 200
	var list = new([size]int32)

	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		list[i] = int32(rand.Int())
	}

	for i := 0; i < size; i++ {
		if tree.Insert(list[i]) {
			cnt++
		}
	}
	for i := 0; i < size; i++ {
		assert(t, tree.Search(list[i]))
		assert(t, !tree.Insert(list[i]))
	}
	for i := 0; i < size; i++ {
		if tree.Remove(list[i]) {
			cnt--
		}
		assert(t, !tree.Search(list[i]))
	}
	assert(t, tree.IsEmpty() && cnt == 0)
	assert(t, !tree.Remove(0))
}
