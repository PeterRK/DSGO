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
