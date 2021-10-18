package sort

import (
	"DSGO/array"
	"constraints"
)

// 快速排序，改进的冒泡排序，不具有稳定性。
// 平均复杂度为O(NlogN) & O(logN)，最坏情况是O(N^2) & O(N)。
// 其中比较操作是O(NlogN)，常数与MergeSort相当；挪移操作是O(NlogN)，常数小于MergeSort。
// 这里采用递归实现，但QuickSort不适合递归实现(有爆栈风险)。
func QuickSort[T constraints.Ordered](list []T) {
	if len(list) < lowerBound {
		SimpleSort(list)
	} else {
		m := partition(list)
		QuickSort(list[:m])
		QuickSort(list[m:])
	}
}

type pair struct {
	a int
	b int
}

func QuickSort_v2[T constraints.Ordered](list []T) {
	var tasks array.Stack[pair]
	tasks.Push(pair{0, len(list)})
	for !tasks.IsEmpty() {
		r := tasks.Pop()
		if r.b-r.a < lowerBound {
			SimpleSort(list[r.a:r.b])
		} else {
			m := partition(list[r.a:r.b]) + r.a
			tasks.Push(pair{m, r.b})
			tasks.Push(pair{r.a, m})
		}
	}
}

func sort3[T constraints.Ordered](list []T, a, b, c int) (int, int, int) {
	if list[a] < list[b] {
		if list[b] < list[c] {
			return a, b, c
		} else if list[a] < list[c] {
			return a, c, b
		} else {
			return c, a, b
		}
	} else {
		if list[a] < list[c] {
			return b, a, c
		} else if list[b] < list[c] {
			return b, c, a
		} else {
			return c, b, a
		}
	}
}

func partition[T constraints.Ordered](list []T) int {
	size := len(list)
	m, s := size/2, size/4
	a, m, b := sort3(list, m-s, m, m+s)
	s = size - 1
	pivot := list[m]
	list[0], list[a] = list[a], list[0]
	list[s], list[b] = list[b], list[s]

	a, b = 1, s-1
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
	return a
}

// 三分快速排序，比二分版本略为复杂
func QuickSortY[T constraints.Ordered](list []T) {
	for len(list) > lowerBoundY {
		fst, snd := triPartition(list)
		QuickSortY(list[:fst])
		QuickSortY(list[snd+1:])
		if list[fst] == list[snd] {
			return
		}
		list = list[fst+1 : snd]
	}
	SimpleSort(list)
}

func sort5[T constraints.Ordered](list []T, a, b, c, d, e int) (int, int, int, int, int) {
	if list[b] < list[a] {
		a, b = b, a
	}
	if list[d] < list[c] {
		c, d = d, c
	}
	if list[c] < list[a] {
		a, c = c, a
		b, d = d, b
	}
	if list[c] < list[e] {
		if list[d] < list[e] {
			if list[b] < list[d] {
				if list[c] < list[b] {
					return a, c, b, d, e
				} else {
					return a, b, c, d, e
				}
			} else if list[b] < list[e] {
				return a, c, d, b, e
			} else {
				return a, c, d, e, b
			}
		} else {
			if list[b] < list[e] {
				if list[c] < list[b] {
					return a, c, b, e, d
				} else {
					return a, b, c, e, d
				}
			} else if list[b] < list[d] {
				return a, c, e, b, d
			} else {
				return a, c, e, d, b
			}
		}
	} else {
		if list[b] < list[c] {
			if list[e] < list[a] {
				return e, a, b, c, d
			} else if list[e] < list[b] {
				return a, e, b, c, d
			} else {
				return a, b, e, c, d
			}
		} else {
			if list[a] < list[e] {
				a, e = e, a
			}
			if list[d] < list[b] {
				b, d = d, b
			}
			return e, a, c, b, d
		}
	}
}

// 返回两个分界元素的位置
func triPartition[T constraints.Ordered](list []T) (fst, snd int) {
	size := len(list)
	m, s := size/2, size/4
	x, l, _, r, y := sort5(list, m-s, m-1, m, m+1, m+s)

	s = size - 1
	pivotL, pivotR := list[l], list[r]
	list[l], list[r] = list[0], list[s]
	list[1], list[x] = list[x], list[1]
	list[s-1], list[y] = list[y], list[s-1]

	l, r = 2, s-2
	for {
		for list[l] < pivotL {
			l++
		}
		for list[r] > pivotR {
			r--
		}
		if list[l] > pivotR {
			list[l], list[r] = list[r], list[l]
			r--
			if list[l] < pivotL {
				l++
				continue
			}
		}
		break
	}

	for k := l + 1; k <= r; k++ {
		if list[k] > pivotR {
			for list[r] > pivotR {
				r--
			}
			if k >= r {
				break
			}
			if list[r] < pivotL {
				list[l], list[k], list[r] = list[r], list[l], list[k]
				l++
			} else {
				list[k], list[r] = list[r], list[k]
			}
			r--
		} else if list[k] < pivotL {
			list[k], list[l] = list[l], list[k]
			l++
		}
	}

	l--
	r++
	list[0], list[l] = list[l], pivotL
	list[s], list[r] = list[r], pivotR
	return l, r
}

func QuickSortY_v2[T constraints.Ordered](list []T) {
	for len(list) > lowerBoundY {
		l, r, skip := triPartition_v2(list)
		QuickSortY_v2(list[:l])
		QuickSortY_v2(list[r:])
		if skip {
			return
		}
		list = list[l:r]
	}
	SimpleSort(list)
}

func sort4[T constraints.Ordered](list []T, a, b, c, d int) (int, int, int, int) {
	if list[b] < list[a] {
		a, b = b, a
	}
	if list[d] < list[c] {
		c, d = d, c
	}
	if list[c] < list[a] {
		if list[a] < list[d] {
			if list[b] < list[d] {
				return c, a, b, d
			} else {
				return c, a, d, b
			}
		} else {
			return c, d, a, b
		}
	} else {
		if list[c] < list[b] {
			if list[d] < list[b] {
				return a, c, d, b
			} else {
				return a, c, b, d
			}
		} else {
			return a, b, c, d
		}
	}
}

func triPartition_v2[T constraints.Ordered](list []T) (l, r int, skip bool) {
	size := len(list)
	m, s := size/2, size/4
	x, l, r, y := sort4(list, m-s, m-1, m+1, m+s)

	s = size - 1
	list[0], list[x] = list[x], list[0]
	list[s], list[y] = list[y], list[s]

	pivotL, pivotR := list[l], list[r]

	l, r = 1, s-1
	for {
		for list[l] < pivotL {
			l++
		}
		for list[r] > pivotR {
			r--
		}
		if list[l] > pivotR {
			list[l], list[r] = list[r], list[l]
			r--
			if list[l] < pivotL {
				l++
				continue
			}
		}
		break
	}

	for k := l + 1; k <= r; k++ {
		if list[k] > pivotR {
			for list[r] > pivotR {
				r--
			}
			if k >= r {
				break
			}
			if list[r] < pivotL {
				list[l], list[k], list[r] = list[r], list[l], list[k]
				l++
			} else {
				list[k], list[r] = list[r], list[k]
			}
			r--
		} else if list[k] < pivotL {
			list[k], list[l] = list[l], list[k]
			l++
		}
	}

	return l, r + 1, pivotL == pivotR
}
