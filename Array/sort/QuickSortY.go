package sort

// 三分快速排序，比二分版本略为复杂
func QuickSortY(list []int) {
	for len(list) > LOWER_BOUND_Y {
		var fst, snd = triPartition(list)
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
func triPartition(list []int) (fst int, snd int) {
	var size = len(list)
	var m1, m2 = len(list)/2 - 1, len(list) / 2
	if list[m1] > list[m2] {
		m1, m2 = m2, m1
	}
	var pivot1, pivot2 = list[m1], list[m2]
	list[m1], list[m2] = list[0], list[size-1]

	var left, right = 1, size - 2
	for k := left; k <= right; k++ {
		if list[k] > pivot2 {
			for k < right && list[right] > pivot2 {
				right--
			}
			list[k], list[right] = list[right], list[k]
			right--
		}
		if list[k] < pivot1 {
			list[k], list[left] = list[left], list[k]
			left++
		}
	}

	list[0], list[left-1] = list[left-1], pivot1
	list[size-1], list[right+1] = list[right+1], pivot2
	return left - 1, right + 1
}
