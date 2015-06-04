package sort

import (
	"time"
)

//快速排序，改进的冒泡排序，不具有稳定性。
//平均复杂度为O(NlogN) & O(logN)，最坏情况是O(N^2) & O(N)。
//其中比较操作是O(NlogN)，常数与MergeSort相当；挪移操作是O(NlogN)，常数小于MergeSort。
//QuickSort不适合递归实现(有爆栈风险)。
func QuickSort(list []int) {
	magic = uint(time.Now().Unix())
	var tasks stack
	tasks.push(0, len(list))
	for !tasks.isEmpty() {
		var start, end = tasks.pop()
		if end-start < sz_limit {
			InsertSort(list[start:end])
		} else {
			var knot = partition(list[start:end]) + start
			tasks.push(knot+1, end)
			tasks.push(start, knot)
		} //每轮保证至少解决一个，否则最坏情况可能是死循环
	}
}

var magic = ^uint(0)

func partition(list []int) int {
	var size = len(list) //不少于3

	var x, y = int(magic % uint(size-1)), int(magic % uint(size-2))
	magic = magic*1103515245 + 12345
	var a, b = 1 + x, 1 + (1+x+y)%(size-1) //a != b
	//三点取中法，每轮至少解决一个
	var barrier = list[0]
	if list[0] > list[a] {
		if list[a] > list[b] {
			barrier, list[a] = list[a], list[0]
		} else { //c >= b
			if list[0] > list[b] {
				barrier, list[b] = list[b], list[0]
			}
		}
	} else { //b >= a
		if list[b] > list[0] {
			if list[a] > list[b] {
				barrier, list[b] = list[b], list[0]
			} else {
				barrier, list[a] = list[a], list[0]
			}
		}
	}

	a, b = 1, size-1
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
	list[0], list[b] = list[b], barrier
	return b
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
