package rbtreex

import (
	"math/rand"
	"testing"
	"time"
)

//func init() {}	//这用于初始化包中数据，这里不合适

var tree_ut = Tree{root: NULL}
var dirty_ut = false

const size_ut = 200

var list1_ut [size_ut]int
var list2_ut [size_ut]int
var list3_ut [size_ut]int
var list4_ut [size_ut]int

func prepareTestData() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < size_ut; i++ {
		list1_ut[i] = rand.Int()
		list2_ut[i] = rand.Int()
		list3_ut[i] = rand.Int()
		list4_ut[i] = rand.Int()
	}
}

func detectAndMark(t *testing.T) {
	if err := recover(); err != nil {
		dirty_ut = true
		t.Fail()
	}
}
func markFail(t *testing.T) {
	dirty_ut = true
	t.FailNow()
}

func Test_Insert(t *testing.T) {
	defer detectAndMark(t)
	prepareTestData()

	for i := 0; i < size_ut; i++ {
		tree_ut.Insert(list1_ut[i])
		tree_ut.Insert(list2_ut[i])
		tree_ut.Insert(list3_ut[i])
		tree_ut.Insert(list4_ut[i])
	}
	for i := 0; i < size_ut; i++ {
		if tree_ut.Insert(list2_ut[i]) {
			markFail(t)
		}
	}
}

func Test_Find(t *testing.T) {
	if dirty_ut {
		t.SkipNow()
	}
	defer detectAndMark(t)

	for i := 0; i < size_ut; i++ {
		if !tree_ut.Find(list4_ut[i]) ||
			!tree_ut.Find(list3_ut[i]) ||
			!tree_ut.Find(list2_ut[i]) ||
			!tree_ut.Find(list1_ut[i]) {
			markFail(t)
		}
	}
}

func Test_Remove(t *testing.T) {
	if dirty_ut {
		t.SkipNow()
	}
	defer detectAndMark(t)

	for i := 0; i < size_ut; i++ {
		tree_ut.Remove(list3_ut[i])
		tree_ut.Remove(list4_ut[i])
	}
	for i := 0; i < size_ut; i++ {
		tree_ut.Remove(list1_ut[i])
		tree_ut.Remove(list2_ut[i])
	}
}

func Test_InsertAfterRemove(t *testing.T) {
	if dirty_ut {
		t.SkipNow()
	}
	defer detectAndMark(t)

	for i := 0; i < size_ut; i++ {
		tree_ut.Insert(list1_ut[i])
		tree_ut.Insert(list3_ut[i])
	}
}

func Test_FindAfterRemove(t *testing.T) {
	if dirty_ut {
		t.SkipNow()
	}
	defer detectAndMark(t)

	for i := 0; i < size_ut; i++ {
		if !tree_ut.Find(list1_ut[i]) {
			markFail(t)
		}
	}
	for i := 0; i < size_ut; i++ {
		if tree_ut.Find(list2_ut[i]) ||
			tree_ut.Find(list4_ut[i]) {
			markFail(t)
		}
	}
	for i := 0; i < size_ut; i++ {
		if !tree_ut.Find(list3_ut[i]) {
			markFail(t)
		}
	}
}
