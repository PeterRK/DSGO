package rbtree

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
	const size = 200
	var list [size * 2]int32

	rand.Seed(time.Now().Unix())
	for i := 0; i < size*2; i++ {
		list[i] = int32(rand.Int())
	}

	//插入两份
	for i := 0; i < size*2; i++ {
		if tree.Insert(list[i]) {
			cnt++
		}
	}
	assert(t, !tree.Insert(list[size]))
	for i := 0; i < size*2; i++ {
		assert(t, tree.Search(list[i]))
	}

	//删除第一份
	for i := 0; i < size; i++ {
		if tree.Remove(list[i]) {
			cnt--
		}
	}
	for i := 0; i < size; i++ {
		assert(t, !tree.Search(list[i]))
	}

	//删除第二份
	for i := size; i < size*2; i++ {
		if tree.Remove(list[i]) {
			cnt--
		}
	}
	assert(t, tree.IsEmpty() && cnt == 0)
}
