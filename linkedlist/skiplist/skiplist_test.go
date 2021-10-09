package skiplist

import (
	"DSGO/utils"
	"math"
	"testing"
)

func Test_SkipList(t *testing.T) {
	defer utils.FailInPanic(t)

	set := New[int]()
	const size = 3000
	list := utils.RandomArray[int](size * 2)

	cnt := 0

	//插入两份
	for i := 0; i < size*2; i++ {
		if set.Insert(list[i]) {
			cnt++
		}
	}
	utils.Assert(t, set.Size() == cnt)
	utils.Assert(t, !set.Insert(list[size]))
	for i := 0; i < size*2; i++ {
		utils.Assert(t, set.Search(list[i]))
	}

	//遍历
	mark := math.MinInt
	set.Travel(func(val int) {
		if val < mark {
			panic(val)
		}
		mark = val
	})

	//删除第一份
	for i := 0; i < size; i++ {
		if set.Remove(list[i]) {
			cnt--
		}
	}
	for i := 0; i < size; i++ {
		utils.Assert(t, !set.Search(list[i]))
	}
	utils.Assert(t, set.Size() == cnt)

	//删除第二份
	for i := size; i < size*2; i++ {
		if set.Remove(list[i]) {
			cnt--
		}
	}
	utils.Assert(t, set.IsEmpty() && cnt == 0)
	utils.Assert(t, !set.Remove(0))
}

func Benchmark_Insert(b *testing.B) {
	b.StopTimer()
	set := New[int]()
	list := utils.RandomArray[int](b.N)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.Insert(list[i])
	}
}

func Benchmark_Search(b *testing.B) {
	b.StopTimer()
	set := New[int]()
	list := utils.RandomArray[int](b.N)
	for i := 0; i < b.N; i++ {
		set.Insert(list[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.Search(list[i])
	}
}

func Benchmark_Remove(b *testing.B) {
	b.StopTimer()
	set := New[int]()
	list := utils.RandomArray[int](b.N)
	for i := 0; i < b.N; i++ {
		set.Insert(list[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		set.Remove(list[i])
	}
}
