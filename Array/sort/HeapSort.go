package sort

// 堆排序，改进的选择排序，不具备稳定性。
// 复杂度为O(NlogN) & O(1)。
// 其中比较操作是O(NlogN)，挪移操作也是O(NlogN)。
func HeapSort(list []int) {
	for idx := len(list)/2 - 1; idx >= 0; idx-- {
		down(list, idx)
	}
	for sz := len(list) - 1; sz > 0; sz-- {
		list[0], list[sz] = list[sz], list[0]
		down(list[:sz], 0)
	}
}

func down(list []int, pos int) {
	var key = list[pos]
	var kid, last = pos*2 + 1, len(list) - 1
	for kid < last {
		if list[kid+1] > list[kid] {
			kid++
		}
		if key >= list[kid] {
			break
		}
		list[pos] = list[kid]
		pos, kid = kid, kid*2+1
	}
	if kid == last && key < list[kid] {
		list[pos], pos = list[kid], kid
	}
	list[pos] = key
}
