package sort

// 快速排序，改进的冒泡排序，不具有稳定性。
// 平均复杂度为O(NlogN) & O(logN)，最坏情况是O(N^2) & O(N)。
// 其中比较操作是O(NlogN)，常数与MergeSort相当；挪移操作是O(NlogN)，常数小于MergeSort。
// 这里采用递归实现，但QuickSort不适合递归实现(有爆栈风险)。
func QuickSort(list []Unit) {
	if len(list) < LOWER_BOUND {
		SimpleSort(list)
	} else {
		m := partition(list)
		QuickSort(list[:m])
		QuickSort(list[m:])
	}
}

func partition(list []Unit) int {
	//谨慎处理，以防越界
	//pivot := list[len(list)/2]
	pivot := median(list[0].val,
		list[len(list)/2].val, list[len(list)-1].val)

	a, b := 0, len(list)-1
	for { //注意对称性
		for list[a].val < pivot {
			a++
		}
		for list[b].val > pivot {
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
func median(a, b, c int) int {
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

/*
func QuickSort(list []Unit) {
	var tasks stack
	tasks.push(0, len(list))
	for !tasks.isEmpty() {
		start, end := tasks.pop()
		if end-start < LOWER_BOUND {
			SimpleSort(list[start:end])
		} else {
			knot := partition(list[start:end]) + start
			tasks.push(knot, end)
			tasks.push(start, knot)
		}
	}
}

type pair struct {
	start int
	end   int
}
type stack struct {
	core []pair
}

func (s *stack) size() int {
	return len(s.core)
}
func (s *stack) isEmpty() bool {
	return len(s.core) == 0
}
func (s *stack) push(start int, end int) {
	s.core = append(s.core, pair{start, end})
}
func (s *stack) pop() (start int, end int) {
	sz := len(s.core) - 1
	unit := s.core[sz]
	s.core = s.core[:sz]
	return unit.start, unit.end
}
*/

//增加判断条件，化不稳定排序为稳定排序
/*
func QuickSortS(list []Unit) {
	for i := 0; i < len(list); i++ {
		list[i].pad[0] = uint32(i)
	}
	doQuickSortS(list)
}

func doQuickSortS(list []Unit) {
	for len(list) > LOWER_BOUND_Y {
		fst, snd := triPartitionS(list)
		doQuickSortS(list[:fst])
		doQuickSortS(list[snd+1:])
		if list[fst] == list[snd] {
			return
		}
		list = list[fst+1 : snd]
	}
	SimpleSort(list)
}

func triPartitionS(list []Unit) (fst, snd int) {
	sz := len(list)
	m, s := sz/2, sz/8
	a, b := m-s, m+s
	if list[a].val > list[b].val ||
		(list[a].val == list[b].val && list[a].pad[0] > list[b].pad[0]) {
		a, b = b, a
	}
	pivot1, pivot2 := list[a], list[b]
	list[a], list[b] = list[0], list[sz-1]

	a, b = 1, sz-2
	for a <= b && (list[a].val < pivot1.val ||
		(list[a].val == pivot1.val && list[a].pad[0] < pivot1.pad[0])) {
		a++
	}
	for k := a; k <= b; k++ {
		if list[k].val > pivot2.val ||
			(list[k].val == pivot2.val && list[k].pad[0] > pivot2.pad[0]) {
			for k < b && (list[b].val > pivot2.val ||
				(list[b].val == pivot2.val && list[b].pad[0] > pivot2.pad[0])) {
				b--
			}
			list[k], list[b] = list[b], list[k]
			b--
		}
		if list[k].val < pivot1.val ||
			(list[k].val == pivot1.val && list[k].pad[0] < pivot1.pad[0]) {
			list[k], list[a] = list[a], list[k]
			a++
		}
	}

	list[0], list[a-1] = list[a-1], pivot1
	list[sz-1], list[b+1] = list[b+1], pivot2
	return a - 1, b + 1
}
*/

// 三分快速排序，比二分版本略为复杂
func QuickSortY(list []Unit) {
	for len(list) > LOWER_BOUND_Y {
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
func triPartition(list []Unit) (fst, snd int) {
	sz := len(list)
	m, s := sz/2, sz/8
	a, b := m-s, m+s
	if list[a].val > list[b].val {
		a, b = b, a
	}
	pivot1, pivot2 := list[a], list[b]
	list[a], list[b] = list[0], list[sz-1]

	a, b = 1, sz-2
	for a <= b && list[a].val < pivot1.val {
		a++
	}
	for k := a; k <= b; k++ {
		if list[k].val > pivot2.val {
			for k < b && list[b].val > pivot2.val {
				b--
			}
			list[k], list[b] = list[b], list[k]
			b--
		}
		if list[k].val < pivot1.val {
			list[k], list[a] = list[a], list[k]
			a++
		}
	}

	list[0], list[a-1] = list[a-1], pivot1
	list[sz-1], list[b+1] = list[b+1], pivot2
	return a - 1, b + 1
}

/*
func triPartition(list []Unit) (fst, snd int) {
	sz := len(list)
	m, s := sz/2, sz/8
	moveMedianTwo(list, m-s, 0, sz-1, m+s)
	pivot1, pivot2 := list[0], list[sz-1]

	a, b := 1, sz-2
	for list[a].val < pivot1.val {
		a++
	}
	for k := a; k <= b; k++ {
		if list[k].val > pivot2.val {
			for list[b].val > pivot2.val {
				b--
			}
			if k > b {
				break
			}
			list[k], list[b] = list[b], list[k]
			b--
		}
		if list[k].val < pivot1.val {
			list[k], list[a] = list[a], list[k]
			a++
		}
	}

	list[0], list[a-1] = list[a-1], pivot1
	list[sz-1], list[b+1] = list[b+1], pivot2
	return a - 1, b + 1
}

func moveMedianTwo(list []Unit, a, b, c, d int) {
	if list[a].val > list[b].val { //保证b非最小
		list[a], list[b] = list[b], list[a]
	}
	if list[c].val < list[d].val { //保证c非最大
		list[c], list[d] = list[d], list[c]
	}
	if list[b].val > list[c].val { //保证b<=c
		list[b], list[c] = list[c], list[b]
	}
}
*/

/*
func triPartition(list []Unit) (fst, snd int) {
	sz := len(list)

	pivot1, pivot2 := list[0], list[sz-1]
	if pivot1.val > pivot2.val {
		pivot1, pivot2 = pivot2, pivot1
		list[0], list[sz-1] = pivot1, pivot2
	}

	ax, bx := 1, sz-2
	a, b := ax, bx
	for {
		for ; a <= b; a++ {
			if list[a].val > pivot2.val {
				break
			} else if list[a].val <= pivot1.val {
				if a != ax {
					list[ax], list[a] = list[a], list[ax]
				}
				ax++
			}
		}
		for ; a <= b; b-- {
			if list[b].val < pivot1.val {
				break
			} else if list[b].val >= pivot2.val {
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
*/
