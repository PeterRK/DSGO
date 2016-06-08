package wavl

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

func randArray(size int) []int {
	rand.Seed(time.Now().Unix())
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
	}
	return list
}
func Test_Tree(t *testing.T) {
	defer guardUT(t)

	var tree Tree
	var cnt = 0
	const size = 200
	var list = randArray(size)

	for i := 0; i < size; i++ {
		if tree.Insert(list[i]) > 0 {
			cnt++
		}
	}

	for i := 0; i < size; i++ {
		assert(t, tree.Search(list[i]) != -1)
		assert(t, tree.Insert(list[i]) < 0)
	}

	for i := 0; i < size; i++ {
		if tree.Remove(list[i]) > 0 {
			cnt--
		}
		assert(t, tree.Search(list[i]) == -1)
	}

	assert(t, tree.IsEmpty() && cnt == 0)
	assert(t, tree.Remove(0) < 0)
}

func randomize(list []int) {
	rand.Seed(time.Now().Unix())
	for i := 1; i < len(list); i++ {
		var j = rand.Int() % (i + 1)
		list[i], list[j] = list[j], list[i]
	}
}
func Test_Rank(t *testing.T) {
	defer guardUT(t)

	var tree Tree
	const size = 200
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i + 1
	}

	randomize(list)
	for i := 0; i < size; i++ {
		assert(t, tree.Insert(list[i]) > 0)
	}
	for i := 0; i < size; i++ {
		assert(t, tree.Search(list[i]) == list[i])
	}
}
