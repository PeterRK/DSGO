package rbt

import (
	"math/rand"
	"testing"
	"time"
)

func Test_Tree(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
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
	if tree.Insert(list[size]) {
		t.Fail()
	}
	for i := 0; i < size*2; i++ {
		if !tree.Search(list[i]) {
			t.Fail()
		}
	}

	//删除第一份
	for i := 0; i < size; i++ {
		if tree.Remove(list[i]) {
			cnt--
		}
	}
	for i := 0; i < size; i++ {
		if tree.Search(list[i]) {
			t.Fail()
		}
	}

	//删除第二份
	for i := size; i < size*2; i++ {
		if tree.Remove(list[i]) {
			cnt--
		}
	}
	if !tree.IsEmpty() || cnt != 0 {
		t.Fail()
	}
}
