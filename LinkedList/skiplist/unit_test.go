package skiplist

import (
	"fmt"
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

func Test_SkipList(t *testing.T) {
	defer guardUT(t)

	var dict = NewSkipList()
	var cnt = 0
	const size = 300
	var list [size * 2]int

	rand.Seed(time.Now().Unix())
	for i := 0; i < size*2; i++ {
		list[i] = rand.Int()
	}

	//插入两份
	for i := 0; i < size*2; i++ {
		if dict.Insert(list[i]) {
			cnt++
		}
	}
	assert(t, !dict.Insert(list[size]))
	for i := 0; i < size*2; i++ {
		assert(t, dict.Search(list[i]))
	}

	//遍历
	utMark = -(int(^uint(0) >> 1)) - 1
	dict.Travel(checkNum)

	//删除第一份
	for i := 0; i < size; i++ {
		if dict.Remove(list[i]) {
			cnt--
		}
	}
	for i := 0; i < size; i++ {
		assert(t, !dict.Search(list[i]))
	}

	//删除第二份
	for i := size; i < size*2; i++ {
		if dict.Remove(list[i]) {
			cnt--
		}
	}
	assert(t, dict.IsEmpty() && cnt == 0)
	assert(t, !dict.Remove(0))
}

var utMark = 0

func checkNum(val int) {
	if val < utMark {
		fmt.Println("X")
		panic(val)
	}
	utMark = val
}
