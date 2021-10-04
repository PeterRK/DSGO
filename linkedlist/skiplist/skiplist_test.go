package skiplist

import (
	"DSGO/array"
	"DSGO/utils"
	"math"
	"math/rand"
	"testing"
	"time"
)

func genRand(size int) []int {
	rand.Seed(time.Now().Unix())
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Int()
	}
	return list
}

func Test_SkipList(t *testing.T) {
	defer utils.FailInPanic(t)

	dict := NewSkipList[int]()
	const size = 300
	list := genRand(size * 2)

	cnt := 0

	//插入两份
	for i := 0; i < size*2; i++ {
		if dict.Insert(list[i]) {
			cnt++
		}
	}
	utils.Assert(t, !dict.Insert(list[size]))
	for i := 0; i < size*2; i++ {
		utils.Assert(t, dict.Search(list[i]))
	}

	//遍历
	mark := math.MinInt
	dict.Travel(func(val int) {
		if val < mark {
			panic(val)
		}
		mark = val
	})

	//删除第一份
	for i := 0; i < size; i++ {
		if dict.Remove(list[i]) {
			cnt--
		}
	}
	for i := 0; i < size; i++ {
		utils.Assert(t, !dict.Search(list[i]))
	}

	//删除第二份
	for i := size; i < size*2; i++ {
		if dict.Remove(list[i]) {
			cnt--
		}
	}
	utils.Assert(t, dict.IsEmpty() && cnt == 0)
	utils.Assert(t, !dict.Remove(0))
}

func Benchmark_Insert(b *testing.B) {
	b.StopTimer()
	dict := NewSkipList[int]()
	list := genRand(b.N)
	array.RandomShuffle(list)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		dict.Insert(list[i])
	}
}

func Benchmark_Remove(b *testing.B) {
	b.StopTimer()
	dict := NewSkipList[int]()
	list := genRand(b.N)
	for i := 0; i < b.N; i++ {
		dict.Insert(list[i])
	}
	array.RandomShuffle(list)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		dict.Remove(list[i])
	}
}

func Benchmark_Search(b *testing.B) {
	b.StopTimer()
	dict := NewSkipList[int]()
	list := genRand(b.N)
	for i := 0; i < b.N; i++ {
		dict.Insert(list[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		dict.Search(list[i])
	}
}
