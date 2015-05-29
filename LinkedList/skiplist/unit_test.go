package skiplist

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_SkipList(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
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
	if dict.Insert(list[size]) {
		t.Fail()
	}
	for i := 0; i < size*2; i++ {
		if !dict.Search(list[i]) {
			t.Fail()
		}
	}

	//遍历
	mark_ut = -(int(^uint(0) >> 1)) - 1
	dict.Travel(checkNum)

	//删除第一份
	for i := 0; i < size; i++ {
		if dict.Remove(list[i]) {
			cnt--
		}
	}
	for i := 0; i < size; i++ {
		if dict.Search(list[i]) {
			t.Fail()
		}
	}

	//删除第二份
	for i := size; i < size*2; i++ {
		if dict.Remove(list[i]) {
			cnt--
		}
	}
	if !dict.IsEmpty() || cnt != 0 {
		t.Fail()
	}
}

var mark_ut = 0

func checkNum(val int) {
	if val < mark_ut {
		fmt.Println("X")
		panic(val)
	}
	mark_ut = val
}
