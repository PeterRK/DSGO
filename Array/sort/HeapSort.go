package sort

//复杂度为O(NlogN) & O(1)，不具备稳定性
//比较操作是O(NlogN)，挪移操作也是O(NlogN)
//建堆开销为O(N)
func HeapSort(list []int) {
	var size = len(list)
	if size < sz_limit {
		InsertSort(list)
	} else {
		for idx := size/2 - 1; idx >= 0; idx-- {
			down(list, idx)
		}
		for sz := size - 1; sz > 0; sz-- {
			list[0], list[sz] = list[sz], list[0]
			down(list[:sz], 0)
		}
	}
}

func down(list []int, spot int) {
	var size = len(list)
	var key = list[spot]
	var left, right = spot*2 + 1, spot*2 + 2
	for right < size {
		var kid int
		if list[left] > list[right] {
			kid = left
		} else {
			kid = right
		}
		if key >= list[kid] {
			goto LabelOver
		}
		list[spot] = list[kid]
		spot, left, right = kid, kid*2+1, kid*2+2
	}
	if right == size && key < list[left] {
		list[spot], list[left] = list[left], key
		return
	}
LabelOver:
	list[spot] = key
}
