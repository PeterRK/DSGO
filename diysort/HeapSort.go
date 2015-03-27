package diysort

//复杂度为O(log N) + O(1)，不具备稳定性
func HeapSort(list []int) {
	var size = len(list)
	if size < 7 {
		SelectSort(list)
		return
	}

	for idx := size/2 - 1; idx >= 0; idx-- {
		var root, left, right = idx, idx*2 + 1, idx*2 + 2
		var key = list[root]
		for right < size {
			var kid int
			if list[left] > list[right] {
				kid = left
			} else {
				kid = right
			}
			if key >= list[kid] {
				goto AdjustOver
			}
			list[root] = list[kid]
			root, left, right = kid, kid*2+1, kid*2+2
		}
		if right == size && key < list[left] {
			list[root], list[left] = list[left], key
			continue
		}
	AdjustOver:
		list[root] = key
	}

	for sz := size - 1; sz > 0; sz-- {
		var key = list[sz]
		list[sz] = list[0]
		var root, left, right = 0, 1, 2
		for right < sz {
			var kid int
			if list[left] > list[right] {
				kid = left
			} else {
				kid = right
			}
			if key >= list[kid] {
				goto AdjustOverX
			}
			list[root] = list[kid]
			root, left, right = kid, kid*2+1, kid*2+2
		}
		if right == sz && key < list[left] {
			list[root], list[left] = list[left], key
			continue
		}
	AdjustOverX:
		list[root] = key
	}
}
