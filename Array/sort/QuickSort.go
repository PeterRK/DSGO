package sort

//快速排序，改进的冒泡排序，不具有稳定性。
//平均复杂度为O(NlogN) & O(logN)，最坏情况是O(N^2) & O(N)。
//其中比较操作是O(NlogN)，常数与MergeSort相当；挪移操作是O(NlogN)，常数小于MergeSort。
func QuickSort(list []int) {
	var tasks stack
	tasks.push(0, len(list))
	for !tasks.isEmpty() {
		var start, end = tasks.pop()
		if end-start < sz_limit {
			InsertSort(list[start:end])
		} else {
			var knot = part(list[start:end]) + start
			tasks.push(knot+1, end)
			tasks.push(start, knot)
		} //每轮保证至少解决一个，否则最坏情况可能是死循环
	}
}

func part(list []int) int {
	var size = len(list)

	//三点取中法，最后保证seed落入中间，每轮至少解决此一处
	var seed = list[0]
	var mid, last = size / 2, size - 1
	if list[0] > list[mid] {
		if list[mid] > list[last] {
			seed, list[mid] = list[mid], list[0]
		} else { //c >= b
			if list[0] > list[last] {
				seed, list[last] = list[last], list[0]
			}
		}
	} else { //b >= a
		if list[last] > list[0] {
			if list[mid] > list[last] {
				seed, list[last] = list[last], list[0]
			} else {
				seed, list[mid] = list[mid], list[0]
			}
		}
	}

	var left, right = 1, last
	for { //注意对称性
		for list[left] < seed {
			left++
		}
		for list[right] > seed {
			right--
		}
		if left >= right {
			break
		}
		list[left], list[right] = list[right], list[left]
		left++
		right--
	}
	list[0], list[right] = list[right], seed

	return right
}
