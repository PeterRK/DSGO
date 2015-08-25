package sort

//快速排序，改进的冒泡排序，不具有稳定性。
//平均复杂度为O(NlogN) & O(logN)，最坏情况是O(N^2) & O(N)。
//其中比较操作是O(NlogN)，常数与MergeSort相当；挪移操作是O(NlogN)，常数小于MergeSort。
//这里采用递归实现，但QuickSort不适合递归实现(有爆栈风险)。
func QuickSort(list []int) {
	if len(list) < sz_limit {
		InsertSort(list)
	} else {
		var knot = partition(list)
		QuickSort(list[:knot])
		QuickSort(list[knot:])
	}
}

func partition(list []int) int {
	var barrier = list[len(list)/2] //该选择影响后面是否可能发生越界
	var a, b = 0, len(list) - 1
	for { //注意对称性
		for list[a] < barrier {
			a++
		}
		for list[b] > barrier {
			b--
		}
		if a >= b {
			break
		}
		//以下的交换操作是主要开销所在
		list[a], list[b] = list[b], list[a]
		a++
		b--
	}
	return a
}

/*
func QuickSort(list []int) {
	var tasks stack
	tasks.push(0, len(list))
	for !tasks.isEmpty() {
		var start, end = tasks.pop()
		if end-start < sz_limit {
			InsertSort(list[start:end])
		} else {
			var knot = partition(list[start:end]) + start
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
	var sz = len(s.core) - 1
	var unit = s.core[sz]
	s.core = s.core[:sz]
	return unit.start, unit.end
}
*/
