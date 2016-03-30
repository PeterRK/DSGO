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
	var sz = len(list)
	var a, b = sz/2 - 1, sz / 2
	if list[a] > list[b] {
		a, b = b, a
	}
	var pivot1, pivot2 = list[a], list[b]
	list[a], list[b] = list[0], list[sz-1]

	a, b = 1, sz-2
	for k := a; k <= b; k++ {
		if list[k] > pivot2 {
			for k < b && list[b] > pivot2 {
				b--
			}
			list[k], list[b] = list[b], list[k]
			b--
		}
		if list[k] < pivot1 {
			list[k], list[a] = list[a], list[k]
			a++
		}
	}

	list[0], list[a-1] = list[a-1], pivot1
	list[sz-1], list[b+1] = list[b+1], pivot2
	return a - 1, b + 1
}
