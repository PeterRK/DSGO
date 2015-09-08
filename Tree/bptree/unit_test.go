package bptree

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
func guard_ut(t *testing.T) {
	if err := recover(); err != nil {
		t.Fail()
	}
}

func Test_Tree(t *testing.T) {
	defer guard_ut(t)

	var tree Tree
	var cnt = 0
	const size = 5000
	var list = new([size]int)

	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
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

	mark_ut = -(int(^uint(0) >> 1)) - 1
	tree.Travel(checkNum)

	for i := 0; i < size; i++ {
		if tree.Remove(list[i]) {
			cnt--
		}
		assert(t, !tree.Search(list[i]))
	}
	assert(t, tree.IsEmpty() && cnt == 0)
	assert(t, !tree.Remove(0))
}

var mark_ut = 0

func checkNum(val int) {
	if val < mark_ut {
		panic(val)
	}
	mark_ut = val
}
