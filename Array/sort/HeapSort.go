package sort

//堆排序，改进的选择排序，不具备稳定性。
//复杂度为O(NlogN) & O(1)。
//其中比较操作是O(NlogN)，挪移操作也是O(NlogN)。
func HeapSort(list []int) {
	if len(list) < sz_limit {
		InsertSort(list)
	} else {
		for idx := len(list)/2 - 1; idx >= 0; idx-- {
			down(list, idx)
		}
		for sz := len(list) - 1; sz > 0; sz-- {
			list[0], list[sz] = list[sz], list[0]
			down(list[:sz], 0)
		}
	}
}

func down(list []int, spot int) {
	var key = list[spot]
	var left, right = spot*2 + 1, spot*2 + 2
	for right < len(list) {
		var kid int
		if list[left] > list[right] {
			kid = left
		} else {
			kid = right
		}
		if key >= list[kid] {
			goto Label_OVER
		}
		list[spot] = list[kid]
		spot, left, right = kid, kid*2+1, kid*2+2
	}
	if right == len(list) && key < list[left] {
		list[spot], spot = list[left], left
	}
Label_OVER:
	list[spot] = key
}
