package array

import (
	"golang.org/x/exp/constraints"
)

// 寻找key的位置(未必是第一个)，没有返回-1
func Search[T constraints.Ordered](list []T, key T) int {
	for a, b := 0, len(list); a < b; {
		m := a + (b-a)/2 //(a+b)/2
		switch {
		case key > list[m]:
			a = m + 1
		case key < list[m]:
			b = m
		default:
			return m
		}
	}
	return -1
}

// 在由小到大的序列中寻找第一个大于key的位置
func SearchSuccessor[T constraints.Ordered](list []T, key T) int {
	a, b := 0, len(list)
	for a < b {
		m := a + (b-a)/2 //(a+b)/2
		if key < list[m] {
			b = m
		} else {
			a = m + 1
		}
	}
	return a
}

// 在由小到大的序列中寻找第一个大于或等于key的位置
func SearchFirstGE[T constraints.Ordered](list []T, key T) int {
	a, b := 0, len(list)
	for a < b {
		m := a + (b-a)/2 //(a+b)/2
		if key > list[m] {
			a = m + 1
		} else {
			b = m
		}
	}
	return a
}

// 在由小到大的序列中寻找最后一个小于或等于key的位置
func SearchLastLE[T constraints.Ordered](list []T, key T) int {
	a, b := len(list)-1, -1
	for a > b {
		m := a + (b-a+1)/2 //(a+b+1)/2，(a+b+2)/2也可以，但(a+b)/2+1不行
		if key < list[m] {
			a = m - 1
		} else {
			b = m
		}
	}
	return a
}

// 在由小到大的序列中寻找目标，找打返回索引范围，没有则返回false
func SearchRange[T constraints.Ordered](list []T, key T) (first, last int, ok bool) {
	last = SearchLastLE(list, key)
	if last == -1 || list[last] != key {
		return -1, -1, false
	}
	first = SearchFirstGE(list, key)
	return first, last, true
}

// 向有序数组插入值
func Insert[T constraints.Ordered](list []T, key T) []T {
	spot := SearchSuccessor(list, key)
	list = append(list, key)
	for i := len(list) - 1; i > spot; i-- {
		list[i] = list[i-1] //后移
	}
	list[spot] = key
	return list
}

// 选取第k大数
func Pick[T constraints.Ordered](list []T, k int) T {
	if k <= 0 || k > len(list) {
		panic("out of range")
	}
	for begin, end := 0, len(list); begin < end-1; {
		pivot := list[(begin+end)/2] //一定要选偏后
		a, b := begin, end-1
		for { //注意对称性
			for list[a] < pivot {
				a++
			}
			for list[b] > pivot {
				b--
			}
			if a >= b {
				break
			}
			list[a], list[b] = list[b], list[a]
			a++
			b--
		}
		if k <= a {
			end = a
		} else {
			begin = a
		}
	}
	return list[k-1]
}
