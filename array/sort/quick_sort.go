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

func partition[T constraints.Ordered](list []T) int {
	//谨慎处理，以防越界
	//pivot := list[len(list)/2]
	pivot := median(list[0], list[len(list)/2], list[len(list)-1])

	a, b := 0, len(list)-1
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

//三点取中不一定很有用
func median[T constraints.Ordered](a, b, c T) T {
	if a > b {
		if b > c {
			return b //a b c
		} else if a > c {
			return c //a c b
		} else {
			return a //c a b
		}
	} else {
		if a > c {
			return a //b a c
		} else if b > c {
			return c //b c a
		} else {
			return b //c b a
		}
	}
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

// 返回两个分界元素的位置
func triPartition[T constraints.Ordered](list []T) (fst, snd int) {
	sz := len(list)
	m, s := sz/2, sz/8
	a, b := m-s, m+s
	if list[a] > list[b] {
		a, b = b, a
	}
	pivot1, pivot2 := list[a], list[b]
	list[a], list[b] = list[0], list[sz-1]

	a, b = 1, sz-2
	for a <= b && list[a] < pivot1 {
		a++
	}
	for k := a; k <= b; k++ {
		if list[k] > pivot2 {
			for k < b && list[b] > pivot2 {
				b--
			}
			list[k], list[b] = list[b], list[k]
			b--
		}
		if list[k] < pivot1 {
			list[k], list[a] = list[a], list[k]
			a++
		}
	}

	list[0], list[a-1] = list[a-1], pivot1
	list[sz-1], list[b+1] = list[b+1], pivot2
	return a - 1, b + 1
}

func moveMedianTwo[T constraints.Ordered](list []T, a, b, c, d int) {
	if list[a] > list[b] { //保证b非最小
		list[a], list[b] = list[b], list[a]
	}
	if list[c] < list[d] { //保证c非最大
		list[c], list[d] = list[d], list[c]
	}
	if list[b] > list[c] { //保证b<=c
		list[b], list[c] = list[c], list[b]
	}
}
func triPartition_v2[T constraints.Ordered](list []T) (fst, snd int) {
	sz := len(list)
	m, s := sz/2, sz/8
	moveMedianTwo(list, m-s, 0, sz-1, m+s)
	pivot1, pivot2 := list[0], list[sz-1]

	a, b := 1, sz-2
	for list[a] < pivot1 {
		a++
	}
	for k := a; k <= b; k++ {
		if list[k] > pivot2 {
			for list[b] > pivot2 {
				b--
			}
			if k > b {
				break
			}
			list[k], list[b] = list[b], list[k]
			b--
		}
		if list[k] < pivot1 {
			list[k], list[a] = list[a], list[k]
			a++
		}
	}

	list[0], list[a-1] = list[a-1], pivot1
	list[sz-1], list[b+1] = list[b+1], pivot2
	return a - 1, b + 1
}

func triPartition_v3[T constraints.Ordered](list []T) (fst, snd int) {
	sz := len(list)

	pivot1, pivot2 := list[0], list[sz-1]
	if pivot1 > pivot2 {
		pivot1, pivot2 = pivot2, pivot1
		list[0], list[sz-1] = pivot1, pivot2
	}

	ax, bx := 1, sz-2
	a, b := ax, bx
	for {
		for ; a <= b; a++ {
			if list[a] > pivot2 {
				break
			} else if list[a] <= pivot1 {
				if a != ax {
					list[ax], list[a] = list[a], list[ax]
				}
				ax++
			}
		}
		for ; a <= b; b-- {
			if list[b] < pivot1 {
				break
			} else if list[b] >= pivot2 {
				if b != bx {
					list[bx], list[b] = list[b], list[bx]
				}
				bx--
			}
		}
		if a >= b {
			break
		}
		if a == ax || b == bx {
			list[a], list[b] = list[b], list[a]
		} else {
			list[a], list[bx] = list[bx], list[a]
			list[b], list[ax] = list[ax], list[b]
			ax++
			bx--
			a++
			b--
		}
	}
	list[0], list[ax-1] = list[ax-1], pivot1
	list[sz-1], list[bx+1] = list[bx+1], pivot2
	return ax - 1, bx + 1
}
