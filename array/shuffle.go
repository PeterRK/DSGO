package array

import (
	"constraints"
	"math/rand"
	"time"
)

// 随机排列
func RandomShuffle[T constraints.Ordered](list []T) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(list); i++ {
		j := rand.Int() % (i + 1)
		list[i], list[j] = list[j], list[i]
	}
}

// 随机组合，前n项为结果
func RandomPartition[T constraints.Ordered](list []T, n int) {
	if n > 0 && n < len(list) {
		rand.Seed(time.Now().UnixNano())
		for i := n; i < len(list); i++ {
			j := rand.Int() % (i + 1)
			list[i], list[j] = list[j], list[i]
		}
	}
}

func Reverse[T any](list []T) {
	for l, r := 0, len(list)-1; l < r; {
		list[l], list[r] = list[r], list[l]
		l++
		r--
	}
}

func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, size := 0, len(a); i < size; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func InsertTo[T any](list []T, pos int, val T) []T {
	if pos < 0 || pos > len(list) {
		panic("illegal pos")
	}
	list = append(list, val)
	for i := len(list)-1; i > pos; i-- {
		list[i] = list[i-1]
	}
	list[pos] = val
	return list
}

func EraseFrom[T any](list []T, pos int, keepOrder bool) []T {
	if pos < 0 || pos >= len(list) {
		panic("illegal pos")
	}
	last := len(list) - 1
	if keepOrder {
		for i := pos; i < last; i++ {
			list[i] = list[i+1]
		}
	} else {
		list[pos] = list[last]
	}
	return list[:last]
}

func SetAll[T any](vec []T, val T) {
	for i := 0; i < len(vec); i++ {
		vec[i] = val
	}
}